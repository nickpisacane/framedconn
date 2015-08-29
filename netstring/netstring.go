// Frame Format len:data,
//

package netstring

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"unicode"

	//"github.com/Nindaff/framedconn"
)

const (
	LEN_DELIMITER   = byte(':')
	FRAME_DELIMITER = byte(',')
	defaultMsgSize  = 4096
)

var (
	ErrBadDelimiter   = errors.New("netstring: bad frame delimiter")
	ErrBadWrite       = errors.New("netstring: bad frame write")
	ErrExpectedDigit  = errors.New("netstring: expected digit")
	ErrUnexpectedChar = errors.New("netstring: unexpected character")
	ErrMessageToLarge = errors.New("netstring: message to large")
)

var (
	errClearBufferFail = errors.New("netstring: failed to clear buffer")
)

type NetStringConn struct {
	MaxBytes int

	conn        net.Conn
	br          *bufio.Reader
	maxLenBytes int
}

func NewNetStringConn(conn net.Conn) *NetStringConn {
	return NewNetStringConnSize(conn, defaultMsgSize)
}

func NewNetStringConnSize(conn net.Conn, size int) *NetStringConn {
	return &NetStringConn{
		MaxBytes:    size,
		conn:        conn,
		br:          bufio.NewReaderSize(conn, size),
		maxLenBytes: len(fmt.Sprintf("%d", size)),
	}
}

// Reads `maxLenBytes + 1` one byte at a time. If a byte is
// read and is not a digit (base 10) as
func (ns *NetStringConn) readLen() (int, error) {
	buf := []byte{}
	i := 0
	for {
		// ReadByte() will return io.EOF error if there is nothing to
		// read. Just continue the loop and block until there is something
		// to read if EOF is returned.
		b, err := ns.br.ReadByte()
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				return 0, err
			}
		}
		i += 1
		if isAsciiDigit(rune(b)) {
			if i == ns.maxLenBytes+1 {
				return 0, ErrUnexpectedChar
			} else {
				buf = append(buf, b)
			}
		} else {
			if b == LEN_DELIMITER {
				break
			} else {
				return 0, ErrUnexpectedChar
			}
		}
	}
	n, err := strconv.ParseUint(string(buf), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// Reads a Frame. Blocks until entire frame is read into a buffer
// then returns the buffer. If the message is to large, an attempt
// is made to clear discard the entire message from the underlying
// buffer, panics if this fails.
func (ns *NetStringConn) ReadFrame() ([]byte, error) {
	msgLen, err := ns.readLen()
	if err != nil {
		return nil, err
	}
	if msgLen > ns.MaxBytes {
		// Discard the message plus the delimiter, if Discard fails
		// to discard `msgLen + 1`, err will not be nil. At this point
		// panicing is the best option as the next read would be compromised
		n, err := ns.br.Discard(msgLen + 1)
		if err != nil || n < msgLen {
			panic(errClearBufferFail)
		}
		return nil, ErrMessageToLarge
	}

	// increment by one for the delimiter
	msgLen += 1
	msg := make([]byte, msgLen)
	bytesRead := 0
	for {
		n, err := ns.br.Read(msg[bytesRead:])
		if err != nil {
			return nil, err
		}
		bytesRead += n
		if bytesRead >= int(msgLen) {
			break
		}
	}

	if msg[msgLen-1] != FRAME_DELIMITER {
		// Attempt to clear buffer
		if _, err := ns.br.Discard(ns.br.Buffered()); err != nil {
			panic(errClearBufferFail)
		}
		return nil, ErrBadDelimiter
	}
	return msg[:msgLen-1], nil
}

func (ns *NetStringConn) WriteFrame(data []byte) error {
	msgLen := len(data)
	msg := joinBytes([]byte(fmt.Sprintf("%d:", msgLen)), data, []byte(","))
	n, err := ns.conn.Write(msg)
	if err != nil {
		return err
	}
	if n < msgLen {
		return ErrBadWrite
	}
	return nil
}

func (ns *NetStringConn) Conn() net.Conn {
	return ns.conn
}

func joinBytes(bs ...[]byte) []byte {
	return bytes.Join(bs, []byte(""))
}

// 0-9 ascii
func isAsciiDigit(r rune) bool {
	return '0' <= r && r <= '9' && r < unicode.MaxASCII
}

// Connection wrapper for framing TCP messages

package netstring

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"unicode"
)

const (
	lenDelimiter     = byte(':')
	frameDelimiter   = byte(',')
	defaultFrameSize = 4096
)

var (
	ErrBadDelimiter   = errors.New("netstring: bad frame delimiter")
	ErrBadWrite       = errors.New("netstring: bad frame write")
	ErrUnexpectedChar = errors.New("netstring: unexpected character")
	ErrFrameToLarge   = errors.New("netstring: frame to large")
)

var (
	errClearBufferFail = errors.New("netstring: failed to clear buffer")
)

// Implements FramedConn
type NetStringConn struct {
	// Underlying conn
	conn net.Conn

	// Max frame size in bytes. This is not a part of netstring.
	maxFrameSize int

	// Max len size in bytes. This is the len in `[len]:[string],` format.
	// The len in bytes is computed automattically, knowing this helps prevent
	// unessary blocking and/or buffer allocation for bad clients that
	// may be abusing or neglecting netstring protocol.
	maxLenSize int

	sync.Mutex // Gaurds buffered Reader
	br         *bufio.Reader
}

// New netstring conn with max frame size of `defaultFrameSize`
func NewNetStringConn(conn net.Conn) *NetStringConn {
	return NewNetStringConnSize(conn, defaultFrameSize)
}

// Create a netstring conn with custom frame size (in bytes)
func NewNetStringConnSize(conn net.Conn, size int) *NetStringConn {
	return &NetStringConn{
		conn:         conn,
		br:           bufio.NewReaderSize(conn, size),
		maxFrameSize: size,
		maxLenSize:   len(fmt.Sprintf("%d", size)),
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
			if i == ns.maxLenSize+1 {
				return 0, ErrUnexpectedChar
			} else {
				buf = append(buf, b)
			}
		} else {
			if b == lenDelimiter {
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
	ns.Lock()
	defer ns.Unlock()

	msgLen, err := ns.readLen()
	if err != nil {
		return nil, err
	}
	if msgLen > ns.maxFrameSize {
		// Discard the message plus the delimiter, if Discard fails
		// to discard `msgLen + 1`, err will not be nil. At this point
		// panicing is the best option as the next read would be compromised
		n, err := ns.br.Discard(msgLen + 1)
		if err != nil || n < msgLen {
			panic(errClearBufferFail)
		}
		return nil, ErrFrameToLarge
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

	if msg[msgLen-1] != frameDelimiter {
		// Attempt to clear buffer
		if _, err := ns.br.Discard(ns.br.Buffered()); err != nil {
			panic(errClearBufferFail)
		}
		return nil, ErrBadDelimiter
	}
	return msg[:msgLen-1], nil
}

// Write bytes to the connection in netstring format. The bytes are formated
// before sent.
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

// Return underlying conn
func (ns *NetStringConn) Conn() net.Conn {
	return ns.conn
}

// Close underlying conn.
func (ns *NetStringConn) Close() error {
	_, _ = ns.br.Discard(ns.br.Buffered())
	return ns.conn.Close()
}

func joinBytes(bs ...[]byte) []byte {
	return bytes.Join(bs, []byte(""))
}

// 0-9 ascii
func isAsciiDigit(r rune) bool {
	return '0' <= r && r <= '9' && r < unicode.MaxASCII
}

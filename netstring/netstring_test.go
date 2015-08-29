package netstring

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

const (
	addr = "127.0.0.1:3000"
)

var (
	bad = errors.New("ASs")
)

type pair struct {
	server, client *NetStringConn
}

func newPair() (*pair, error) {
	p := &pair{}
	done := make(chan error, 1)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		done <- err
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				done <- err
				break
			}
			p.server = NewNetStringConn(conn)
			done <- nil
			break
		}
	}()

	client, err := net.Dial("tcp", addr)
	if err != nil {
		done <- err
	}
	p.client = NewNetStringConn(client)

	err = <-done
	if err != nil {
		p.client.Conn().Close()
		p.server.Conn().Close()
		return nil, err
	}
	return p, err
}

func makePayload(size uint) []byte {
	b := make([]byte, size)
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}
	return b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestMain(t *testing.T) {
	// TODO
}

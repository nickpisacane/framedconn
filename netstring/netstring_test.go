package netstring

import (
	"bytes"
	"net"
	"testing"
)

const (
	addr = "127.0.0.1:3000"
)

func createPair() (client *NetStringConn, server *NetStringConn, e error) {
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
			server = NewNetStringConn(conn)
			done <- nil
			break
		}
	}()

	c, err := net.Dial("tcp", addr)
	if err != nil {
		done <- err
	}
	client = NewNetStringConn(c)

	e = <-done
	if err != nil {
		client.Conn().Close()
		server.Conn().Close()
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestFrame(t *testing.T) {
	// create client, server pair
	client, server, err := createPair()
	check(err)

	// Frame equality server
	frame := []byte("Hello, World.")
	err = client.WriteFrame(frame)
	check(err)
	f, err := server.ReadFrame()
	check(err)

	if !bytes.Equal(frame, f) {
		t.Error("Expected frame %s, recieved %s", frame, f)
	}

	// Frame equality client
	err = server.WriteFrame(frame)
	check(err)
	f, err = client.ReadFrame()
	check(err)

	if !bytes.Equal(frame, f) {
		t.Error("Expected frame %s, recieved %s", frame, f)
	}

	// Send a bad frame to the server. When server reads the frame
	// err should not be nil
	badFrame := []byte("2:ping,")
	// Bypass netstring
	_, err = client.conn.Write(badFrame)
	check(err)
	_, err = server.ReadFrame()
	if err == nil {
		t.Error("ReadFrame() should have failed for frame overflow")
	}

	// Check that the server cleared the buffer after the frame overflow
	// as to not corrupt any succeding frames
	goodFrame := []byte("good frame")
	err = client.WriteFrame(goodFrame)
	check(err)
	f, err = server.ReadFrame()
	check(err)
	if !bytes.Equal(goodFrame, f) {
		t.Errorf("Buffer failed to clear. Expected frame %s, recieved %s", goodFrame, f)
	}

	// Check frames are consistent across go routines
	frame = []byte("concurrency")
	err = server.WriteFrame(frame)
	check(err)
	first := make(chan []byte, 1)
	second := make(chan []byte, 1)
	go func() {
		f, err = client.ReadFrame()
		check(err)
		first <- f
	}()
	go func() {
		f, err = client.ReadFrame()
		check(err)
		second <- f
	}()

	verify := func(incoming []byte, msg []byte) {
		if !bytes.Equal(incoming, msg) {
			t.Errorf("Expected frame %s, recieved %s", msg, incoming)
		}
	}

	// Which ever finishes first check that the frame is equal to
	// `frame`, then send a message to client and the opposing go routine
	// should read the correct frame
	select {
	case firstRecv := <-first:
		close(first)
		if !bytes.Equal(frame, firstRecv) {
			t.Errorf("Expected frame %s, recieved %s", frame, firstRecv)
		}
		secondFrame := []byte("Second")
		err = server.WriteFrame(secondFrame)
		check(err)
		verify(<-second, secondFrame)
		close(second)
	case secondRecv := <-second:
		close(second)
		if !bytes.Equal(frame, secondRecv) {
			t.Errorf("Expected frame %s, recieved %s", frame, secondRecv)
		}
		firstFrame := []byte("First")
		err = server.WriteFrame(firstFrame)
		check(err)
		verify(<-first, firstFrame)
		close(first)
	}

	err = client.Close()
	check(err)
	err = server.Close()
	check(err)
}

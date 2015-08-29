# Netstring Framing
See <a href="http://cr.yp.to/proto/netstrings.txt">Netstrings</a>

## Getting Started
```sh
$ go get github.com/Nindaff/framedconn/netstring
```

## Usage

```go

func main() {
  conn, _ := net.Dial("tcp", "127.0.0.1:3000")
  nsClient := netstring.NewNetStringConn(conn)

  _ = nsClient.WriteFrame([]byte("ping")) // writes `4:ping,` to the connection
  // Assume server responds with `4:pong,`
  frame, _ = nsClient.ReadFrame()
  string(frame) == "pong" // true
}
```

# netstring
--
    import "github.com/Nindaff/framedconn/netstring"

#### type NetStringConn

```go
type NetStringConn struct {
	sync.Mutex // Gaurds buffered Reader
}
```

Implements FramedConn

#### func  NewNetStringConn

```go
func NewNetStringConn(conn net.Conn) *NetStringConn
```
New netstring conn with max frame size of `defaultFrameSize`

#### func  NewNetStringConnSize

```go
func NewNetStringConnSize(conn net.Conn, size int) *NetStringConn
```
Create a netstring conn with custom frame size (in bytes)

#### func (*NetStringConn) Close

```go
func (ns *NetStringConn) Close() error
```
Close underlying conn.

#### func (*NetStringConn) Conn

```go
func (ns *NetStringConn) Conn() net.Conn
```
Return underlying conn

#### func (*NetStringConn) ReadFrame

```go
func (ns *NetStringConn) ReadFrame() ([]byte, error)
```
Reads a Frame. Blocks until entire frame is read into a buffer then returns the
buffer. If the message is to large, an attempt is made to clear discard the
entire message from the underlying buffer, panics if this fails.

#### func (*NetStringConn) WriteFrame

```go
func (ns *NetStringConn) WriteFrame(data []byte) error
```
Write bytes to the connection in netstring format. The bytes are formated before
sent.

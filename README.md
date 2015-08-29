# framedconn -- Simple tcp framing
This package serve's as an interface that framing packages in this repo will adhere to.


#### Installation
--
    import "github.com/Nindaff/framedconn"

#### type FramedConn

```go
type FramedConn interface {
  // Reads the next "frame" from the connection.
  ReadFrame() ([]byte, error)

  // Writes a "frame" to the connection.
  WriteFrame(p []byte) error

  // Closes the connections, truncates any buffers.
  Close() error

  // Returns the underlying connection.
  Conn() net.Conn
}
```

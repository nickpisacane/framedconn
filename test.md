# framedconn
--
    import "github.com/Nindaff/framedconn"


## Usage

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

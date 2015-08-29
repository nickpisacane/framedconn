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

## Godoc
	
		<div id="short-nav">
			<dl>
			<dd><code>import "/home/nindaff/go/src/github.com/Nindaff/framedconn/netstring"</code></dd>
			</dl>
			<dl>
			<dd><a href="#pkg-overview" class="overviewLink">Overview</a></dd>
			<dd><a href="#pkg-index" class="indexLink">Index</a></dd>
			
			
			</dl>
		</div>
		<!-- The package's Name is printed as title by the top-level template -->
		<div id="pkg-overview" class="toggleVisible">
			<div class="collapsed">
				<h2 class="toggleButton" title="Click to show Overview section">Overview ▹</h2>
			</div>
			<div class="expanded">
				<h2 class="toggleButton" title="Click to hide Overview section">Overview ▾</h2>
				
			</div>
		</div>
		

		<div id="pkg-index" class="toggleVisible">
		<div class="collapsed">
			<h2 class="toggleButton" title="Click to show Index section">Index ▹</h2>
		</div>
		<div class="expanded">
			<h2 class="toggleButton" title="Click to hide Index section">Index ▾</h2>

		<!-- Table of contents for API; must be named manual-nav to turn off auto nav. -->
			<div id="manual-nav">
			<dl>
			
			
				<dd><a href="#pkg-variables">Variables</a></dd>
			
			
			
				
				<dd><a href="#NetStringConn">type NetStringConn</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NewNetStringConn">func NewNetStringConn(conn net.Conn) *NetStringConn</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NewNetStringConnSize">func NewNetStringConnSize(conn net.Conn, size int) *NetStringConn</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#NetStringConn.Close">func (ns *NetStringConn) Close() error</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NetStringConn.Conn">func (ns *NetStringConn) Conn() net.Conn</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NetStringConn.ReadFrame">func (ns *NetStringConn) ReadFrame() ([]byte, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NetStringConn.WriteFrame">func (ns *NetStringConn) WriteFrame(data []byte) error</a></dd>
				
			
			
			</dl>
			</div><!-- #manual-nav -->

		

		
			<h4>Package files</h4>
			<p>
			<span style="font-size:90%">
			
				<a href="/src/target/netstring.go">netstring.go</a>
			
			</span>
			</p>
		
		</div><!-- .expanded -->
		</div><!-- #pkg-index -->

		<div id="pkg-callgraph" class="toggle" style="display: none">
		<div class="collapsed">
			<h2 class="toggleButton" title="Click to show Internal Call Graph section">Internal call graph ▹</h2>
		</div> <!-- .expanded -->
		<div class="expanded">
			<h2 class="toggleButton" title="Click to hide Internal Call Graph section">Internal call graph ▾</h2>
			<p>
			  In the call graph viewer below, each node
			  is a function belonging to this package
			  and its children are the functions it
			  calls&mdash;perhaps dynamically.
			</p>
			<p>
			  The root nodes are the entry points of the
			  package: functions that may be called from
			  outside the package.
			  There may be non-exported or anonymous
			  functions among them if they are called
			  dynamically from another package.
			</p>
			<p>
			  Click a node to visit that function's source code.
			  From there you can visit its callers by
			  clicking its declaring <code>func</code>
			  token.
			</p>
			<p>
			  Functions may be omitted if they were
			  determined to be unreachable in the
			  particular programs or tests that were
			  analyzed.
			</p>
			<!-- Zero means show all package entry points. -->
			<ul style="margin-left: 0.5in" id="callgraph-0" class="treeview"></ul>
		</div>
		</div> <!-- #pkg-callgraph -->

		
		
			<h2 id="pkg-variables">Variables</h2>
			
				<pre>var (
    <span id="ErrBadDelimiter">ErrBadDelimiter</span>   = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;netstring: bad frame delimiter&#34;)
    <span id="ErrBadWrite">ErrBadWrite</span>       = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;netstring: bad frame write&#34;)
    <span id="ErrUnexpectedChar">ErrUnexpectedChar</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;netstring: unexpected character&#34;)
    <span id="ErrFrameToLarge">ErrFrameToLarge</span>   = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;netstring: frame to large&#34;)
)</pre>
				
			
		
		
		
			
			
			<h2 id="NetStringConn">type <a href="/src/target/netstring.go?s=622:1118#L25">NetStringConn</a></h2>
			<pre>type NetStringConn struct {
    <a href="/pkg/sync/">sync</a>.<a href="/pkg/sync/#Mutex">Mutex</a> <span class="comment">// Gaurds buffered Reader</span>
    <span class="comment">// contains filtered or unexported fields</span>
}</pre>
			<p>
Implements FramedConn
</p>


			

			

			
			
			

			
				
				<h3 id="NewNetStringConn">func <a href="/src/target/netstring.go?s=1184:1235#L43">NewNetStringConn</a></h3>
				<pre>func NewNetStringConn(conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>) *<a href="#NetStringConn">NetStringConn</a></pre>
				<p>
New netstring conn with max frame size of `defaultFrameSize`
</p>

				
				
			
				
				<h3 id="NewNetStringConnSize">func <a href="/src/target/netstring.go?s=1355:1420#L48">NewNetStringConnSize</a></h3>
				<pre>func NewNetStringConnSize(conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, size <a href="/pkg/builtin/#int">int</a>) *<a href="#NetStringConn">NetStringConn</a></pre>
				<p>
Create a netstring conn with custom frame size (in bytes)
</p>

				
				
			

			
				
				<h3 id="NetStringConn.Close">func (*NetStringConn) <a href="/src/target/netstring.go?s=4119:4157#L165">Close</a></h3>
				<pre>func (ns *<a href="#NetStringConn">NetStringConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
				<p>
Close underlying conn.
</p>

				
				
				
			
				
				<h3 id="NetStringConn.Conn">func (*NetStringConn) <a href="/src/target/netstring.go?s=4031:4071#L160">Conn</a></h3>
				<pre>func (ns *<a href="#NetStringConn">NetStringConn</a>) Conn() <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a></pre>
				<p>
Return underlying conn
</p>

				
				
				
			
				
				<h3 id="NetStringConn.ReadFrame">func (*NetStringConn) <a href="/src/target/netstring.go?s=2647:2699#L100">ReadFrame</a></h3>
				<pre>func (ns *<a href="#NetStringConn">NetStringConn</a>) ReadFrame() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
				<p>
Reads a Frame. Blocks until entire frame is read into a buffer
then returns the buffer. If the message is to large, an attempt
is made to clear discard the entire message from the underlying
buffer, panics if this fails.
</p>

				
				
				
			
				
				<h3 id="NetStringConn.WriteFrame">func (*NetStringConn) <a href="/src/target/netstring.go?s=3735:3789#L146">WriteFrame</a></h3>
				<pre>func (ns *<a href="#NetStringConn">NetStringConn</a>) WriteFrame(data []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
				<p>
Write bytes to the connection in netstring format. The bytes are formated
before sent.
</p>

				
				
				
			
		
	

	






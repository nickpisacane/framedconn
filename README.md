# framedconn -- Simple tcp framing
This package serve's as an interface that framing packages in this repo will adhere to.

	
## Godoc		
		<div id="short-nav">
			<dl>
			<dd><code>import "/home/nindaff/go/src/github.com/Nindaff/framedconn"</code></dd>
			</dl>
			<dl>
			<dd><a href="#pkg-overview" class="overviewLink">Overview</a></dd>
			<dd><a href="#pkg-index" class="indexLink">Index</a></dd>
			
			
				<dd><a href="#pkg-subdirectories">Subdirectories</a></dd>
			
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
			
			
			
			
				
				<dd><a href="#FramedConn">type FramedConn</a></dd>
				
				
			
			
			</dl>
			</div><!-- #manual-nav -->

		

		
			<h4>Package files</h4>
			<p>
			<span style="font-size:90%">
			
				<a href="/src/target/framedconn.go">framedconn.go</a>
			
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

		
		
		
		
			
			
			<h2 id="FramedConn">type <a href="/src/target/framedconn.go?s=39:338#L1">FramedConn</a></h2>
			<pre>type FramedConn interface {
    <span class="comment">// Reads the next &#34;frame&#34; from the connection.</span>
    ReadFrame() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// Writes a &#34;frame&#34; to the connection.</span>
    WriteFrame(p []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// Closes the connections, truncates any buffers.</span>
    Close() <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// Returns the underlying connection.</span>
    Conn() <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>
}</pre>
			

			

			

			
			
			

			

			
		
	

	





	
	
		<h2 id="pkg-subdirectories">Subdirectories</h2>
	
	


	<div class="pkg-dir">
		<table>
			<tr>
				<th class="pkg-name">Name</th>
				<th class="pkg-synopsis">Synopsis</th>
			</tr>

			
			<tr>
				<td colspan="2"><a href="..">..</a></td>
			</tr>
			

			
				
					<tr>
						<td class="pkg-name" style="padding-left: 0px;">
							<a href="netstring/">netstring</a>
						</td>
						<td class="pkg-synopsis">
							
						</td>
					</tr>
				
			
		</table>
	</div>


	


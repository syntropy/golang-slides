Title: A Practical Introduction to Go

I could bore you with describing every single feature of Go one by one.

...or, I could simply show you how to develop a simple network server that does
something useful and is easy to understand.

---
Title: Telnet Chat Server

I will show you how to develop a telnet chat server, step by step.

This is what it's supposed to be doing:

* The user connects to the chat server by telnet and is asked for a nickname.

* After entering the nickname, the user joins the chat room.

* By typing and pressing return, the user can post text that all other connected users can see.

* When the user disconnects, all other users are informed that the user left the chat room.

---
Title: First iteration: accepting connections

<pre class="prettyprint" data-lang="go">
package main

import ( "log"; "net" ; "os" )

func main() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
</pre>
---
Title: First iteration: handling the connection

<pre class="prettyprint" data-lang="go">
func handleConnection(c net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		n, err = c.Write(buf[0:n])
		if err != nil {
			c.Close()
			break
		}
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())
}
</pre>
---
Title: What are goroutines?

Goroutines are

* a function executing concurrently with other goroutines in the same address space.

* very lightweight, small stack, grows and shrinks on demand.

* goroutines are multiplexed to multiple OS threads. If a goroutine blocks waiting for I/O, others will run.

* think of them as very lightweight threads that are cheap to create and destroy.

---
Title: Second iteration

* What we created so far is a simple echo server.

* To get closer to our goal, let's extend the echo server to send everything it receives to a central goroutine that logs all received data to stdout.

---
Title: Second iteration: providing a communication channel

<pre class="prettyprint" data-lang="go">
	msgchan := make(chan string)
	go printMessages(msgchan)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn, msgchan)
	}
</pre>
---
Title: Second iteration: printing all received messages

<pre class="prettyprint" data-lang="go">
func printMessages(msgchan <-chan string) {
	for msg := range msgchan {
		log.Printf("new message: %s", msg)
	}
}
</pre>

---
Title: Second iteration: sending all messages to the central process

<pre class="prettyprint" data-lang="go">
func handleConnection(c net.Conn, msgchan chan<- string) {
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		msgchan <- string(buf[0:n])
		// ...
	}
}
</pre>

---
Title: What are channels?

* Channels are a means of communication that allows sending and receiving objects of a certain data type.

* In addition, the semantics of channels provide a way to synchronize different goroutines.

* You can send objects to a channel, and receive objects from a channel.

* Receiving from a channel will always block until there is data to receive.

* Sending to an unbuffered channel will block until data is read from the channel.

* Sending to a buffered channel will only block until the data has been copied to the buffer. If the buffer is full, it will block until space has been freed (i.e. by receiving from the channel).

* Think of channels as something like type-safe Unix pipes within your program.

---
Title: Combining goroutines and channels

* Usually, concurrent programming is hard and only low-level primitives are available. Managing concurrent access to shared variables is error-prone.

* Go encourages another way of tackling concurrency instead: data is passed around between goroutines through channels. Only the responsible goroutine gets to modify data. Thus data races cannot occur.

* This way of designing concurrent programs can be subsumed with:
  <i></i>
  <h2>Do not communicate by sharing memory; instead, share memory by communicating.</h2>

* The theoretical foundation of goroutines and channels are Communicating Sequential Processes (CSP) by C. A. R. Hoare.

---
Title: Third iteration: what's missing?

* So far, we have an application where client connections send everything to a central processing goroutine.

* Now we need to extend this with the following parts:

    * a way to register new clients

    * a way to unregister clients that have disconnected

    * messages sent by a client shall be broadcast to all clients

---
Title: Third iteration: registering new clients

First, we create a channel to register new clients and hand it to all the functions that need it:
<pre class="prettyprint" data-lang="go">
addchan := make(chan Client)
go handleMessages(msgchan, addchan)
// ...
go handleConnection(conn, msgchan, addchan)
</pre>

That's what the `Client` type looks like:

<pre class="prettyprint" data-lang="go">
type Client struct {
	conn net.Conn
	ch   chan<- string
}
</pre>

---
Title: Third iteration: registering new clients

Then, in `handleConnection`, we register the client, identified by the
connection, together with a channel with which `handleConnection` will receive
messages that are sent back to the client.

<pre class="prettyprint" data-lang="go">
func handleConnection(c net.Conn, msgchan chan<- string, addchan chan<- Client) {
	ch := make(chan string)

	addchan <- Client{c, ch}
	// ...
}
</pre>

---
Title: Third iteration: registering new clients and broadcasting messages

in `handleMessages`, new clients are received and added to a map that contains
all active connections together with the associated channel. New messages are
also broadcast here.

<pre class="prettyprint" data-lang="go">
func handleMessages(msgchan <-chan string, addchan <-chan Client) {
	clients := make(map[net.Conn]chan<- string)
	for {
		select {
		case msg := <-msgchan:
			for _, ch := range clients {
				go func(mch chan<- string) { mch <- "\033[1;33;40m" + msg + "\033[m\r\n" }(ch)
			}
		case client := <-addchan:
			clients[client.conn] = client.ch
		}
	}
}
</pre>

---
Title: Third iteration: unregistering disconnected clients

To unregister disconnected clients, we create another channel and – again – hand it
to all the functions that need it:

<pre class="prettyprint" data-lang="go">
// in func main()
rmchan := make(chan Client)

go handleMessages(msgchan, addchan, rmchan)
// ...
go handleConnection(conn, msgchan, addchan, rmchan)
</pre>

---
Title: Third iteration: unregistering disconnected clients

To unregister disconnected clients, we simply send them to this channel in the `handleConnection` function
when the function returns.

<pre class="prettyprint" data-lang="go">
func handleConnection(c net.Conn, msgchan chan<- string, addchan chan<- Client, rmchan chan<- Client) {
	// ...
	defer func() {
		rmchan <- c
	}()
	// ...
}
</pre>

---
Title: Third iteration: unregistering disconnected clients

In `handleMessages`, these disconnected clients are received and their connection removed from the map.

<pre class="prettyprint" data-lang="go">
func handleMessages(msgchan <-chan string, addchan <-chan Client, rmchan <-chan net.Conn) {
	clients := make(map[net.Conn]chan<- string)

	for {
		select {
		// ...
		case conn := <-rmchan:
			delete(clients, conn.ch)
		}
	}
}
</pre>

---
Title: Third iteration: fixing the handleConnection function

* the `handleConnection` function needs to both receive data and forward it to `handleMessages` and send other messages back to the client.

* blocking reads and `select` don't work together... so, we need to move the reading to another goroutine.

* the `handleConnection` should also query the newly connected user for the desired nickname.

---
Title: Third iteration: fixing the handleConnection function

<pre class="prettyprint" data-lang="go">
func handleConnection(c net.Conn, msgchan chan<- string, addchan chan<- Client, rmchan chan<- Client) {
	bufc := bufio.NewReader(c)
	defer c.Close()
	client := Client{
			conn:     c,
			nickname: promptNick(c, bufc),
			ch:       make(chan string),
	}
	if strings.TrimSpace(client.nickname) == "" {
		io.WriteString(c, "Invalid Username\n")
		return
	}
</pre>

---
Title: Third iteration: fixing the handleConnection function

<pre class="prettyprint" data-lang="go">
	addchan <- client
	defer func() {
		msgchan <- fmt.Sprintf("User %s left the chat room.\n", client.nickname)
		log.Printf("Connection from %v closed.\n", c.RemoteAddr())
		rmchan <- client
	}
	io.WriteString(c, fmt.Sprintf("Welcome, %s!\n\n", client.nickname))
	msgchan <- fmt.Sprintf("New user %s has joined the chat room.\n", client.nickname)

	go client.ReadLinesInto(msgchan)
	client.WriteLinesFrom(client.ch)
}
</pre>

---
Title: Third iteration: promptNick()

<pre class="prettyprint" data-lang="go">
func promptNick(c net.Conn, bufc *bufio.Reader) string {
	io.WriteString(c, "\033[1;30;41mWelcome to the fancy demo chat!\033[0m\n")
	io.WriteString(c, "What is your nick? ")
	nick, _, _ := bufc.ReadLine()
	return string(nick)
}
</pre>

---
Title: Third iteration: ReadLinesInto()

<pre class="prettyprint" data-lang="go">
func (c Client) ReadLinesInto(ch chan<- string) {
	bufc := bufio.NewReader(c.conn)
	for {
		line, err := bufc.ReadString('\n')
		if err != nil {
			break
		}
		ch <- fmt.Sprintf("%s: %s", c.nickname, line)
	}
}
</pre>

---
Title: Third iteration: WriteLinesFrom()

<pre class="prettyprint" data-lang="go">
func (c Client) WriteLinesFrom(ch <-chan string) {
	for msg := range ch {
		_, err := io.WriteString(c.conn, msg)
		if err != nil {
			return
		}
	}
}
</pre>

---
Title: Third iteration: finished!

We now have a fully functional telnet multi-user chat application, in about 100 lines of Go code.

That wasn't too terrible, was it?

The full source code for these examples is available here:

[https://github.com/akrennmair/telnet-chat](https://github.com/akrennmair/telnet-chat)

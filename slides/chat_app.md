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

import ( "fmt" ; "net" ; "os" )

func main() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
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
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
</pre>
---
Title: What are goroutines?

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
			fmt.Println(err)
			continue
		}

		go handleConnection(conn, msgchan)
	}
</pre>
---
Title: Second iteration: printing all received messages

<pre class="prettyprint" data-lang="go">
func printMessages(msgchan <-chan string) {
	for {
		msg := <-msgchan
		fmt.Printf("new message: %s\n", msg)
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

* Go encourages another way of tackling concurrency instead: data is passed around between goroutines through channels. Only the responsible goroutine gets to modify data. Data races thus cannot occur.

* This way of designing concurrent programs can be subsumed with 

<h2>Do not communicate by sharing memory; instead, share memory by communicating.</h2>


Title: So, what else is there?

In the following slides, I will show you all the features of Go that I haven't discussed before.

* Object-oriented programming

* A brief overview over the standard library

---
Title: Object-oriented programming

* When it comes to object-oriented programming, Go handles things a bit differently.

* Focus on simplicity, which may confuse users who already know other OO languages.

* No inheritance.

---
Title: Data types, methods and interfaces

* Data types define the content

* Methods define operations on the content

* Interfaces define the behaviour of a data type (i.e. a set of methods that you can use on an object)

---
Title: Data types

You can create new types from any existing type:

<pre class="prettyprint" data-lang="go">
type km_h float64

type m_s float64

type Coordinate2D struct {
	X, Y int
}
</pre>

Unlike <code>typedef</code>s in C, <code>type</code>s in Go aren't mere aliases. They are different types
that you cannot simply assign to each other, you need an explicit conversion.

	t.go:10: cannot use bar (type m_s) as type km_h in assignment

---
Title: Methods

You can define methods on data types:

<pre class="prettyprint" data-lang="go">
func (c *Coordinate2D) DistanceToZero() int {
	return int(math.Sqrt(float64(c.X*c.X + c.Y*c.Y)))
}
</pre>

---
Title: Interfaces

Interfaces are implicit. Once you define an interface, every data type that implements the required
methods fulfills that interface.

<pre class="prettyprint" data-lang="go">
type Reader interface {
	Read(data []byte) (int, error)
}
</pre>

By convention, interface names end in "-er" and are usually named by the method name they contain, e.g. `Reader`, `Writer`, `ReadWriteCloser`. The standard libary makes extensive use of interfaces.

<pre class="prettyprint" data-lang="go">
conn, err := net.Dial("tcp", "example.com:9876")
compressor, err := gzip.NewWriter(conn)
json_encoder := json.NewEncoder(compressor)
json_encoder.Encode(my_data)
</pre>

This code connects to `example.com:9876`, encodes `my_data` (an object of any type) to JSON, gzip-compresses it, and sends it over the TCP connection.
---
Title: Inheritance

As mentioned before, Go doesn't support inheritance. Instead, it has embedding.

You can embed an existing data type into a struct type. The struct type then inherits (ha!) all the methods of the embedded type.

<pre class="prettyprint" data-lang="go">
type Foo int

func (f *Foo) Bla() {
	fmt.Println("called Bla()")
}

type Bar struct {
	Foo
}

func (b *Bar) Whoop() {
	fmt.Println("called Whoop()")
}
</pre>

---
Title: Standard library: web applications

Go's standard library is rich in features that allow getting quickly into programming web-related tools and servers.

* HTTP server, client (`net/http`)

* HTML template system (`html/template`)

* JSON and XML encoders/decoders (`encoding/json`, `encoding/xml`)

---
Title: HTTP server

The most simple web "app" you can write in Go:

<pre class="prettyprint" data-lang="go">
package main
import (
   "fmt"
   "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
func main() {
   http.ListenAndServe(":8000", http.HandlerFunc(handler))
}
</pre>

More sophisticated features are available, including prefix-based routing and
TLS. For an even more complete experience that builds upon Go's standard
library, using the [Gorilla web toolkit](http://gorilla-web.appspot.com/) is
recommended.
---
Title: JSON

The simplest way to read/write JSON is to use `encoding/json`'s `Marshal()` and `Unmarshal()` functions.

<pre class="prettyprint" data-lang="go">
type Comment {
	Id     int64  `json:"id"`
	Author string `json:"author"`
	Text   string `json:"comment"`
}

// ...

comment := &Comment{Id: 42, Author: "Anonymous Coward", Text: "First post!!11!"}
if json_data, err := json.Marshal(comment); err == nil {
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_data)
}
</pre>
---
Title: JSON, XML

<pre class="prettyprint" data-lang="go">
var coordinates []int
if err = json.Unmarshal([]byte(json_coords), &coordinates); err == nil {
	// ...
}
</pre>

Working with XML is practically the same, except the annotations are different:

<pre class="prettyprint" data-lang="go">
type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
}
</pre>
---
Title: Even more standard library (1)

* Archive readers/writers (tar, zip)

* Compression algorithms (bzip2, deflate, gzip, lzw, zlib)

* Cryptography (AES, DES, DSA, ECDSA, HMAC, MD5, RC4, RSA, SHA{1,256,512}, X.509)

* Generic SQL database interface (with external drivers)

* Encodings (ASCII85, ASN.1, base{32,64}, CSV, hexadecimal, JSON, PEM, XML)

* Parser for Go source code

* Hash algorithms (Adler32, CRC32, CRC64, FNV-1(a))

* HTML and text templates

---
Title: Even more standard library (2)

* Image reading/writing (GIF, JPEG, PNG)

* Reflection

* Regular expressions

* Unit tests

* ...

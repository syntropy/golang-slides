Title: Declaring Variables

You can declare a list of variables including their data type with the var keyword:

<pre class="prettyprint" data-lang="go">
var x, y, z int
var name string
</pre>

Initializing variables in the declaration is also possible. It can even derive the data type for you:

<pre class="prettyprint" data-lang="go">
var x, y, z int = 23, 42, 9001
var name = "Max Mustermann"
pi := 3.1415926 // short for var pi = ..., works only within functions
</pre>

---
Title: Constants

Constants are declared like variables, but with the const keyword. They can be optionally typed:

<pre class="prettyprint" data-lang="go">
const pi = 3.1415926
const e float64 = 2.71828
</pre>

Untyped constants take the type that they need by the context:

<pre class="prettyprint" data-lang="go">
const (
	three = 3
)
var (
	x int = three
	y float64 = three
)
</pre>

---
Title: Enumerated Constants

If you need an enumerated list of integer constants without wanting to write down all values
manually, Go provides you with the iota keyword. For every constant in a const list, the
iota keyword is equivalent to the constant's position in the list, starting with 0. Further
constants without explicit value derive their value from the previous expression with an
incremented iota.

<pre class="prettyprint" data-lang="go">
const (
	a = 23 + iota // a = 23
	b             // b = 24
	c             // c = 25
)

const (
	FOO = 1 << iota // FOO = 1
	BAR             // BAR = 2
	BAZ             // BAZ = 4
	QUUX            // QUUX = 8
)
</pre>

---
Title: Basic types

Go provides the following basic data types

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32; represents a Unicode code point

	float32 float64

	complex64 complex128

---
Title: Structs

A struct contains a number of fields with a data type and a name.

<pre class="prettyprint" data-lang="go">
type LogLevel int

type LogMsg struct {
	Level     LogLevel
	Msg       string
	Timestamp time.Time
}
</pre>


Title: Declaring Variables

You can declare a list of variables including their data type with the var keyword:

<pre class="prettyprint" data-lang="go">
var x, y, z int
var name string
</pre>

Initializing variables in the declaration is also possible. It can even derive the data type for you:

<pre class="prettyprint" data-lang="go">
var x, y, z int = 23, 42, 1e6
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

---
Title: Using structs

Fields in a struct can be accessed through ".":

<pre class="prettyprint" data-lang="go">
type Coordinate struct {
	X, Y, Z float64
}

c := Coordinate{1.9, 2.2, -2.5}
// you can also name the fields in the initialization:
c := Coordinate{Z:1.1, X:9.3, Y:0.1}
</pre>

Visibility of fields to other packages is controlled by the case of a field
name's first character. Upper-case field names are public, lower-case
field-names are private.

---
Title: Pointers

Go has pointers, but no pointer arithmetics (well, at least none without the unsafe package).

The fields of structs can be access through pointers in the same way as with
structs: the dot operator is transparent here.

<pre class="prettyprint" data-lang="go">
c := Coordinate{0.0, 1.0, 2.0}
x := &c                           // creates a pointer to c
p := &Coordinate{3.0, 4.0, 5.0}   // creates a new Coordinate object and lets p point to it
x.X, p.X = 23, 42                 // accessing fields
</pre>

For basic data types, you need to use the new operator:

<pre class="prettyprint" data-lang="go">
i := new(int)
*i = 23

c := new(Coordinate)
</pre>
---
Title: Maps

A map allows you store key-value pairs and retrieve the value by key. Maps must be created
with make.

<pre class="prettyprint" data-lang="go">
var m map[string]Coordinate

m = make(map[string]Coordinate
// or:
n := make(map[string]Coordinate)

m["Couch"] = Coordinate{0.0, 0.0, 0.0}

fmt.Printf("%#v", m["Couch"])

o := map[string]Coordinate{
	"foo": Coordinate{0.0, 1.0, 2.0},
	"bar": Coordinate{3.0, 4.0, 5.0},
}
</pre>
---
Title: Working with maps

Insert or update an element:

<pre class="prettyprint" data-lang="go">
m[key] = value
</pre>

Retrieve an element by key:
<pre class="prettyprint" data-lang="go">
value = m[key]
</pre>

Delete an element by key:
<pre class="prettyprint" data-lang="go">
delete(m, key)
</pre>

Test whether an element is present.  If an element for key was found, then ok
is true. Otherwise, ok is false and value is zero.
<pre class="prettyprint" data-lang="go">
value, ok = m[key]
</pre>

---
Title: Functions

A function takes 0 or more arguments, and has 0 or more return values. Arguments
are always named, return values can optionally be named.

<pre class="prettyprint" data-lang="go">
func add(x, y int) int {
	return x + y
}
</pre>

You can have more than one return value, and you can name them:

<pre class="prettyprint" data-lang="go">
func fib(a, b int) (x, y int) {
	x, y = b, a + b
	return
}
</pre>

---
Title: Functions as values

In Go, functions can also be used as values:

<pre class="prettyprint" data-lang="go">
mul := func(a, b int) int {
	return a * b
}
fmt.Println(mul(3, 4))
</pre>

---
Title: Functions as closures

In Go, functions are also closures:

<pre class="prettyprint" data-lang="go">
func create_id_generator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}
</pre>
---
Title: If statements (1)

If statements look like in C, with a few exceptions: the ( ) aren't necessary,
and the { } are mandatory.

<pre class="prettyprint" data-lang="go">
if x < 0 {
	// ... is executed if "x < 0" is true
} else if y > 0 {
	// ... is executed if "y > 0" is true
} else {
	// ... is executed otherwise
}
</pre>

---
Title: If statements (2)

The if statement allows you to execute a statement before the actual condition.
Variables declared in the statement are scoped to the if statement.

<pre class="prettyprint" data-lang="go">
if data, err := fetch(); err == nil {
	process(data)
} else {
	fmt.Printf("an error occured while fetching data: %v\n", err)
	return err
}
</pre>
---
Title: The switch statement

The main difference of Go's switch statement to that of other languages is that
case bodies break automatically. To prevent such an automatic break, Go provides
the fallthrough statement. In addition, if you leave out the condition, switch
can serve as a clean alternative to writing long if-then-else chains.

<pre class="prettyprint" data-lang="go">
switch x := foo(); x {
case 0:
case 1: fmt.Println("foobar")
case 2: fmt.Println("x is 2")
		fallthrough
case 3: fmt.Println("quux")
}

switch {
case x < 0: // ...
case y > 0: // ...
default:    // ...
}
</pre>
---
Title: Loops

Go only has one loop: for. It looks like in C, except the ( ) are gone, and the
{ } are mandatory. Pre statement, loop condition, and post statement are
semicolon-separated.

If you leave out pre and post statement or even all of the three statements,
you can leave out the semicolons.

<pre class="prettyprint" data-lang="go">
for i := 0 ; i < 20 ; i++ {
	// for loop like everyone knows it
}

for i < 30 {
	// equivalent with C's while loop
}

for { 
	// infinite loop
}
</pre>

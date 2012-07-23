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


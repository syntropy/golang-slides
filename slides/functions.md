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
func CreateIDGenerator() <b>func() int</b> {
	id := 0
	return func() int {
		id++
		return id
	}
}
</pre>

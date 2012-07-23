Title: Arrays

An array is a numbered sequence of elements of a certain type. Arrays are values,
that means all elements are copied when you assign an array to another. The size
of an array is part of the type. That means that two arrays of the same element
type but of different sizes are of different types.

<pre class="prettyprint" data-lang="go">
var a [10]int // array of 10 element of type int

b := [...]float64{0.1, 7.5, 3.1} // array of 3 elements of type float64

fmt.Printf("%f", b[1])
</pre>

Using arrays in Go is not very common. Slices are a similar but more general
mechanism that is more powerful than arrays.

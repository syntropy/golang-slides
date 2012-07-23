Title: Slices

A slice points to a range of element of an array, and includes components of length and capacity.

You can create slices by making them point to an existing array, or by creating a new one
with the make function.

<pre class="prettyprint" data-lang="go">
// create an array:
arr := [..]int{3,5,2,1}

// create a slice from parts of arr, starting at index 1, ranging up to but excluding index 3.
slice := arr[1:3] // -> []int{5, 2}

// create another empty slice
another_slice := []int{}
</pre>

---
Title: Manipulating slices

All you basically need to work with slices is the len function to compute a slice's length,
the append function to append elements to a slice and the range syntax to select a range
from an existing slice or array to create a new one.

<pre class="prettyprint" data-lang="go">
a = append(a, x)									// append element x to a slice

a = append(a, b...)									// append slice b to slice a

a = append(a[:i], a[i+1:]...)						// delete element at index i

a = append(a[:i], append([]int{x}, a[i:]...)...)	// insert element x at index i

a = append(a[:i], append(b, a[i:]...)...)			// insert slice b at index i
</pre>

More such tricks: [http://code.google.com/p/go-wiki/wiki/SliceTricks](http://code.google.com/p/go-wiki/wiki/SliceTricks)

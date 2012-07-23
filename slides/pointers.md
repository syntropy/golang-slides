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

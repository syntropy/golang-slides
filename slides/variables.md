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


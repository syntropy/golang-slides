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


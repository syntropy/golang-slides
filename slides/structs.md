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


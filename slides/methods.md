Title: Methods

Go does *not* have classes. However, you can define methods on custom data types. These
can, but need not be structs.

You can define methods similar to regular functions. The method receiver is specified
between the func keyword and the method name.

<pre class="prettyprint" data-lang="go">
type Foo int

func (f *Foo) Bar() {
	*f++
}
// ...
x := Foo(3)
x.Bar()		// will increase x to 4
</pre>

If you don't specify the method receiver as pointer, the method will work on a copy of the
method receiver and thus not manipulate the method receiver.

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

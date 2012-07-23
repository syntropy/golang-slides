Title: Loops

Go only has one loop: for. It looks like in C, except the ( ) are gone, and the
{ } are mandatory. Pre statement, loop condition, and post statement are
semicolon-separated.

If you leave out pre and post statement or even all of the three statements,
you can leave out the semicolons.

<pre class="prettyprint" data-lang="go">
for i := 0 ; i < 20 ; i++ {
	// for loop like everyone knows it
}

for i < 30 {
	// equivalent with C's while loop
}

for { 
	// infinite loop
}
</pre>

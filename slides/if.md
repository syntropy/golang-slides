Title: If statements (1)

If statements look like in C, with a few exceptions: the ( ) aren't necessary,
and the { } are mandatory.

<pre class="prettyprint" data-lang="go">
if x < 0 {
	// ... is executed if "x < 0" is true
} else if y > 0 {
	// ... is executed if "y > 0" is true
} else {
	// ... is executed otherwise
}
</pre>

---
Title: If statements (2)

The if statement allows you to execute a statement before the actual condition.
Variables declared in the statement are scoped to the if statement.

<pre class="prettyprint" data-lang="go">
if data, err := fetch(); err == nil {
	process(data)
} else {
	fmt.Printf("an error occured while fetching data: %v\n", err)
	return err
}
</pre>

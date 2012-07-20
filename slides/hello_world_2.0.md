title: Hello, world! 2.0

<pre class="prettyprint" data-lang="go">
package main
import (
   "fmt"
   "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) { 
   fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]) 
}
func main() {
   http.ListenAndServe(":8000", http.HandlerFunc(handler))
}
</pre>

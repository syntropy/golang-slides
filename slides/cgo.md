Title: Integrating with existing C libraries: cgo

Go can easily integrate with existing C libraries

You need to start your own library, though. But besides that, it's really easy:

* Start a new .go source file

* `import "C"`

* Embed the C code you want to interface as comment (e.g. `#include <foo.h>`) prior to the `import`

* Imported C functions, types and global variables are available under the virtual "C" package

* Optional: `#cgo` preprocessor instructions to set C defines and linker flags.

---
Title: First Example (1)

<pre class="prettyprint" data-lang="go">
package pwnam
/*
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Passwd struct {
	Uid   uint32
	Gid   uint32
	Dir   string
	Shell string
}
</pre>

---
Title: First Example (2)

<pre class="prettyprint" data-lang="go">
func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.cfree(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	return &Passwd{ 
		Uid: uint32(cpw.pw_uid), 
		Gid: uint32(cpw.pw_gid), 
		Dir: C.GoString(cpw.pw_dir), 
		Shell: C.GoString(cpw.pw_shell) }
}
</pre>

---
Title: Linking

* Usually, you want to access functions outside of libc

* Linking of external libraries is required

* Use `#cgo` directive

* cgo can also use pkg-config

---
Title: Linking Examples

<pre class="prettyprint" data-lang="go">
package pcap
/*
#cgo LDFLAGS: -lpcap
#include <stdlib.h>
#include <pcap.h>
*/
import "C"
</pre>

<pre class="prettyprint" data-lang="go">
package stfl
/*
#cgo pkg-config: stfl
#cgo LDFLAGS: -lncursesw
#include <stdlib.h>
#include <stfl.h>
*/
import "C"
</pre>

---
Title: Mapping the C namespace to Go

* Everything declared in the C code is available in the `C` pseudo-package

* Fundamental C data types have their counterpart, e.g. `int` &rarr; `C.int`, `unsigned short` &rarr; `C.ushort`, etc.

* The Go equivalent to `void *` is `unsafe.Pointer`

* `typedef`s are available under their own name

* `struct`s are available with `struct_` prefix, e.g. `struct foo` &rarr; `C.struct_foo`, same goes for `union`s and `enum`s

---
Title: Conversion between C and Go strings

* The `C` package contains conversion functions to convert Go to C strings and vice versa

* Also: opaque data (behind `void *`) to `[]byte`

<pre class="prettyprint" data-lang="go">
// Go string to C string; result must be freed with C.free
func C.CString(string) *C.char

// C string to Go string
func C.GoString(*C.char) string

// C string, length to Go string
func C.GoStringN(*C.char, C.int) string

// C pointer, length to Go []byte
func C.GoBytes(unsafe.Pointer, C.int) []byte
</pre>

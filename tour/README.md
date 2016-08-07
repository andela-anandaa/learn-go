## Short Notes from _The Tour_

_Tour_ documentation offline:

```
$ go tool tour
```

Full documentation offline:
```
$ godoc -http=:xxxx
```

Quick reference check from command-line, eg for `fmt.Print`:
```
$ godoc cmd/fmt Print
```

Factored import statements, e.g.

```go
import(
	"fmt"
	"math"
)
```

- When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package. e.g. `math.pi` vs. `math.Pi`
- Naked return statements should be used only in short functions. They can harm readability in longer functions. 
- Variables declared without an explicit initial value are given their zero value.
- Constants cannot be declared using the `:=` syntax. 


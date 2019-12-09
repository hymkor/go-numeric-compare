go-numeric-compare
==================

Compare strings that contain numbers.

The numbers in the strings are compared by their values, 
not in dictionary order.

```go
package main

import (
	"fmt"
	"github.com/zetamatta/go-numeric-compare"
)

func test(a, b string) {
	cmp := numeric.Compare(a, b)
	fmt.Printf("%2d == numeric.Compare(\"%s\",\"%s\")\n", cmp, a, b)
}

func main() {
	test("100", "20")
	test("v80", "v100")
	test("v100", "v100")
	test("v120", "v100")
	test("a1", "a001")
}
```

```
$ go run example.go
 1 == numeric.Compare("100","20")
-1 == numeric.Compare("v80","v100")
 0 == numeric.Compare("v100","v100")
 1 == numeric.Compare("v120","v100")
 1 == numeric.Compare("a1","a001")
```

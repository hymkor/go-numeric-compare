go-numeric-compare
==================

Compare strings that contain numbers.

The numbers in the strings are compared by their values, 
not in dictionary order.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/zetamatta/go-numeric-compare"
)

func test(a, b string) {
	cmp := strings.Compare(a, b)
	fmt.Printf("%5d == strings.Compare (\"%s\",\"%s\")\n", cmp, a, b)
	cmp = numeric.Compare(a, b)
	fmt.Printf("%5d == numeric.Compare (\"%s\",\"%s\")\n", cmp, a, b)
	lt := numeric.LessThan(a, b)
	fmt.Printf("%5v == numeric.LessThan(\"%s\",\"%s\")\n", lt, a, b)
	fmt.Println()
}

func main() {
	test("alpha", "beta")
	test("v80", "v100")
	test("1.10.4.1", "1.2.3.4")
	test("a1", "a001")
}
```

```
$ go run example.go
   -1 == strings.Compare ("alpha","beta")
   -1 == numeric.Compare ("alpha","beta")
 true == numeric.LessThan("alpha","beta")

    1 == strings.Compare ("v80","v100")
   -1 == numeric.Compare ("v80","v100")
 true == numeric.LessThan("v80","v100")

   -1 == strings.Compare ("1.10.4.1","1.2.3.4")
    1 == numeric.Compare ("1.10.4.1","1.2.3.4")
false == numeric.LessThan("1.10.4.1","1.2.3.4")

    1 == strings.Compare ("a1","a001")
    1 == numeric.Compare ("a1","a001")
false == numeric.LessThan("a1","a001")``
```

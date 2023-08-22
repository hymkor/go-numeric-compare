//go:build run

package main

import (
	"fmt"
	"strings"

	"github.com/hymkor/go-numeric-compare"
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

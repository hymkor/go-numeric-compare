// +build run

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

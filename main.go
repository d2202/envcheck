package main

import (
	"fmt"

	"github.com/d2202/envcheck/parser"
)

func main() {
	fmt.Println(parser.Parse("tests/test.env"))
}

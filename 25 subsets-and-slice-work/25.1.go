package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Эта программа позволяет определить вхождение строки в подстроку")

	var str, substr string

	flag.StringVar(&str, "str", "default str", "set str")
	flag.StringVar(&substr, "substr", "default substr", "set substr")
	flag.Parse()

	fmt.Println(strings.Contains(str, substr))
}

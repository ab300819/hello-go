package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every", "The greeting object.")
}

func main() {
	fmt.Printf("hello, %s!\n", name)
}

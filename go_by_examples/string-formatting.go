package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Println("%v\n", p)
	fmt.Println("%+v\n", p)
	fmt.Println("%#v\n", p)



}

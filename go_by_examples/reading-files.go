package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//dat, err := ioutil.ReadFile("/Users/mason/Project/biomart/build.gradle")
	//check(err)
	//fmt.Print(string(dat))

	f, err := os.Open("/Users/mason/Project/biomart/build.gradle")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2,err:=f.Seek(6,0)
	check(err)
	b2:=make([]byte,2)
	n2,err:=f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ",n2,o2)
	fmt.Printf("%v\n",string(b2[:n2]))
}

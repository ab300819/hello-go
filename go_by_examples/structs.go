package main

import "fmt"

type person struct{
	name string
	age int
}

func NewPerson(name string) *person{
	p:=person{name:name}
	p.age=42
	return &p
}

func main(){
	
}
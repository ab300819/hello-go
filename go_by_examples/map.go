package main

import "fmt"

func main(){
	
	m:=make(map[string]int)

	m["k1"]=10
	m["k2"]=11
	fmt.Println("map: ",m)

	v1:=m["k1"]
	fmt.Println("v1:",v1)
	fmt.Println("len: ",len(m))

	delete(m,"k2")
	fmt.Println("map: ",m)

	val,prs:=m["k1"]
	fmt.Println("prs",val,prs)

	n:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map",n)

}
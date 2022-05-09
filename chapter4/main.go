package main

import "fmt"

func main() {

	a := 10
	b := &a
	c := *b
	fmt.Println(a == c)
}

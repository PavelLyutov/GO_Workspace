package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Golang"

	for i,s := range a{
		fmt.Printf("%d : %#s\n" ,i,s)
	}
}

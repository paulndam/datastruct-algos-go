package main

import "fmt"


func main(){


	// slices are abstraction over go arrays.
	var slice = [] int {1,2,3,4,5,6}
	slice = append(slice, 7,8)
	fmt.Println("my capacity is --->", cap(slice))
	fmt.Println("my length is ---->", len(slice))
}
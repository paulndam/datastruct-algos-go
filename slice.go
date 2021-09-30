package main

import "fmt"

//passing a slice to a reference as a function.
func twiceValue(slice [] int){
	var i int 
	var v int 
	
	for i, v = range slice {
		slice[i] = 2 * v
	}
}

func main(){
	// slices are abstraction over go arrays.
	var slice = [] int {1,2,3,4,5,6}
	slice = append(slice, 7,8)
	fmt.Println("my capacity is --->", cap(slice))
	fmt.Println("my length is ---->", len(slice))

	var slices = [] int {1,3,5,7,9,12}
	twiceValue(slices)

	var i int 

	for i = 0; i < len(slices); i++{
		fmt.Println("new slice is ----->", slices[i])
	}

}
package main

import "fmt"

func main(){
	var arr = [] int {2,4,5,7,9,3}
	//excluding index 3.
	var slice1 = arr[:3]
	fmt.Println("slice1 --->",slice1)
	//including index 1 thru  index 5 but excluding index 5.
	var slice2 = arr[1:5]
	fmt.Println("slice 2 --->",slice2)
	//appending or adding element to slice.
	var slice3 = append(slice2,54)
	fmt.Println("--- slice3 -->", slice3)
}
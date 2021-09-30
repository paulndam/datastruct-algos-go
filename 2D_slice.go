package main

import "fmt"


func main(){
	var TwoDArray [8][8] int 
	TwoDArray[3][6] = 18 
	TwoDArray [7][4] = 3
	fmt.Println(TwoDArray)

	//slices of slices as two dimensional arrays.

	var rows int 
	var cols int 

	rows = 7
	cols = 9

	var TwoD_Slice = make([][] int, rows)
	var i int
	for i = range TwoD_Slice{
		TwoD_Slice[i] = make([] int, cols)
	}
	fmt.Println(TwoD_Slice)
}
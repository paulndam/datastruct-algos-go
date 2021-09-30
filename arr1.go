package main

import "fmt"


 
func main(){

	var arr = [5] int {1,2,3,4,5}
	// var value int
	for i,value := range arr{
		fmt.Println("my index is ---->",i,"and my value is --->",value)

	}

	// to remove index 
	for _,v := range arr{
		fmt.Println("no index but my value is ----->",v)
	}

}
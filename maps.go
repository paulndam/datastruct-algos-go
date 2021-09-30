package main

import "fmt"


func main(){

	// maps are use to keep track of keys which can be of type strings,int,pointers,interface,float,double,structs and arrays.Their values can be of different types.
	//                   key    value
	var languages = map [int] string {
		// keys are integers and values are strings
		1:"french",
		2:"italian",
		3:"english",
	}

	// maps can also br created with the make() method.
	var shoes = make(map[int]string)
	shoes[1] = "nike"
	shoes[2] = "puma"
	shoes[3] = "addidas"

	var i int 
	var v string

	for i,v = range languages{
		fmt.Println("my index is ---->",i,"and my value is ---->",v)
	}
	fmt.Println("shoes with at key 2 is ---->",shoes[2])

	// retrieving a value and deleting a value from a map.
	fmt.Println(shoes[3])
	delete(shoes,"puma")
	
}
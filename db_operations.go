package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//customer class.

type Customer struct {

	Id int 
	Name string 
	SSN string
	
}

// connection method which returns sql.db
func GetConnection() (database *sql.DB){
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "Gr@ndkanu05$"
	databaseName := "crm"

	database,err := sql.Open(databaseDriver,databaseUser + ":" + databasePass + "@/" + databaseName)


	if err != nil {
		panic(err.Error())
		fmt.Println("----- Database connection failed ----")
	}

	return database

}

// method to return customer array

func GetCustomers() []Customer{
	var database *sql.DB
	database = GetConnection()

	var err error
	var rows *sql.Rows
	rows,err = database.Query("SELECT * FROM Customer ORDER BY Id DESC")

	if err != nil{
		panic(err.Error())
		fmt.Println("can't return all users in DB")
	}

	var customer Customer 
	customer = Customer{}

	var customers []Customer
	customers = []Customer{}

	for rows.Next(){
		var Id int 
		var name string 
		var ssn string

		err = rows.Scan(&Id,&name,&ssn)

		if err != nil {
			panic(err.Error())
			fmt.Println("scanning customers attributes to DB failed")
		}

		customer.Id = Id
		customer.Name = name 
		customer.SSN = ssn 

		customers = append(customers,customer)

	}
	defer database.Close()

	return customers
}

// insert method.
func Insert(customer Customer) {
	var database * sql.DB
	database = GetConnection()

	var err error
	var insert *sql.Stmt

	insert,err = database.Prepare("INSERT INTO CUSTOMER(Id,Name,SSN) VALUES(?,?,?)")

	if err != nil {
		panic(err.Error())
		fmt.Println("couldn't insert customers successfully.")
	}

	insert.Exec(customer.Id,customer.Name,customer.SSN)
	defer database.Close()
}


//update method.
func Update(customer Customer){
	var database * sql.DB
	database = GetConnection()

	var err error
	var update *sql.Stmt

	update,err = database.Prepare("UPDATE CUSTOMER SET Name=?, SSN=? WHERE Id=?")

	if err != nil {
		panic(err.Error())
		fmt.Println("can't update user/customer")

	}

	update.Exec(customer.Name,customer.SSN,customer.Id)
	
	defer database.Close()


}


// delet customer.

func Delete(customer Customer){
	var database * sql.DB
	database = GetConnection()

	var err error
	var delete *sql.Stmt

	delete,err = database.Prepare("DELETE FROM Customer WHERE Id=?")

	if err != nil{
		panic(err.Error())
		fmt.Println("can't delete user/customer")

	}

	delete.Exec(customer.Id)

	defer database.Close()



}



func main(){

	var customers []Customer
	customers = GetCustomers()
	fmt.Println("----- Before inserting,updating and deleting customers in DB ---",customers)

	var customer Customer

	customer.Id = 1
	customer.Name = "Mary Joe"
	customer.SSN = "123-98-0536"

	// customer.Id = 3
	// customer.Name = "Tom Brady"
	// customer.SSN = "205-66-6666"


	// Insert(customer)

	// Update(customer)

	Delete(customer)

	customers = GetCustomers()
	// fmt.Println("---- Returning all customers from DB after inserting one or more customers --->",customers)

	// fmt.Println("----- Returning all customers from DB after updating customers  ---",customers)

	fmt.Println("----- Returning all customers from DB after deleting customers  ---",customers)
}
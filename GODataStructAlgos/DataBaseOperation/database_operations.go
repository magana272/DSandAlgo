package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// Make a connection to the database
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

// GetCustomers method return Customer array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	rows, err := database.Query("SELECT * FROM Customer ORDER Customerid DESC")
	if err != nil {
		fmt.Println("Couldn't connect")
		panic(err.Error())
	}
	var customer Customer
	customer = Customer{}
	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		err = rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}
func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Customers", customers)
}

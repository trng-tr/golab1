package utils2

import (
	"fmt"

	"github.com/google/uuid"
)

func inputCustomer() Customer {
	var customer Customer = Customer{}    // initialiser customer
	var uuid string = uuid.New().String() // généré from package  github.com/google/uuid v1.6.0
	customer.uuid = uuid
	fmt.Print("Entrer le firstname du customer: ")
	fmt.Scan(&customer.firstname)
	fmt.Print("Entrer le lastname du customer: ")
	fmt.Scan(&customer.lastname)
	fmt.Print("Entre l'email du customer :")
	fmt.Scan(&customer.email)

	return customer
}

func PrintCustomer() {
	var customer Customer = inputCustomer()
	fmt.Println("Customer object :", customer)
}

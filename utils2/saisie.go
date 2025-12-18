package utils2

import (
	"fmt"
)

func inputCustomer() Customer {
	var customer Customer = Customer{} // initialiser customer
	fmt.Print("Entrer l'id du customer: ")
	fmt.Scan(&customer.id)
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

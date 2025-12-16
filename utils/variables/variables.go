package variables

import "fmt"

func VariablesDeclaration() {
	var (
		age    byte
		taille byte
	)
	const hasBankAccount bool = true

	fmt.Println("saisir ton age et ta taille sur la même ligne séparer par des espaces: ")
	fmt.Scanln(&age, &taille)
	fmt.Printf("son age et sa taille sont %d ans %d cm\n", age, taille)
	fmt.Println("il/elle possede un bank account ", hasBankAccount)
}

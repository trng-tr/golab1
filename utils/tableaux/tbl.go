package tableaux

import "fmt"

func Accumulateur1() {
	declarerTableau()
	printTableau1()
	printTableau2()
	printTableau3()
	printTableau4()
	printTableau5()
	var jours [7]string = modifyElementsTable()
	fmt.Println(jours[:])
}

func Accumulateur2() {
	declareMatrix()
	matrix := initMatrix()
	fmt.Println(matrix)
}

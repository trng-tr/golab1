package anonymousfunc

import (
	"fmt"
	utils "golab1/utils"
)

func AnonymousFunction1() {
	value, err := utils.SaisirNumber()
	if err != nil {
		fmt.Println("", err)
		return
	}

	// fonction anonyme appelée immediatement
	func() {
		fmt.Println("Anonymous function 1 print :", value)
	}() // fonction anonyme appelée immediatement
}

func AnonymousFunction2() {
	value, err := utils.SaisirNumber()
	if err != nil {
		fmt.Println("", err)
		return
	}

	f := func() { // stockage de la fonc anonyme dans une variable
		fmt.Println("Anonymous function 2 print: ", value)
	}

	f() // appel de la fonction anonyme

}

/*
declaration compacte des parametres v1 et v2,
équivaut à AnonymousFunction3(v1 float, v2 float64)
*/
func AnonymousFunction3(v1, v2 float64) {
	// fonction anonyme avec params
	somme := func(n1 float64, n2 float64) {
		fmt.Printf("la somme de %f et %f est %f \n", n1, n2, (n1 + n2))
	}

	somme(v1, v2) // appell fonc anonyme
}

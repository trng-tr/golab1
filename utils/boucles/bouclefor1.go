package boucles

import (
	"errors"
	"fmt"
	utils "golab1/utils"
)

func NombrePairs(n int) ([]int, error) { // chercher les nombres pairs entre 0 et n>=10
	if n < 10 {
		return nil, errors.New("error: est trop petit")
	}
	var nPairs []int = []int{} // créer et initialiser un slice vide
	for counter := 0; counter <= n; counter++ {
		if counter%2 == 0 {
			nPairs = append(nPairs, counter)
		}
	}
	return nPairs, nil
}

func EcoleDimanche() string {
	var age float64
	var err error

	for age < 10 || age > 18 { // while condition vraie, rester dans la boucle
		age, err = utils.SaisirNumber() // reaffectation des nouvelles valeurs
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	return fmt.Sprintf("Bienvenue chez les ados, age %d ans\n", int(age))
}

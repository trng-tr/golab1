package boucles

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func saisirMinMax() ([]float64, error, error) {
	scanner := bufio.NewScanner(os.Stdin)
	//saisir min
	fmt.Print("Saisir la valeur du min: ")
	scanner.Scan()
	var str1 string = scanner.Text()
	if str1 == "" {
		return nil, errors.New("error: chaine vide"), nil
	}
	min, err1 := strconv.ParseFloat(str1, 64)
	if err1 != nil {
		return nil, errors.New("erreur de conversion"), nil
	}
	// saisir max
	fmt.Print("Saisir la valeur du max: ")
	scanner.Scan()
	var str2 string = scanner.Text()
	if str2 == "" {
		return nil, nil, errors.New("error: chaine vide")
	}
	max, err2 := strconv.ParseFloat(str2, 64)
	if err2 != nil {
		return nil, nil, errors.New("error: chaine vide")
	}

	if max <= min {
		return nil, errors.New("error: max <= min"), errors.New("error: max <= min")
	}
	var minMax []float64 = make([]float64, 0) // créer et initialiser slice
	return append(minMax, min, max), nil, nil // return a slice et err1=""et err2=""
}

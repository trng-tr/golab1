package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func SaisirNumber() (float64, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Saisir un nombre: ")
	scanner.Scan()
	var str string = scanner.Text()
	if str == "" {
		return -1, errors.New("error: chaine vide")
	}
	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return -1, errors.New("error: chaine vide")
	}

	return value, nil
}

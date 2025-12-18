package toint

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func convertToInt1() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Saisir une valeur à convertir en int: ")
	scanner.Scan()
	var varName string = scanner.Text()
	if varName == "" {
		return -1, errors.New("error: chaine vide")
	}
	var value, err = strconv.Atoi(varName)
	if err != nil {
		return -1, errors.New("erreur de conversion en int")
	}
	return value, nil
}

package tofloat

import (
	"errors"
	"strconv"
)

func ConvertToFloat(str string) (float64, error) {
	if str == "" {
		return -1, errors.New("error: chaine vide")
	}

	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return -1, errors.New("erreur de conversion en float")
	}
	return value, nil
}

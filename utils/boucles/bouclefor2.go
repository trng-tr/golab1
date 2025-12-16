package boucles

import (
	"math/rand"
)

func GetRandomValue() (int, error) {
	minMax, err1, err2 := saisirMinMax()
	if err1 != nil {
		return -1, err1
	} else if err2 != nil {
		return -1, err2
	}
	min, max := int(minMax[0]), int(minMax[1])
	return rand.Intn(max-min+1) + min, nil
}

package main

import (
	"fmt"
	"golab1/utils/calculs"
	"golab1/utils/conditions"
	"golab1/utils/reader"
)

func main() {
	for {
		poids, err := reader.ReadFromKeyboard()
		if err != "" {
			fmt.Printf("%#v\n", err)
			return
		}
		fmt.Printf("retour %d\n", poids)

		poids2, err2 := conditions.SaisirPoids()
		if err2 != nil {
			fmt.Printf("%#v\n", err2)
			return
		}

		age, err3 := conditions.SaisirAge()
		if err3 != nil {
			fmt.Printf(" %#v\n", err3)
			return
		}

		fmt.Printf("votre poids %d kg. votre age %d ans\n", poids2, age)

		//random value
		calculs.GetRandomValue()

	}
}

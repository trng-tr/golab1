package main

import (
	"fmt"
	"golab1/utils"
	"golab1/utils/boucles"
	"golab1/utils/conditions"
)

func main() {
	for {

		randValue, err := boucles.GetRandomValue()
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("random value %d\n", randValue)

		fmt.Printf("%s", boucles.EcoleDimanche())

		value, err1 := utils.SaisirNumber()
		if err1 != nil {
			fmt.Printf("%#v\n", err1)
			return
		}

		nPairs, err2 := boucles.NombrePairs(int(value))
		if err2 != nil {
			fmt.Printf("%#v\n", err2)
			return
		}

		fmt.Printf("%#v\n", nPairs)

		conditions.UsingSwitchCase()
		conditions.GetWeekEnd()
		conditions.TestOtherSwith()

	}
}

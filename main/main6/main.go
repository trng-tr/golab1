package main

import (
	"fmt"
	"golab1/utils"
	"golab1/utils/anonymousfunc"
	"math/rand"
)

func main() {
	anonymousfunc.AnonymousFunction1()
	anonymousfunc.AnonymousFunction2()

	v1, err1 := utils.SaisirNumber()
	if err1 != nil {
		fmt.Println("", err1)
		return
	}

	v2, err2 := utils.SaisirNumber()
	if err2 != nil {
		fmt.Println("", err2)
		return
	}

	anonymousfunc.AnonymousFunction3(v1, v2)

	var product int = anonymousfunc.AnonymousFunction4(int(v1), int(v2))
	fmt.Printf("le produit de %v et %v est %v \n", v1, v2, product)

	factoriel, err3 := anonymousfunc.AnonymousFunction5(int(v1))
	if err3 != nil {
		fmt.Println("error is raised :", err3)
		return
	}
	fmt.Printf("le factoriel de %v est %v", v1, factoriel)

	var slice []float64 = make([]float64, 0, int(v1))
	for i := 0; i < int(v1); i++ {
		slice = append(slice, float64(rand.Intn(int(v1)-10+1)+10))
	}

	sliceNew, err := anonymousfunc.AnonymousFunction6(slice, v1)

	if err != nil {
		fmt.Println("an error has been raised: ", err)
		return
	}
	fmt.Println("original slice :", slice)
	fmt.Println("divided slice :", sliceNew)

}

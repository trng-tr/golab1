package main

import (
	"fmt"
	"golab1/utils"
	"golab1/utils/anonymousfunc"
	"math/rand"
)

func main() {

	v1, err1 := utils.SaisirNumber()
	if err1 != nil {
		fmt.Println("", err1)
		return
	}

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

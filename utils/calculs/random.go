package calculs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetRandomValue() {
	min, err0 := saisirMin()
	if err0 != nil {
		fmt.Printf("Error: %#v\n", err0)
		return
	}
	max, err1 := saisirMax()
	if err1 != nil {
		fmt.Printf("Error: %#v\n", err1)
		return
	}

	if max <= min {
		fmt.Println("Error: le max est inférieur au min")
		return
	}
	fmt.Printf("min:%#v, max %#v", min, max)

}

func saisirMin() (int, error) {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Saisir Min")
	scanner.Scan()
	var min string = scanner.Text()
	minValue, err0 := strconv.Atoi(min)
	if err0 != nil {
		return -1, err0 // get string error
	} else if minValue < 0 {
		return -1, errors.New("error: la valeur du min doit etre supérieur à 0")
	} else {
		return minValue, nil
	}
}

func saisirMax() (int, error) {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Saisir Max")
	scanner.Scan()
	var max string = scanner.Text()
	maxValue, err0 := strconv.Atoi(max)

	if err0 != nil {
		return -1, err0
	} else if maxValue < 0 {
		return -1, errors.New("error: la valeur du max doit etre supérieur à 0")
	} else {
		return maxValue, nil
	}
}

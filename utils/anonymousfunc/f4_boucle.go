package anonymousfunc

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin) // package var

func initSlice(size int, min int, max int) ([]float64, error) {
	var slice []float64 = make([]float64, 0, size)
	for i := 0; i <= size; i++ {
		slice = append(slice, float64(rand.Intn(max-min+1)+min))
	}

	return slice, nil
}

func inputSliceSize() (int, error) {
	fmt.Print("Saisir la taille de la slice: ")
	if !scanner.Scan() {
		return -1, errors.New("error: lecture impossible")
	}
	var str string = scanner.Text()
	if str == "" {
		return -1, errors.New("error: chaine vide")
	}
	size, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return -1, errors.New("error: conversion error")
	} else if size < 10 {
		return -1, errors.New("error: taille slice trop petite")
	}

	return int(size), nil
}

func inputMinMax() (int, int, error) {
	fmt.Print("Saisir min: ")
	if !scanner.Scan() {
		return -1, -1, errors.New("error: lecture impossible")
	}
	var str1 string = scanner.Text()
	if str1 == "" {
		return -1, -1, errors.New("error: min vide")
	}
	min, err := strconv.Atoi(str1)
	if err != nil {
		return -1, -1, errors.New("error conversion")
	}

	fmt.Print("Saisir max: ")
	if !scanner.Scan() {
		return -1, -1, errors.New("error: lecture impossible")
	}
	var str2 string = scanner.Text()
	if str2 == "" {
		return min, -1, errors.New("error: max vide")
	}
	max, err := strconv.Atoi(str2)
	if err != nil {
		return min, -1, errors.New("error conversion")
	}

	if min >= max {
		return -1, -1, errors.New("error: min >=max")
	}

	return min, max, nil
}

func powerN(power float64) ([]float64, error) {
	size, err := inputSliceSize()
	if err != nil {
		return nil, err
	}
	min, max, err := inputMinMax()

	if err != nil {
		return nil, err
	}
	s, err := initSlice(size, min, max)
	if err != nil {
		return nil, err
	}
	fmt.Printf("old slice: %v\n", s)
	for i, v := range s {
		func(v1 float64, n1 float64) {
			v1 = math.Pow(v1, n1)
			s[i] = v1
		}(v, power)
	}
	return s, nil
}

func PrintSlice() {
	fmt.Print("Saisir la puissance :")
	scanner.Scan()
	var str string = scanner.Text()
	if str == "" {
		fmt.Println("error: chaine vide")
		return
	}
	pow, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("error de conversion")
		return
	} else if pow > 5 {
		fmt.Print("la puissance est trop grand")
		return
	}
	newSlice, err := powerN(pow)
	if err != nil {
		fmt.Println("an error has been raised :", err)
		return
	}

	fmt.Printf("new slice: %v\n", newSlice)
}

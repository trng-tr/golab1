package main

import (
	"fmt"
	"golab1/utils/calculs"
)

func main() {
	for { // for ici est une boucle infinie en GO, c'est l'équivalent de while(true)
		var a float64 = saisieValeurA()
		var b float64 = saisieValeurB()
		result1, err1 := calculs.Addition(int(a), int(b)) // on cast des float64 a et b en int64
		if err1 != nil {
			fmt.Printf("une erreur est produite: %s\n", err1)
		} else {
			fmt.Printf("la somme de %d et %d est %d \n", int64(a), int64(b), result1) // on cast des float64 a et b en int64
		}
		result2, err2 := calculs.Division(a, b)
		if err2 != nil {
			fmt.Printf("une erreur de division est produite: %s", err2)
		} else {
			fmt.Printf("la division de %f par %f est %f \n", a, b, result2)
		}

		var v1 float64 = saisieValeurA()
		var v2 float64 = saisieValeurB()
		fmt.Println("la somme des varargs :", calculs.Addition2(a, b, v1, v2))
	}
}
func saisieValeurA() float64 {
	var a float64
	fmt.Print("Saisir la valeur de a:")
	fmt.Scanln(&a)
	return a
}
func saisieValeurB() float64 {
	var b float64
	fmt.Print("Saisir la valeur de b: ")
	fmt.Scanln(&b)
	return b
}

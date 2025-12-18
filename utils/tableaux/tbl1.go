package tableaux

import "fmt"

func declarerTableau() {
	var table [5]int                          // déclarer tableau de  5 nombres
	fmt.Println("tableau des nombres", table) //[0,0,0,0,0]
}

var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

func printTableau1() {

	// récupération de l'index et de la valeur
	for index, jour := range jours {
		fmt.Println(jour, "est le jour numéro", (index + 1), "de la semaine")
	}
}

func printTableau2() {
	// on récupère tous les éléments
	fmt.Println(jours[:]) //[lundi mardi mercredi jeudi vendredi samedi dimanche]
}

func printTableau3() {
	// on récupère tous les éléments jusqu'à l'index 3 exclu
	fmt.Println(jours[:3]) //[lundi mardi mercredi]
}

func printTableau4() {
	// on récupère tous les éléments à partir l'index 3 inclu
	fmt.Println(jours[3:]) // [jeudi vendredi samedi dimanche]
}

func printTableau5() {
	// on récupère les éléments à partir l'index 2 inclu jusqu'à l'index 5 exclu
	fmt.Println(jours[2:5]) // [mercredi jeudi vendredi samedi dimanche]
}

func modifyElementsTable() [7]string { //retourne un tableau de len 5
	for indice, jour := range jours {
		if jour == "samedi" || jour == "dimanche" {
			jours[indice] = "weekend"
		}
	}
	return jours
}

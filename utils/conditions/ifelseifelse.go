package conditions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func SaisirPoids() (int, error) { //public function
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Saisir votre poids :")
	scanner.Scan()
	var poids string = scanner.Text()
	poidsValue, err2 := strconv.Atoi(poids)
	if err2 != nil {
		return -100, errors.New("erreur de conversion")
	} else if !isValidInputPoids(poidsValue) {
		// fmt.Sprintf est l'équivalent de String.format()
		return -100, errors.New("error => Votre poids est < au minimum ou > 250 kg autorisé")
	} else {
		return poidsValue, nil
	}
}

func SaisirAge() (int8, error) { // public function
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Saisir votre age :")
	scanner.Scan() // lancer le scan
	var age string = scanner.Text()
	ageValue, err1 := strconv.Atoi(age)
	if err1 != nil {
		return -1, errors.New("erreur de conversion")
	} else if !isValidInputAge(ageValue) {
		return -1, errors.New("error: votre age est < minimun ou > au maximum")
	} else {
		return int8(ageValue), nil
	}
}
func isValidInputPoids(poids int) bool { // private function
	return poids >= 70 && poids <= 250
}
func isValidInputAge(age int) bool { // private funcion
	return age >= 20 && age <= 40
}

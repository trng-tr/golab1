package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFromKeyboard() (int, string) {
	/* creation d'un scanner pour capturer entre clavier
	depuis le clavier*/
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Saisir quelque chose au clavier:")
	// lancement du scanner
	scanner.Scan()
	// stockage de l'entrée utilisateur dans une variable
	var input string = scanner.Text()
	fmt.Printf("vous avez saisi %#v\n", input)
	/* convertir le input string en int,
	func Atoi(s string) (int, error)*/
	poids, err := strconv.Atoi(input)
	if err != nil {
		return -1, err.Error() // get error as string
	} else if poids <= 30 {
		return -1, "Error: tu es encore jeune"
	} else {
		return poids, ""
	}
}

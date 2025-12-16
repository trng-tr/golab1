package conditions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UsingSwitchCase() {
	choice, err := readInt64()
	if err != "" {
		fmt.Printf("%#v\n", err)
		return
	}
	switch choice { // variable a vérifier
	case 0, 1: // 1 ou 0
		fmt.Print("George Boole !")
	case 7:
		fmt.Println("William Van de Walle!")
	case 23:
		fmt.Println("Pour certains, le nombre 23 est source de nombreux mystères, qu'en penses-tu Jim Carrey?")
	case 42:
		fmt.Println("la réponse à la question ultime du sens de la vie!")
	default:
		println("Mauvais choix !") // Valeur par défaut
	}
}
func readInt64() (int64, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("faire votre choix")
	scanner.Scan()
	var str = scanner.Text()
	if str == "" {
		return -1, "Error: chaine vide"
	}
	choice, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		return -1, err.Error()
	}
	return choice, ""
}

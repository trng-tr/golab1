package constantes

import "fmt"

func PrintConstante() {
	const constante string = "Je suis une constante de type string, je ne peux pas être modifié après affectation"
	fmt.Printf("%s\n", constante)
}

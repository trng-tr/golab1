package main

import (
	"fmt"
	"golab1/utils"
	"golab1/utils/constantes"
	"golab1/utils/variables"
)

func main() {
	fmt.Println("Hello World !")
	DisplayMessage()
	InputName()
	var message string = utils.GetGreeting("Placide Nduwayo")
	fmt.Printf("%s\n", message)
	var autre string
	fmt.Printf("valeur par défaut de la variable autre %#v\n", autre)
	variables.VariablesDeclaration()
	constantes.PrintConstante()
}
func DisplayMessage() {
	fmt.Println("apprendre le langage Go !")
	fmt.Println("le langage Go est fastoche")
	fmt.Println("Ceci est un message custom !")
}
func DisplayCustomMessage(message string) {
	fmt.Printf("message customized %#v\n", message) //
}
func InputName() {
	var firstname, lastname string
	fmt.Print("saisir votre prénom et votre nom sur la même ligne séparer par des espaces:")
	/*pointeur à l'adresse memoire stockant les valeurs
	de firstname et lastname, on les saisit sur la meme
	ligne séparés par des espaces */
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("il/elle s appelle %#v %#v\n", firstname, lastname)
	fmt.Printf("bienvenue %#v %#v\n", firstname, lastname)
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"golab1/utils/wrappers/tobool"
	"os"
)

func main() {
	for {
		str2, err3 := saisieBool()
		if err3 != nil {
			fmt.Printf("%#v\n", err3)
			return
		}

		value3, err4 := tobool.ConvertToBool(str2)
		if err4 != nil {
			fmt.Printf("%#v\n", err4)
			return
		}
		fmt.Printf("Bool value %#v\n", value3)

	}
}
func saisieBool() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("saisir un boolean: ")
	scanner.Scan()
	var value string = scanner.Text()
	if value == "" {
		return "Error =>", errors.New("error: chaine vide")
	}
	return value, nil
}

package utils

import "fmt"

func GetGreeting(name string) string {
	return fmt.Sprintf("Bonjour %v !", name)
}

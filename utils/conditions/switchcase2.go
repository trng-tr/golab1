package conditions

import (
	"fmt"
	"time"
)

func GetWeekEnd() {
	switch time.Now().Weekday() { // recuperer le joude de la semaine
	case time.Saturday:
		fmt.Println("Samedi")
	case time.Sunday:
		fmt.Println("Dimanche")
	default:
		fmt.Println("au boulot, le weekend est fini")
	}
}

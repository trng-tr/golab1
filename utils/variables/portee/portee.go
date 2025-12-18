package portee

import (
	"fmt"
)

func getApplicationInfo(f func()) {
	f()
}

func GetApplicationInfo() {
	getApplicationInfo(func() {
		fmt.Printf("Application name %v, version: %v\n",
			getApplicationName(),
			getApplicationVersion())
	})
}

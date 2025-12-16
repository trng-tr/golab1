package tobool

import (
	"errors"
	"strconv"
)

func ConvertToBool(str string) (bool, error) {
	value, err := strconv.ParseBool(str)
	if err != nil {
		return false, errors.New("erreur de conversion")
	}
	// string reconnues comme true: "1","t","T","true","TRUE","True"
	// string reconnues comme false: "0","f","F","false","FALSE","False"
	// tout autre string retourne une error
	return value, nil
}

package toint

func ConvertToInt() (int, int, error, error) {
	value1, err1 := convertToInt1()
	if err1 != nil {
		return -1, -1, err1, nil
	}

	value2, err2 := convertToInt2()
	if err2 != nil {
		return -1, -1, nil, err2
	}

	return value1, int(value2), nil, nil

}

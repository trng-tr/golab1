package anonymousfunc

import "errors"

func compute3(slice []float64, value float64,
	divide func([]float64, float64) ([]float64, error)) ([]float64, error) {

	newSlice, err := divide(slice, value)
	if err != nil {
		return nil, err
	}
	return newSlice, nil
}

func AnonymousFunction6(slice []float64, value float64) ([]float64, error) {
	newSlice, err := compute3(slice, value, func(s []float64, v float64) ([]float64, error) {
		if v == 0.0 {
			return nil, errors.New("error: il n'est pas encore possible de diviser par zero")
		}
		var result []float64 = make([]float64, 0) // initialiser du slice
		for _, value := range s {
			result = append(result, value/v)
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}

	return newSlice, nil
}

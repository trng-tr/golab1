package calculs

func Addition2(values ...float64) float64 {
	var (
		somme float64 = 0.0
	)

	for _, value := range values {
		somme += value
	}

	return somme
}

package anonymousfunc

import "errors"

/* fonction classique compute() à laquelle on passe un callback
pour effectuer des operation aux parametres v1 et v2*/
func compute(v1 int, v2 int, operation func(int, int) int) int {
	return operation(v1, v2)
}

// calling  anonymous function callback on classic function
func AnonymousFunction4(n1 int, n2 int) int {
	/*le callback va effectuer le produit de v1 et v2
	passés à compute et retourner le resultat*/
	resut := compute(n1, n2, func(i1, i2 int) int {
		return (i1 * i2)
	})

	return resut
}

func compute2(value int, factoriel func(int) (int, error)) (int, error) {
	fact, err := factoriel(value)
	if err != nil {
		return -1, err
	}
	return fact, nil
}

func AnonymousFunction5(nb int) (int, error) {
	result, err := compute2(nb, func(n int) (int, error) {
		if n < 0 {
			return -1, errors.New("error: nombre negatif pour le factoriel")
		} else if n > 20 {
			return -1, errors.New("error: overflow, nombre trop grand pour le calcul du factoriel")

		} else if n == 0 || n == 1 {
			return 1, nil
		}
		var fact int = 1
		for i := 2; i <= n; i++ {
			fact *= i
		}
		return fact, nil
	})
	if err != nil {
		return -1, err
	}

	return result, nil
}

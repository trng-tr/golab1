package calculs

import "errors"

func Addition(a int, b int) (int, error) { /* en Go une fonction peut retourner
	plusieurs valeus: le resultat ou l'erreur */
	if a <= 0 || b <= 0 {
		return -1, errors.New("une des valeurs est <=0")
	}

	return a + b, nil
}

func Division(a float64, b float64) (float64, error) { /*la function publique de division
	retourne deux valeurs soit le resultat soit une erreur */
	if b == 0 {
		return -1, errors.New("impossible de diviser par zero")
	}
	return a / b, nil
}

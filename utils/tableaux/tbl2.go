package tableaux

import "fmt"

const (
	nbLines  int = 3 // nombres de tableaux, chaque ligne est un tableau
	nbColums int = 5 // nb d'element pour chaque ligne
)

func declareMatrix() { // delcare a matrix l lines and c columns
	var matrix [nbLines][nbColums]float64
	fmt.Println(matrix)
}

func initMatrix() [nbLines][nbColums]float64 {
	//initialize matrix
	var matrix [nbLines][nbColums]float64 = [nbLines][nbColums]float64{
		{1, 3, 5, 7, 9}, {0, 2, 4, 6, 8}, {5, 7, 11, 13, 17}}
	return matrix
}

package main

import (
	"fmt"
)

func first(a [3][5]int, b [5][4]int) [3][4]int {
	var x [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for z := 0; z < 5; z++ {
				x[i][j] += a[i][z] * b[z][j]
			}
		}
	}
	return x
}

func main() {
	fmt.Println("Эта прога вычисляет определитель матрицы 3*3")
	aMatrix := [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	bMatrix := [5][4]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	fmt.Println(first(aMatrix, bMatrix))
}

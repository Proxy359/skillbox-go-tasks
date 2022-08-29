package main

import (
	"fmt"
)

func first(a [3][3]int) int {
	x := a[0][0]*a[1][1]*a[2][2] + a[0][2]*a[1][0]*a[2][1] + a[0][1]*a[1][2]*a[2][0]
	x += -a[0][2]*a[1][1]*a[2][0] - a[0][1]*a[1][0]*a[2][2] - a[0][0]*a[1][2]*a[2][1]
	return x
}

func main() {
	fmt.Println("Эта прога вычисляет определитель матрицы 3*3")
	matrix := [3][3]int{
		{1, 2, 3},
		{5, 6, 7},
		{9, 10, 11},
	}
	fmt.Println(first(matrix))
}

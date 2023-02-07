package main

import (
	"fmt"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	r := FindDiagonalOrder(matrix)

	fmt.Println(r)
	t.Log(r)
}

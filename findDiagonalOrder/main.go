package main

import "fmt"

func main() {

	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// matrix := [][]int{
	// 	{1, 2, 3, 10},
	// 	{4, 5, 6, 11},
	// 	{7, 8, 9, 12},
	// }

	matrix := [][]int{}

	rt := FindDiagonalOrder(matrix)

	fmt.Println(rt)

}

// FindDiagonalOrder returns the diagonal
func FindDiagonalOrder(matrix [][]int) []int {

	var data []int
	// 行数
	row := len(matrix)

	if row == 0 {
		return data
	}

	// 列数
	col := len(matrix[0])
	// 最终输出集合数
	count := row + col - 1

	var sum int

	for k := 0; k < count; k++ {
		sum = k

		if k%2 == 0 {
			for j := 0; j < col; j++ {
				i := sum - j

				if i <= (row-1) && i >= 0 {
					data = append(data, matrix[i][j])
				}
			}
		} else {
			for j := (col - 1); j >= 0; j-- {
				i := sum - j

				if i <= (row-1) && i >= 0 {
					data = append(data, matrix[i][j])
				}
			}
		}

	}

	return data
}

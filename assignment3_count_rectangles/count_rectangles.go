package main

import "fmt"

func countRectangle(rectangles [][]int) int {
	nrows, ncols := getArrayShape(rectangles)
	mark_matrix := createCopyArray(rectangles)
	count_rectangles := 0

	for row_idx := 0; row_idx < nrows; row_idx++ {
		for col_idx := 0; col_idx < ncols; col_idx++ {
			// check adjancent area for same rectangle
			if (rectangles[row_idx][col_idx] == 1) {
				belong_to_exist_rect, rectangle_idx := inspectAdjacentCells(mark_matrix, row_idx, col_idx)
				if (belong_to_exist_rect) {
					mark_matrix[row_idx][col_idx] = rectangle_idx
				} else {
					count_rectangles++
					mark_matrix[row_idx][col_idx] = count_rectangles
				}
			}
		}
	}
	return count_rectangles
}

func getArrayShape(arr [][]int) (int, int) {
	nrows := len(arr)
	ncols := len(arr[0])
	return nrows, ncols
}

func createCopyArray(arr [][]int) [][]int{
	nrows, ncols := getArrayShape(arr)
	new_arr := make([][]int, nrows)
	for i := range new_arr {
		new_arr[i] = make([]int, ncols)
	}
	return new_arr
}

func inspectAdjacentCells(mark_matrix [][]int, row_idx, col_idx int) (bool, int) {
	// check cell on top
	if ((row_idx > 0) && (mark_matrix[row_idx-1][col_idx] > 0)){
		return true, mark_matrix[row_idx-1][col_idx]
	}

	// check cell on left
	if ((col_idx > 0) && (mark_matrix[row_idx][col_idx-1] > 0)){
		return true, mark_matrix[row_idx][col_idx-1]
	}
	return false, 0
}

func main() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := countRectangle(arr)
	fmt.Printf("%v", count)
}
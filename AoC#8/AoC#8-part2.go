package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var grid [][]int
var grid_visited [][]bool

func checkDepth(row, col int) {
	if grid_visited[row][col] == false {
		grid_visited[row][col] = true

		if col == 0 {
			if row == 0 {
				if !grid_visited[row][col+1] {
					size++
					checkDepth(row, col+1)
				}
				if !grid_visited[row+1][col] {
					size++
					checkDepth(row+1, col)
				}
			} else if row == len(grid)-1 {
				if !grid_visited[row][col+1] {
					size++
					checkDepth(row, col+1)
				}
				if !grid_visited[row-1][col] {
					size++
					checkDepth(row-1, col)
				}
			} else {
				if !grid_visited[row][col+1] {
					size++
					checkDepth(row, col+1)

				}
				if !grid_visited[row-1][col] {
					size++
					checkDepth(row-1, col)

				}
				if !grid_visited[row+1][col] {
					size++
					checkDepth(row+1, col)

				}
			}
		} else if row == 0 {
			if col == len(grid[row])-1 {
				if !grid_visited[row][col-1] {
					size++
					checkDepth(row, col-1)

				}
				if !grid_visited[row+1][col] {
					size++
					checkDepth(row+1, col)

				}
			} else {
				if !grid_visited[row][col+1] {
					size++
					checkDepth(row, col+1)

				}
				if !grid_visited[row][col-1] {
					size++
					checkDepth(row, col-1)

				}
				if !grid_visited[row+1][col] {
					size++
					checkDepth(row+1, col)

				}
			}
		} else if col == len(grid[row])-1 {
			if row == len(grid)-1 {
				if !grid_visited[row][col-1] {
					size++
					checkDepth(row, col-1)

				}
				if !grid_visited[row-1][col] {
					size++
					checkDepth(row-1, col)

				}
			} else {
				if !grid_visited[row][col-1] {
					size++
					checkDepth(row, col-1)

				}
				if !grid_visited[row+1][col] {
					size++
					checkDepth(row+1, col)

				}
				if !grid_visited[row-1][col] {
					size++
					checkDepth(row-1, col)

				}
			}
		} else if row == len(grid)-1 {
			if !grid_visited[row][col-1] {
				size++
				checkDepth(row, col-1)

			}
			if !grid_visited[row][col+1] {
				size++
				checkDepth(row, col+1)

			}
			if !grid_visited[row-1][col] {
				size++
				checkDepth(row-1, col)

			}
		} else {
			if !grid_visited[row][col-1] {
				size++
				checkDepth(row, col-1)

			}
			if !grid_visited[row][col+1] {
				size++
				checkDepth(row, col+1)

			}
			if !grid_visited[row-1][col] {
				size++
				checkDepth(row-1, col)

			}
			if !grid_visited[row+1][col] {
				size++
				checkDepth(row+1, col)

			}
		}
	}
}

var size int

func checkAdjacent() []int {
	basins_size := []int{}
	size = 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			checkDepth(i, j)
			if size > 1 {
				basins_size = append(basins_size, size)
			}
			size = 1
		}
	}

	return basins_size
}

func main() {
	grid = make([][]int, 0)
	grid_visited = make([][]bool, 0)
	var row string

	for {
		_, err := fmt.Scanf("%s", &row)

		if err != nil {
			break
		}

		row_split := strings.Split(row, "")

		row_slice := []int{}
		row_slice_bool := []bool{}
		for _, v := range row_split {
			v_int, _ := strconv.Atoi(v)
			row_slice = append(row_slice, v_int)
			if v_int == 9 {
				row_slice_bool = append(row_slice_bool, true)
			} else {
				row_slice_bool = append(row_slice_bool, false)
			}
		}

		grid = append(grid, row_slice)
		grid_visited = append(grid_visited, row_slice_bool)
	}

	basins_size := checkAdjacent()
	// fmt.Println(checkAdjacent(grid))
	sort.Ints(basins_size)
	fmt.Println(basins_size[len(basins_size)-3:])

	res := 1
	for _, v := range basins_size[len(basins_size)-3:] {
		res *= v
	}
	fmt.Println(res)
}

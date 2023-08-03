package main

import "fmt"

func countRectangles(rec [][]int) int {
	count := 0
	for i := 0; i < len(rec); i++ {
		for j := 0; j < len(rec[i]); j++ {
			if rec[i][j] == 1 {
				count++
				markRectangleAsVisited(rec, i, j)
			}
		}
	}
	return count
}

func markRectangleAsVisited(rec [][]int, i int, j int) {
	if i < 0 || i >= len(rec) || j < 0 || j >= len(rec[i]) || rec[i][j] == 0 {
		return
	}

	rec[i][j] = 0
	markRectangleAsVisited(rec, i+1, j)
	markRectangleAsVisited(rec, i-1, j)
	markRectangleAsVisited(rec, i, j+1)
	markRectangleAsVisited(rec, i, j-1)
}

func main() {
	arr := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	count := countRectangles(arr)
	fmt.Println(count)
}

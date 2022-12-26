package main

import (
	"bufio"
	"fmt"
	"os"
)

const n = 99 // side of a input square

func readInput() [n][n]int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var arr [n][n]int
	for row := 0; row < n; row++ {
		scanner.Scan()
		line := scanner.Text()
		for i := 0; i < n; i++ {
			arr[row][i] = int(line[i]) - 48
		}
	}
	return arr
}

func checkNeighbors(arr [n][n]int, row int, col int) bool {
	cell := arr[row][col]
	for i := row + 1; i < n; i++ {
		if arr[i][col] >= cell {
			goto Bottom
		}
	}
	return true
Bottom:
	for i := row - 1; i >= 0; i-- {
		if arr[i][col] >= cell {
			goto Right
		}
	}
	return true
Right:
	for i := col + 1; i < n; i++ {
		if arr[row][i] >= cell {
			goto Left
		}
	}
	return true
Left:
	for i := col - 1; i >= 0; i-- {
		if arr[row][i] >= cell {
			return false
		}
	}
	return true
}

func countNonVisible(arr [n][n]int) (visible int) {
	visible = 4*n - 4
	for row := 1; row < n-1; row++ {
		for col := 1; col < n-1; col++ {
			if checkNeighbors(arr, row, col) {
				visible++
			}
		}
	}
	return
}

func main() {
	data := readInput()
	fmt.Println(countNonVisible(data))
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countDiff() int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	text := scanner.Text()
	const n = 4
	start := 0
	diff := 0
Loop:
	for k := range text {
		for i := start; i < start+diff; i++ {
			if text[i] == text[k] {
				start = k
				diff = 1
				continue Loop
			}
		}
		if diff == n-1 {
			return k
		}
		diff++
	}
	return 0
}
func main() {
	fmt.Println(countDiff())
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func game() int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	point := 0
	for scanner.Scan() {
		text := scanner.Text()
		fst := text[0]
		snd := text[2]
		if fst == 'A' {
			if snd == 'X' {
				point += 4
			} else if snd == 'Y' {
				point += 8
			} else if snd == 'Z' {
				point += 3
			}
		} else if fst == 'B' {
			if snd == 'X' {
				point += 1
			} else if snd == 'Y' {
				point += 5
			} else if snd == 'Z' {
				point += 9
			}
		} else if fst == 'C' {
			if snd == 'X' {
				point += 7
			} else if snd == 'Y' {
				point += 2
			} else if snd == 'Z' {
				point += 6
			}
		}
	}
	return point
}
func main() {
	fmt.Println(game())
}

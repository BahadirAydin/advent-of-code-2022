package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func draw(clock, reg int) {

	if reg == clock%40 || reg-1 == clock%40 || reg+1 == clock%40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if clock%40 == 0 {
		fmt.Println()
	}
}
func simulate() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	reg := 1
	clock := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		clock++
		draw(clock, reg)
		str := scanner.Text()
		if str[0] == 'a' { //addx
			clock++
			n, _ := strconv.Atoi(str[5:])
			reg += n
			draw(clock, reg)
		}
	}
}

func main() {
	simulate()
}

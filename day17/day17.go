package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int64
}
type Rock struct {
	stones     []Point
	leftmost   int
	rightmost  int
	bottommost int64
}

var rock1 = Rock{
	stones: []Point{
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
	},
	leftmost:   2,
	rightmost:  5,
	bottommost: 0,
}

var rock2 = Rock{
	stones: []Point{
		{2, 1},
		{3, 0},
		{3, 1},
		{3, 2},
		{4, 1},
	},
	leftmost:   2,
	rightmost:  4,
	bottommost: 0,
}

var rock3 = Rock{
	stones: []Point{
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
	},
	leftmost:   2,
	rightmost:  4,
	bottommost: 0,
}

var rock4 = Rock{
	stones: []Point{
		{2, 0},
		{2, 1},
		{2, 2},
		{2, 3},
	},
	leftmost:   2,
	rightmost:  2,
	bottommost: 0,
}

var rock5 = Rock{
	stones: []Point{
		{2, 0},
		{2, 1},
		{3, 0},
		{3, 1},
	},
	leftmost:   2,
	rightmost:  3,
	bottommost: 0,
}

var rocks = []Rock{rock1, rock2, rock3, rock4, rock5}
var num_iter_p1 = 2022
var num_iter_p2 = 1000000000000

func parse() (wind []int) {
	f, _ := os.ReadFile("input.txt")
	str := string(f)
	val := 0
	for i := 0; i < len(str)-1; i++ {
		if str[i] == '>' {
			val = 1
		} else if str[i] == '<' {
			val = -1
		}
		wind = append(wind, val)
	}
	return wind
}

type Chamber struct {
	arr     [][]bool
	highest [7]int64
}

func initChamber() *Chamber {
	arr := make([][]bool, 1000)
	for i := range arr {
		arr[i] = make([]bool, 7)
	}
	return &Chamber{arr, [7]int64{-1, -1, -1, -1, -1, -1, -1}}
}

func max(arr [7]int64) int64 {
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func min(arr [7]int64) int64 {
	min := arr[0]
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	if min == -1 {
		return 0
	}
	return min
}

func isHit(chamber *Chamber, rock *Rock, point Point) bool {
	if (rock.bottommost == 0 && point.y == -1) || (rock.leftmost == 0 && point.x == -1) || (rock.rightmost == 6 && point.x == 1) {
		return true
	}
	for stone := range rock.stones {
		if chamber.arr[rock.stones[stone].y+point.y][rock.stones[stone].x+point.x] {
			return true
		}
	}
	return false
}

func resizeChamber(chamber *Chamber, newHeight int) {
	if newHeight >= len(chamber.arr) {
		newArr := make([][]bool, newHeight+1)
		copy(newArr, chamber.arr)
		for i := len(chamber.arr); i <= newHeight; i++ {
			newArr[i] = make([]bool, 7)
		}
		chamber.arr = newArr
	}
}

func day17(wind []int, n int) int64 {
	windPos := 0
	chamber := initChamber()
	offset := int64(0)
	i := 0
	for i < n {
		for originalRock := range rocks {
			if i == n {
				break
			}
			i++
			spawn_y := max(chamber.highest) + 4
			if spawn_y+5 >= int64(len(chamber.arr)) {
				resizeChamber(chamber, len(chamber.arr)*2)
			}
			rock := rocks[originalRock]
			rock.stones = make([]Point, len(rocks[originalRock].stones))
			copy(rock.stones, rocks[originalRock].stones)
			for stone := range rock.stones {
				rock.stones[stone].y = spawn_y + rock.stones[stone].y
			}
			rock.bottommost = spawn_y
			for {

				currentWind := int64(wind[windPos])
				if !isHit(chamber, &rock, Point{currentWind, 0}) {
					for stone := range rock.stones {
						rock.stones[stone].x += currentWind
					}
					rock.leftmost += int(currentWind)
					rock.rightmost += int(currentWind)
				}

				windPos = (windPos + 1) % len(wind)

				isHit := isHit(chamber, &rock, Point{0, -1})
				if isHit || rock.bottommost == 0 {
					for stone := range rock.stones {
						y, x := rock.stones[stone].y, rock.stones[stone].x
						chamber.arr[y][x] = true
						if y > chamber.highest[x] {
							chamber.highest[x] = y
						}

					}
					break
				}
				// bottom
				rock.bottommost--
				for stone := range rock.stones {
					rock.stones[stone].y--
				}
			}
			for stone := range rock.stones {
				y, x := rock.stones[stone].y, rock.stones[stone].x
				chamber.arr[y][x] = true
				if y > chamber.highest[x] {
					chamber.highest[x] = y
				}
			}
			// take out everything after min(chamber.highest)

		}
		min_ := min(chamber.highest)
		chamber.arr = chamber.arr[min_:]
		offset += min_
		for i := range chamber.highest {
			chamber.highest[i] -= min_
		}

	}
	// adding 1 because of index starting at 0
	return max(chamber.highest) + 1 + offset
}

func main() {
	wind := parse()
	fmt.Println(day17(wind, num_iter_p1))
	fmt.Println(day17(wind, num_iter_p2))
}

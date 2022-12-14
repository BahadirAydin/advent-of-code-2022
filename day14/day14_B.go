package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Point struct {
  x int
  y int
}
type Line struct {
  head Point
  tail Point
}

func readData() ([]Line, int, int, int) {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  lines := make([]Line, 0)
  minX := 500
  maxX := 500
  maxY := 0
  for scanner.Scan() {
    text := scanner.Text()
    list := strings.Split(text, " -> ")
    f := make([]Point, 0)
    for i := 0; i < len(list); i++ {
      split := strings.Split(list[i], ",")
      x, _ := strconv.Atoi(split[0])
      y, _ := strconv.Atoi(split[1])
      f = append(f, Point{x, y})
      if x < minX {
        minX = x
      } else if x > maxX {
        maxX = x
      }
      if y > maxY {
        maxY = y
      }
    }
    for i := 0; i < len(f)-1; i++ {
      lines = append(lines, Line{f[i], f[i+1]})
    }
  }
  return lines, minX, maxX + 1, maxY + 1
}
func createGrid(x, y int) [][]int {
  grid := make([][]int, y)
  for i := 0; i < len(grid); i++ {
    grid[i] = make([]int, x)
  }
  return grid
}
func drawLine(l Line, grid [][]int, offset, add int) {
  head := l.head
  tail := l.tail
  if head.x == tail.x {
    start, end := head.y, tail.y
    if tail.y < start {
      start, end = tail.y, head.y
    }
    for i := start; i <= end; i++ {
      grid[i][head.x-offset+add] = 'w'
    }
  } else {
    start, end := head.x, tail.x
    if tail.x < start {
      start, end = tail.x, head.x
    }
    for i := start; i <= end; i++ {
      grid[head.y][i-offset+add] = 'w'
    }
  }
}
func fillGrid(lines []Line, grid [][]int, offset, add int) {
  for i := 0; i < len(lines); i++ {
    drawLine(lines[i], grid, offset, add)
  }
}
func simulateSand(grid [][]int, x, maxX, maxY int) bool {
  y := 0
  for {
    if y == maxY || x == maxX {
      return false
    }
    if y+1 < maxY && grid[y+1][x] == 0 {
      y++
    } else if y+1 < maxY && x-1 >= 0 && grid[y+1][x-1] == 0 {
      x--
      y++
    } else if y+1 < maxY && x+1 < maxX && grid[y+1][x+1] == 0 {
      x++
      y++
    } else if y+1 >= maxY || x-1 < 0 || x+1 > maxX {
      return false
    } else {
      break
    }
  }
  grid[y][x] = 'o'
  return true
}
func sandCount() int {
  lines, minX, maxX, maxY := readData()
  offset := minX
  maxX -= offset
  const add = 400 // 400 is not very close to infinity but it works :)
  maxY += 2
  maxX += add
  minX = 0
  grid := createGrid(maxX, maxY)
  for i := 0; i < maxX; i++ {
    grid[maxY-1][i] = 'w'
  }
  fillGrid(lines, grid, offset, add/2)
  count := 0
  spawnPosition := 500 - offset + add/2
  for simulateSand(grid, spawnPosition, maxX, maxY) {
    count++
    if grid[0][spawnPosition] == 'o' {
      break
    }
  }
  return count
}
func main() {
  fmt.Println(sandCount())
}

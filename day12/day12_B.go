package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
)

const row = 41
const col = 138

func readData() ([row][col]int32, []int, []int, int, int) {
  f, _ := os.Open("input.txt")
  var arr [row][col]int32
  defer f.Close()
  scanner := bufio.NewScanner(f)
  starty := make([]int, 0)
  startx := make([]int, 0)
  r := 0
  endy, endx := 0, 0
  for scanner.Scan() {
    for c, v := range scanner.Text() {
      if v == 'E' {
        endy, endx = r, c
        v = 'z'
      } else if v == 'S' {
        v = 'a'
      }
      if v == 'a' {
        starty = append(starty, r)
        startx = append(startx, c)
      }
      arr[r][c] = v
    }
    r++
  }
  return arr, starty, startx, endy, endx
}
func allPaths() int {
  data, starty, startx, endy, endx := readData()
  min := math.MaxInt32
  for k := range starty {
    v := int(shortestPath(data, starty[k], startx[k], endy, endx))
    if v < min {
      min = v
    }
  }
  return min
}
func shortestPath(data [row][col]int32, starty, startx, endy, endx int) int32 {
  var distances [row][col]int32
  var visited [row][col]bool
  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      distances[i][j] = math.MaxInt32
    }
  }
  y, x := starty, startx
  distances[y][x] = 0
  for !(y == endy && x == endx) {
    visited[y][x] = true
    if y+1 < row && !visited[y+1][x] && data[y+1][x]-data[y][x] <= 1 {
      if distances[y][x]+1 < distances[y+1][x] {
        distances[y+1][x] = distances[y][x] + 1
      }
    }
    if x+1 < col && !visited[y][x+1] && data[y][x+1]-data[y][x] <= 1 {
      if distances[y][x]+1 < distances[y][x+1] {
        distances[y][x+1] = distances[y][x] + 1
      }
    }
    if x-1 >= 0 && !visited[y][x-1] && data[y][x-1]-data[y][x] <= 1 {
      if distances[y][x]+1 < distances[y][x-1] {
        distances[y][x-1] = distances[y][x] + 1
      }
    }
    if y-1 >= 0 && !visited[y-1][x] && data[y-1][x]-data[y][x] <= 1 {
      if distances[y][x]+1 < distances[y-1][x] {
        distances[y-1][x] = distances[y][x] + 1
      }
    }
    min := math.MaxInt32 // ideally i should use a heap here, this is not the best implementation
    for i := 0; i < row; i++ {
      for j := 0; j < col; j++ {
        if distances[i][j] < int32(min) && !visited[i][j] {
          min = int(distances[i][j])
          y, x = i, j
        }
      }
    }
    if visited[y][x] { // stuck
      return math.MaxInt32
    }
  }
  return distances[y][x]
}
func main() {
  fmt.Println(allPaths())
}

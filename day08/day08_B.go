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

func countScore(arr [n][n]int, row int, col int) int {
  cell := arr[row][col]
  scoreTop := 0
  for i := row + 1; i < n; i++ {
    if arr[i][col] >= cell {
      scoreTop++
      break
    }
    scoreTop++
  }
  scoreBottom := 0
  for i := row - 1; i >= 0; i-- {
    if arr[i][col] >= cell {
      scoreBottom++
      break
    }
    scoreBottom++
  }
  scoreRight := 0
  for i := col + 1; i < n; i++ {
    if arr[row][i] >= cell {
      scoreRight++
      break
    }
    scoreRight++
  }

  scoreLeft := 0
  for i := col - 1; i >= 0; i-- {
    if arr[row][i] >= cell {
      scoreLeft++
      break
    }
    scoreLeft++
  }
  return scoreBottom * scoreLeft * scoreRight * scoreTop
}

func countNonVisible(arr [n][n]int) (max int) {
  // one problem is that my code does not check the edge trees. in some extreme cases the edge cases could be answer too.
  max = 0
  for row := 1; row < n-1; row++ {
    for col := 1; col < n-1; col++ {
      score := countScore(arr, row, col)
      if score > max {
        max = score
      }
    }
  }
  return
}

func main() {
  data := readInput()
  fmt.Println(countNonVisible(data))
}

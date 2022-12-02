package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func findMax() int {

  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  val := 0
  max := 0
  for scanner.Scan() {
    if scanner.Text() == "" {
      if val > max {
        max = val
      }
      val = 0
    } else {
      converted, _ := strconv.Atoi(scanner.Text())
      val += converted
    }
  }
  return max
}

func main() {
  max := findMax()
  fmt.Println(max)
}

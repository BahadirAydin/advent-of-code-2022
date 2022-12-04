package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func countContained() int {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  count := 0
  for scanner.Scan() {
    text := scanner.Text()
    dash := strings.Index(text, "-")
    comma := strings.Index(text, ",")
    x1, _ := strconv.Atoi(text[0:dash])
    x2, _ := strconv.Atoi(text[dash+1 : comma])
    text = text[comma+1:]
    dash = strings.Index(text, "-")
    y1, _ := strconv.Atoi(text[0:dash])
    y2, _ := strconv.Atoi(text[dash+1:])
    if (x1 >= y1 && x2 <= y2) || (y1 >= x1 && y2 <= x2) {
      count++
    }
  }
  return count
}
func main() {
  fmt.Println(countContained())
}

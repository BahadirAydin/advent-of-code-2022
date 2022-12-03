package main

import (
  "bufio"
  "fmt"
  "os"
)

func rucksack() int {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  point := 0
  for scanner.Scan() {
    text := scanner.Text()
    length := len(text)
    for i := 0; i < length/2; i++ {
      item := text[i]
      found := false
      for j := length / 2; j < length; j++ {
        if item == text[j] {
          ascii_val := int(item)
          if ascii_val > 96 {
            point += int(item) - 96
          } else {
            point += int(item) - 38
          }
          found = true
          break
        }
      }
      if found {
        break
      }
    }
  }
  return point
}
func main() {
  fmt.Println(rucksack())
}

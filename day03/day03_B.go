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
  var group [3]string
  count := 0
  for scanner.Scan() {
    text := scanner.Text()
    group[count] = text
    count++
    if count == 3 {
    Loop:
      for i := 0; i < len(group[0]); i++ {
        for j := 0; j < len(group[1]); j++ {
          for k := 0; k < len(group[2]); k++ {
            if group[0][i] == group[1][j] && group[0][i] == group[2][k] && group[1][j] == group[2][k] {
              ascii_val := int(group[0][i])
              if ascii_val > 96 {
                point += int(group[0][i]) - 96
              } else {
                point += int(group[0][i]) - 38
              }
              break Loop
            }
          }
        }
      }
      count = 0
    }
  }
  return point
}
func main() {
  fmt.Println(rucksack())
}

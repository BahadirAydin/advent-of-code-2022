package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func findMax3() int {

  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  val := 0
  var top3 [3]int
  for scanner.Scan() {
    if scanner.Text() == "" {
      for k, v := range top3 {
        if val > v {
          if k == 0 {
            tmp := top3[1]
            top3[1] = top3[0]
            top3[2] = tmp
          } else if k == 1 {
            top3[2] = top3[1]
          }
          top3[k] = val
          break
        }
      }
      val = 0
    } else {
      converted, _ := strconv.Atoi(scanner.Text())
      val += converted
    }
  }
  return top3[0] + top3[1] + top3[2]
}

func main() {
  max := findMax3()
  fmt.Println(max)
}

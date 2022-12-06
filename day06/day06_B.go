package main

import (
  "bufio"
  "fmt"
  "os"
)

func countDiff() int {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  text := scanner.Text()
  const n = 14
  var arr [n - 1]int32
  diff := 0
Loop:
  for k, v := range text {
    for i := 0; i < diff; i++ {
      if arr[i] == v {
        arr[0] = v
        diff = 1
        continue Loop
      }
    }
    if diff == n-1 {
      return k
    }
    arr[diff] = v
    diff++
  }
  return 0
}
func main() {
  fmt.Println(countDiff())
}

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func reorderTo(arr *[]byte, no int) {
  for i := 0; i < no; i++ {
    (*arr) = append([]byte{0}, (*arr)...)
  }
}
func reorderFrom(arr *[]byte, no int) {
  (*arr) = (*arr)[no:]
}

func day5() {
  crates := 9
  arr := make([][]byte, crates)
  arr[0] = []byte{'W', 'L', 'S'}
  arr[1] = []byte{'Q', 'N', 'T', 'J'}
  arr[2] = []byte{'J', 'F', 'H', 'C', 'S'}
  arr[3] = []byte{'B', 'G', 'N', 'W', 'M', 'R', 'T'}
  arr[4] = []byte{'B', 'Q', 'H', 'D', 'S', 'L', 'R', 'T'}
  arr[5] = []byte{'L', 'R', 'H', 'F', 'V', 'B', 'J', 'M'}
  arr[6] = []byte{'M', 'J', 'N', 'R', 'W', 'D'}
  arr[7] = []byte{'J', 'D', 'N', 'H', 'F', 'T', 'Z', 'B'}
  arr[8] = []byte{'T', 'F', 'B', 'N', 'Q', 'L', 'H'}

  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    text := scanner.Text()
    index := strings.Index(text, "f")
    no, _ := strconv.Atoi(text[5 : index-1])
    from := int(text[index+5]) - 49
    to := int(text[index+10]) - 49
    if no > len(arr[from]) {
      no = len(arr[from])
    }
    reorderTo(&arr[to], no)
    for i := 0; i < no; i++ {
      arr[to][i] = arr[from][i]
    }
    reorderFrom(&arr[from], no)
  }
  for i := 0; i < crates; i++ {
    fmt.Print(string(arr[i][0]))
  }
  fmt.Println()
}
func main() {
  day5()
}

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func signalPower(clock, reg int) int {
  if clock == 20 || clock == 60 || clock == 100 || clock == 140 || clock == 180 || clock == 220 {
    return clock * reg
  }
  return 0
}

func simulate() (sum int) {
  f, _ := os.Open("input.txt")
  defer f.Close()
  reg := 1
  clock := 0
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    clock++
    sum += signalPower(clock, reg)
    str := scanner.Text()
    if str[0] == 'a' { //addx
      clock++
      sum += signalPower(clock, reg)
      n, _ := strconv.Atoi(str[5:])
      reg += n
    }
  }
  return
}

func main() {
  fmt.Println(simulate())
}

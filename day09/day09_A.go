package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

type Point struct {
  x int
  y int
}

// I have first implemented BUNCH OF if else statements keeping track of each
// possible move and state which was a literal pain.
// then I realized unless the absolute length between two points is
// larger than sqrt(2) which is what happens when they're diagonal
// we can also say that this is a condition to "touch"
// the position DOES NOT change.
// and when the position changes (len > sqrt(2)) it is just the head's old position

// IMPORTANT: There is a more sophisticated solution in part B which solves all ropes with -> length > 1.

// Finds the change in tails position.
func findTailPos(oldHead, head, tail Point) (int, int) {
  hx, hy, tx, ty, ox, oy := head.x, head.y, tail.x, tail.y, oldHead.x, oldHead.y
  len := (hx-tx)*(hx-tx) + (hy-ty)*(hy-ty)
  if len <= 2 {
    return tx, ty
  }
  return ox, oy
}

func day9() {
  f, _ := os.Open("input.txt")
  defer f.Close()
  arr := make(map[Point]bool, 0)
  var tail Point
  var head Point
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    line := scanner.Text()
    dir := line[0]
    n, _ := strconv.Atoi(line[2:])
    for i := 0; i < n; i++ {
      oldHead := head
      if dir == 'R' {
        head.x += 1
      } else if dir == 'L' {
        head.x -= 1
      } else if dir == 'U' {
        head.y += 1
      } else {
        head.y -= 1
      }
      tail.x, tail.y = findTailPos(oldHead, head, tail)
      _, ok := arr[tail]
      if !ok {
        arr[tail] = true
      }
    }
  }
  fmt.Println(len(arr))
}

func main() {
  day9()
}

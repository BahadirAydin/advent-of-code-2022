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

// New addition in part B is that now the head can move diagonal.

// Finds the change in tails position.
func findTailPos(oldHead, head, tail Point) (int, int) {
  hx, hy, tx, ty, ox, oy := head.x, head.y, tail.x, tail.y, oldHead.x, oldHead.y
  lenHeadTail := (hx-tx)*(hx-tx) + (hy-ty)*(hy-ty)
  lenHeadOldhead := (hx-ox)*(hx-ox) + (hy-oy)*(hy-oy)

  if lenHeadTail == 4 { //new mechanic: head goes diagnoal but the tail does not. check the examples in AOC website
    return (tx + hx) / 2, (ty + hy) / 2
  } else if lenHeadOldhead == 2 && lenHeadTail > 2 { // new mechanic: head goes diagonal and tail too. check the examples in AOC website
    return tx + hx - ox, ty + hy - oy
  } else if lenHeadTail <= 2 {
    return tx, ty
  }
  return ox, oy
}

func day9() {
  f, _ := os.Open("input.txt")
  defer f.Close()
  arr := make(map[Point]bool, 0)
  const knot = 10
  var knots [knot]Point
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := scanner.Text()
    dir := line[0]
    n, _ := strconv.Atoi(line[2:])
    for i := 0; i < n; i++ {
      oldHead := knots[0]
      if dir == 'R' {
        knots[0].x += 1
      } else if dir == 'L' {
        knots[0].x -= 1
      } else if dir == 'U' {
        knots[0].y += 1
      } else {
        knots[0].y -= 1
      }
      for i := 1; i < knot; i++ {
        cpy := knots[i]
        knots[i].x, knots[i].y = findTailPos(oldHead, knots[i-1], knots[i])
        oldHead = cpy
      }
      _, ok := arr[knots[knot-1]]
      if !ok {
        arr[knots[knot-1]] = true
      }
    }
  }
  fmt.Println(len(arr))
}
func main() {
  day9()
}

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
  hx, hy, tx, ty := head.x, head.y, tail.x, tail.y
  len := (hx-tx)*(hx-tx) + (hy-ty)*(hy-ty)
  len2 := (hx-oldHead.x)*(hx-oldHead.x) + (hy-oldHead.y)*(hy-oldHead.y)
  if len == 4 { //new mechanic head goes diagnoal but the tail does not. check the examples in AOC website
    return (tx + hx) / 2, (ty + hy) / 2
  } else if len2 == 2 && len > 2 { // new mechanic head goes diagonal and tail too. check the examples in AOC website
    return tx + (hx - oldHead.x), ty + (hy - oldHead.y)
  } else if len <= 2 {
    return tx, ty
  }
  return oldHead.x, oldHead.y
}

func day9() {
  f, _ := os.Open("input.txt")
  defer f.Close()
  arr := make(map[Point]bool, 0)
  const knot = 9
  var knots [knot]Point
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
      knotCopy := knots[0]
      knots[0].x, knots[0].y = findTailPos(oldHead, head, knots[0])
      for i := 1; i < knot; i++ {
        cpy := knots[i]
        knots[i].x, knots[i].y = findTailPos(knotCopy, knots[i-1], knots[i])
        knotCopy = cpy
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

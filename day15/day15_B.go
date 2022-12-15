package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

type Point struct {
  x int
  y int
}
type Sensor struct {
  beacon Point
  x      int
  y      int
}

func contains(arr []Point, p Point) bool {
  for _, v := range arr {
    if v == p {
      return true
    }
  }
  return false
}
func calculateDistance(p Point, s Sensor) int {
  dis := math.Abs(float64(p.x-s.x)) + math.Abs(float64(p.y-s.y))
  return int(dis)
}
func pointRange(s Sensor, y, distance int) (int, int, bool) {
  d := distance - int(math.Abs(float64((s.y - y))))
  if d < 0 {
    return 0, 0, false
  }
  return s.x - d, d + s.x, true
}
func intersection(p1 *Point, p2 *Point) (res Point, flag bool) {
  if p1.x >= p2.x && p1.y <= p2.y {
    res = *p2
  } else if p2.x >= p1.x && p2.y <= p1.y { // vice versa
    res = *p1
  } else if p1.y >= p2.x && p1.x <= p2.x {
    cp2 := Point{p1.y + 1, p2.y}
    cp1 := Point{p1.x, p2.x - 1}
    res = Point{p2.x, p1.y}
    *p1 = cp1
    *p2 = cp2
  } else if p2.y >= p1.x && p2.x <= p1.x {
    cp2 := Point{p2.y + 1, p1.y}
    cp1 := Point{p2.x, p1.x - 1}
    res = Point{p1.x, p2.y}
    *p1 = cp1
    *p2 = cp2
  } else {
    flag = true
  }
  flag = !flag
  return
}
func readData() []Sensor {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  sensors := make([]Sensor, 0)
  for scanner.Scan() {
    text := scanner.Text()
    comma := strings.Index(text, ",")
    colon := strings.Index(text, ":")
    sx, _ := strconv.Atoi(text[12:comma])
    sy, _ := strconv.Atoi(text[comma+4 : colon])
    text = text[colon+25:]
    comma = strings.Index(text, ",")
    bx, _ := strconv.Atoi(text[:comma])
    by, _ := strconv.Atoi(text[comma+4:])
    sensors = append(sensors, Sensor{beacon: Point{bx, by}, x: sx, y: sy})
  }
  return sensors
}
func split(arr []Point) []Point {
  // repeatedly find the intersections between the elements
  // make that intersection a new element
  // so if there are 2 elements intersecting with another
  // after the intersection there are 3 elements which are all distinct.
  // if one of them is contained in the other simply remove the other.
  // as a result all elements will be distinct.
  for {
    flag := true
    for i := 0; i < len(arr)-1; i++ {
      for j := i + 1; j < len(arr); j++ {
        x, ok := intersection(&arr[i], &arr[j])
        if ok {
          flag = false
          if x == arr[i] {
            arr[j], arr[len(arr)-1] = arr[len(arr)-1], arr[j]
            arr = arr[:len(arr)-1]
          } else if x == arr[j] {
            arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
            arr = arr[:len(arr)-1]
          } else {
            arr = append(arr, x)
          }
        }
      }
    }
    if flag {
      break
    }
  }
  return arr
}
func simulate(Y int) []Point {
  sensors := readData()
  arr := make([]Point, 0) // x field means "from", y field means "to" in this context
  beaconSet := make([]Point, 0)
  for _, s := range sensors {
    dis := calculateDistance(s.beacon, s)
    if !contains(beaconSet, s.beacon) {
      beaconSet = append(beaconSet, s.beacon)
    }
    from, to, ok := pointRange(s, Y, dis)
    if from > to {
      from, to = to, from
    }
    if ok && !(from < 0 && to < 0) {
      if from < 0 {
        from = 0
      }
      arr = append(arr, Point{from, to})
    }
  }
  arr = split(arr)
  return arr
}
func findTuning() int {
  const min = 0
  const max = 4000000
  for i := min; i <= max; i++ {
    arr := simulate(i)
    for j := min; j <= max; j++ {
      flag := true
      for _, v := range arr {
        if v.x == j {
          flag = false
          j = v.y
        }
      }
      if flag {
        return 4000000*j + i
      }
    }
  }
  return -1
}
func main() {
  fmt.Println(findTuning())
}

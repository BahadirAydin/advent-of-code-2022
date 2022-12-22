package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Valve struct {
  id       int
  name     string
  flowRate int
  edges    []*Valve
}

func readData() []Valve {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  arr := make([]Valve, 0)
  edges := make([][]string, 0)
  id := 0
  for scanner.Scan() {
    text := strings.Fields(scanner.Text())
    name := text[1]
    rate, _ := strconv.Atoi(text[4][5 : len(text[4])-1])
    neighbors := make([]string, 0)
    for i := 9; i < len(text); i++ {
      curr := text[i]
      if curr[len(curr)-1] == ',' {
        curr = curr[:len(curr)-1]
      }
      neighbors = append(neighbors, curr)
    }
    arr = append(arr, Valve{id, name, rate, []*Valve{}})
    id++
    edges = append(edges, neighbors)
  }
  for i := 0; i < len(edges); i++ {
    c := edges[i]
    for _, item := range c {
      for k, v := range arr {
        if v.name == item {
          arr[i].edges = append(arr[i].edges, &arr[k])
        }
      }
    }
  }
  return arr
}
func contains(arr []int, v int) bool {
  for i := 0; i < len(arr); i++ {
    if arr[i] == v {
      return true
    }
  }
  return false
}

func shortestPath(valves []Valve, source int) []int {
  queue := make([]int, 0)
  queue = append(queue, source)
  visited := make([]int, 0)
  visited = append(visited, source)
  distances := make([]int, len(valves))
  distances[source] = 1
  for len(queue) > 0 {
    top := queue[0]
    queue = queue[1:]
    neighbors := valves[top].edges
    for _, v := range neighbors {
      if !contains(visited, v.id) {
        queue = append(queue, v.id)
        visited = append(visited, v.id)
        distances[v.id] = distances[top] + 1
      }
    }
  }
  return distances
}
func createDistanceMatrix(valves []Valve) [][]int {
  distanceMatrix := make([][]int, 0)
  for i := 0; i < len(valves); i++ {
    distanceMatrix = append(distanceMatrix, shortestPath(valves, i))
  }
  return distanceMatrix
}

type Info struct {
  index1, index2, minute1, minute2 int
}

func day16() int {
  valves := readData()
  s := 0
  for k := range valves {
    if valves[k].name == "AA" {
      s = k
      break
    }
  }
  matrix := createDistanceMatrix(valves)
  opened := make([]int, 0)
  for k := range valves {
    if valves[k].flowRate == 0 {
      opened = append(opened, k)
    }
  }
  table := make(map[Info]int)
  x := dfs(valves, matrix, opened, s, s, 0, 26, 26, &table)
  return x
}

func memorized(table *map[Info]int, index1, index2, minute1, minute2 int) int {
  val1, ok1 := (*table)[Info{index1, index2, minute1, minute2}]
  val2, ok2 := (*table)[Info{index2, index1, minute2, minute1}]
  if ok1 {
    return val1
  } else if ok2 {
    return val2
  }
  return -1
}

func dfs(valves []Valve, matrix [][]int, opened []int, index1, index2, totalPressure, minute1, minute2 int, table *map[Info]int) int {
  if val := memorized(table, index1, index2, minute1, minute2); totalPressure < val {
    return val
  }
  opened = append(opened, index1)
  opened = append(opened, index2)
  row1 := matrix[index1]
  row2 := matrix[index2]
  max := 0
  res := 0

  for i := 0; i < len(row1); i++ {
    for j := 0; j < len(row2); j++ {
      if i == j {
        continue
      }
      flag1 := contains(opened, i)
      flag2 := contains(opened, j)
      if row1[i] < minute1 && row2[j] < minute2 && !flag1 && !flag2 {
        addedPressure := (minute1-row1[i])*valves[i].flowRate + (minute2-row2[j])*valves[j].flowRate
        res = dfs(valves, matrix, opened, i, j, totalPressure+addedPressure, minute1-row1[i], minute2-row2[j], table)
      } else if minute1 >= row1[i] && row2[j] < minute2 && !flag2 {
        addedPressure := (minute2 - row2[j]) * valves[j].flowRate
        res = dfs(valves, matrix, opened, index1, j, totalPressure+addedPressure, minute1, minute2-row2[j], table)
      } else if minute2 >= row2[j] && row1[i] < minute1 && !flag1 {
        addedPressure := (minute1 - row1[i]) * valves[i].flowRate
        res = dfs(valves, matrix, opened, i, index2, totalPressure+addedPressure, minute1-row1[i], minute2, table)
      }
      if res > max {
        max = res
      }
    }
  }
  if max == 0 {
    return totalPressure
  }
  _, ok1 := (*table)[Info{index1, index2, minute1, minute2}]
  _, ok2 := (*table)[Info{index2, index1, minute2, minute1}]
  if ok1 {
    (*table)[Info{index1, index2, minute1, minute2}] = max
  } else if ok2 {
    (*table)[Info{index2, index1, minute2, minute1}] = max
  } else {
    (*table)[Info{index1, index2, minute1, minute2}] = max
  }
  return max
}
func main() {
  fmt.Println(day16())
}

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
  x := dfs(valves, matrix, opened, s, 0, 30)
  return x
}

func dfs(valves []Valve, matrix [][]int, opened []int, index, totalPressure, minutes int) int {
  if minutes != 30 {
    opened = append(opened, index)
  }
  row := matrix[index]
  max := 0
  for k := range row {
    if row[k] <= minutes && !contains(opened, k) {
      res := dfs(valves, matrix, opened, k, totalPressure+(minutes-row[k])*valves[k].flowRate, minutes-row[k])
      if res > max {
        max = res
      }
    }
  }
  if max == 0 {
    return totalPressure
  }
  return max
}
func main() {
  fmt.Println(day16())
}

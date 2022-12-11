package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Monkey struct {
  id          int
  items       []int
  opAddition  bool // if true operation is addition else it is multiplication
  opN         int  // IF it is -1, it means the "old" value which is the item itself.
  test        int  // after 	doing the operation on item and dividing it by 3 check if it is divisible by this variable
  trueMonkey  int  // throw id if test is true
  falseMonkey int  // throw id if test is false
  inspect     int
}

func readInput() []Monkey {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  monkeys := make([]Monkey, 0)
  for scanner.Scan() {
    text := scanner.Text()
    if len(text) > 0 && text[0] == 'M' { // Monkey
      var newMonkey Monkey
      // get id
      id, _ := strconv.Atoi(text[7 : len(text)-1])
      newMonkey.id = id
      // get items
      scanner.Scan()
      text = scanner.Text()
      items := strings.Split(text[18:], ", ")
      for _, v := range items {
        conv, _ := strconv.Atoi(v)
        newMonkey.items = append(newMonkey.items, conv)
      }
      // get if addition or multiplication and the operand
      scanner.Scan()
      text = scanner.Text()
      fields := strings.Fields(text[23:])
      if fields[0] == "+" {
        newMonkey.opAddition = true // since it is default false, we don't need an else statement
      }
      if fields[1] == "old" {
        newMonkey.opN = -1
      } else {
        conv, _ := strconv.Atoi(fields[1])
        newMonkey.opN = conv
      }
      // get test number (divisible by ....)
      scanner.Scan()
      text = scanner.Text()
      test, _ := strconv.Atoi(text[21:])
      newMonkey.test = test
      // get monkey ids in the last two lines
      scanner.Scan()
      text = scanner.Text()
      m1, _ := strconv.Atoi(text[29:])
      scanner.Scan()
      text = scanner.Text()
      m2, _ := strconv.Atoi(text[30:])
      newMonkey.trueMonkey = m1
      newMonkey.falseMonkey = m2
      // append the monkey to the list
      monkeys = append(monkeys, newMonkey)
    }
  }
  return monkeys
}

func monkeyBehaviour(m *Monkey, itemIndex int, allDivs int) int {
  m.inspect++
  itemVal := m.items[itemIndex]
  if m.opAddition {
    itemVal += m.opN
  } else {
    if m.opN == -1 {
      itemVal *= itemVal
    } else {
      itemVal *= m.opN
    }
  }
  itemVal %= allDivs
  m.items[itemIndex] = itemVal
  if itemVal%m.test == 0 {
    return m.trueMonkey
  }
  return m.falseMonkey
}

func round(monkeys *[]Monkey, allDivs int) {
  for k := range *monkeys {
    if len((*monkeys)[k].items) == 0 {
      continue
    }
    for i := range (*monkeys)[k].items {
      id := monkeyBehaviour(&(*monkeys)[k], i, allDivs)
      (*monkeys)[id].items = append((*monkeys)[id].items, (*monkeys)[k].items[i])
    }
    (*monkeys)[k].items = nil
  }
}
func multiplyMax2(monkeys []Monkey) int {
  var max2 [2]int
  for _, v := range monkeys {
    if v.inspect > max2[0] {
      max2[1] = max2[0]
      max2[0] = v.inspect
    } else if v.inspect > max2[1] {
      max2[1] = v.inspect
    }
  }
  return max2[0] * max2[1]
}
func simulate() int {
  monkeys := readInput()
  const roundN = 10000
  allDivs := 1
  // we will do val = val % allDivs everytime to avoid the numbers getting too big for computer to handle.
  for _, v := range monkeys {
    allDivs *= v.test
  }
  for i := 0; i < roundN; i++ {
    round(&monkeys, allDivs)
  }
  return multiplyMax2(monkeys)
}

func main() {
  fmt.Println(simulate())
}

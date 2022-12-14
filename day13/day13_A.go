package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func compareNum(l, r int) int {
	if l < r {
		return 1
	} else if l == r {
		return 0
	}
	return -1
}
func compare(l, r []interface{}) int {
	size := len(l)
	if size > len(r) {
		size = len(r)
	}
	for i := 0; i < size; i++ {
		itemL, checkL := l[i].([]interface{})
		itemR, checkR := r[i].([]interface{})
		res := 0
		if checkL && checkR { // both are arrays
			res = compare(itemL, itemR)
		} else if checkL {
			rval, _ := r[i].(int)
			res = compare(itemL, []interface{}{rval})
		} else if checkR {
			lval := l[i].(int)
			res = compare([]interface{}{lval}, itemR)
		} else {
			lval, _ := l[i].(int)
			rval, _ := r[i].(int)
			res = compareNum(lval, rval)
		}
		if res != 0 {
			return res
		}
	}
	if len(l) == len(r) {
		return 0
	} else if size == len(r) {
		return -1
	}
	return 1
}

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func parseLine(str string) []interface{} {
	var list []interface{}
	var sym Stack
	var s Stack
	for k := 0; k < len(str); k++ {
		v := str[k]
		if v == '[' {
			s.Push(k)
			sym.Push(int(v))
		} else if v == ']' {
			sym.Pop()
			ind := s.Pop()
			if len(sym) == 0 {
				list = append(list, parseLine(str[ind+1:k]))
			}
		} else if len(sym) == 0 && v != ',' {
			i := k
			for ; i < len(str) && str[i] != ']' && str[i] != ','; i++ {
			}
			conv, _ := strconv.Atoi(str[k:i])
			list = append(list, conv)
			k = i
		}
	}
	return list
}

func compute() (sum int) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	pair := 0
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		pair++
		scanner.Scan()
		r := scanner.Text()
		l = l[1 : len(l)-1]
		r = r[1 : len(r)-1]
		parsedL := parseLine(l)
		parsedR := parseLine(r)
		res := compare(parsedL, parsedR)
		if res == 1 {
			sum += pair
		}
	}
	return sum
}
func main() {
	fmt.Println(compute())
}

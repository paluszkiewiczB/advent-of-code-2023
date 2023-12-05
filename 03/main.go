package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("part one: %s\n", partOne())
	//fmt.Printf("part two: %s\n", partTwo())
}

type window struct {
	val [3]row // value is either dot, symbol, or number
}

type row []any

func (r row) String() string {
	sb := strings.Builder{}
	for _, v := range r {
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}

func newWindow() *window {
	return &window{val: [3]row{}}
}

func (s *window) slide(v row) {
	s.val[0], s.val[1], s.val[2] = s.val[1], s.val[2], v
}

type pos struct {
	x1, x2 int
}
type dot struct {
	pos
}

func (d dot) String() string {
	return "."
}

type symbol struct {
	pos
	val rune
}

func (s symbol) String() string {
	return string(s.val)
}

type number struct {
	pos
	val int
}

func (n number) String() string {
	return strconv.Itoa(n.val)
}

type puzzle struct {
	w   *window
	sum int
}

func (p puzzle) String() string {
	sb := strings.Builder{}
	sb.WriteString("===========\n")
	for _, r := range p.w.val {
		sb.WriteString(fmt.Sprintf("%v\n", r))
	}
	sb.WriteString("===========\n")
	return sb.String()
}

func newPuzzle() *puzzle {
	return &puzzle{
		w: newWindow(),
	}
}

func (p *puzzle) loadLine(line string) {
	vals := parseLine(line)
	p.w.slide(vals)
}

func (p *puzzle) digest() {
	for _, val := range p.w.val[1] {
		if n, ok := val.(number); ok {
			if p.isAdjacent(n) {
				p.sum += n.val
			} else {
			}
		}
	}
}

func debugPrint(p puzzle, n number) {
	show := func(a any, overlap bool) {
		switch t := a.(type) {
		case dot:
			if t.x1 >= n.x1-1 && t.x1 <= n.x2+1 {
				print(t.String())
			}
		case symbol:
			if t.x1 >= n.x1-1 && t.x1 <= n.x2+1 {
				print(t.String())
			}
		case number:
			if t == n {
				print(n.String())
				return
			}
			if !(t.x1 >= n.x1-1 && t.x1 <= n.x2+1) {
				return
			}
			str := strconv.Itoa(t.val)
			left := max(n.x1-1, t.x2)
			right := min(n.x2+1, t.x1)
			if left > right {
				left, right = right, left
			}
			if !overlap {
				if left <= t.x1-1 && right >= t.x1-1 { // is on the left side
					print(str[t.x1-1])
					print(n.String())
				} else if left <= t.x2+1 && right >= t.x2+1 { // is on the right side
					print(n.String())
					print(str[t.x2+1])
				}
			} else {
				trimL := left - t.x1
				trimR := t.x2 - right
				trimmed := str[trimL : len(str)-trimR]
				print(trimmed)
			}

		}
	}

	for _, up := range p.w.val[0] {
		show(up, true)
	}

	println()

	for _, mid := range p.w.val[1] {
		show(mid, false)
	}

	println()

	for _, down := range p.w.val[2] {
		show(down, true)
	}

	println()

}

func (p *puzzle) isAdjacent(n number) bool {
	for _, up := range p.w.val[0] {
		if s, ok := up.(symbol); ok {
			if s.x1 >= n.x1-1 && s.x1 <= n.x2+1 {
				return true
			}
		}
	}

	for _, mid := range p.w.val[1] {
		if s, ok := mid.(symbol); ok {
			if s.x1 == n.x1-1 || s.x1 == n.x2+1 {
				return true
			}
		}
	}

	for _, down := range p.w.val[2] {
		if s, ok := down.(symbol); ok {
			if s.x1 >= n.x1-1 && s.x1 <= n.x2+1 {
				return true
			}
		}
	}

	return false
}

func parseLine(line string) row {
	var vals row
	from, to := -1, -1
	num := 0
	addNum := func() {
		if from == -1 {
			return
		}
		vals = append(vals, number{pos: pos{x1: from, x2: to}, val: num})
		from = -1
		to = -1
		num = 0
	}
	for x, c := range line {
		switch c {
		case '.':
			addNum()
			vals = append(vals, dot{pos: pos{x1: x}})
		default:
			if unicode.IsDigit(c) {
				if from == -1 {
					from = x
					to = x
					num = int(c - '0')
				} else {
					to = x
					num = num*10 + int(c-'0')
				}
			} else {
				addNum()
				vals = append(vals, symbol{pos: pos{x1: x}, val: c})
			}
		}
	}

	if from != -1 {
		addNum()
	}

	return vals
}

func partOne() string {
	s, cF := readFile("input.txt")
	defer cF()

	p := newPuzzle()
	for s.Scan() {
		must(s.Err())
		p.loadLine(s.Text())
		p.digest()
	}
	p.loadLine("") // slide window
	p.digest()

	return strconv.Itoa(p.sum)
}

func partTwo() string {
	s, cF := readFile("sample-input.txt")
	defer cF()

	for s.Scan() {
		must(s.Err())
		line := s.Text()
		println(line)
	}

	return "todo"
}

func readFile(file string) (s *bufio.Scanner, cF func() error) {
	f, err := os.Open(file)
	must(err)
	s = bufio.NewScanner(f)
	return s, f.Close
}

func mustf(f func() error) {
	err := f()
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

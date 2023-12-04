package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part one: %s\n", partOne())
	fmt.Printf("part two: %s\n", partTwo())
}

type game struct {
	id int
	// one element is a set; key is color, value is count
	cubes []map[string]int
}

func (g *game) UnmarshalText(text []byte) error {
	idS, idE := bytes.IndexRune(text, ' '), bytes.IndexRune(text, ':')
	id, err := strconv.Atoi(string(text[idS+1 : idE]))
	if err != nil {
		return err
	}
	g.id = id

	sets := strings.Split(string(text[idE+2:]), ";")
	g.cubes = make([]map[string]int, len(sets))
	for i := range g.cubes {
		g.cubes[i] = make(map[string]int)
	}
	for si, set := range sets {
		if set[0] == ' ' {
			set = set[1:]
		}
		colors := strings.Split(set, ", ")
		for _, color := range colors {
			withValue := strings.Split(color, " ")
			if len(withValue) != 2 {
				return fmt.Errorf("invalid color, two parts expected: %s", color)
			}

			count, err := strconv.Atoi(withValue[0])
			if err != nil {
				return fmt.Errorf("invalid count: %s, %w", withValue[0], err)
			}

			g.cubes[si][withValue[1]] = count
		}
	}

	return nil
}

var rules = map[string]int{"red": 12, "green": 13, "blue": 14}

func partOne() string {
	s, cF := readFile("input.txt")
	defer cF()

	var sum int
	for s.Scan() {
		must(s.Err())
		g := &game{}
		err := g.UnmarshalText(s.Bytes())
		must(err)
		if valid(g) {
			sum += g.id
		}
	}

	return strconv.Itoa(sum)
}

func valid(g *game) bool {
	for _, cube := range g.cubes {
		for color, count := range cube {
			if count > rules[color] {
				return false
			}
		}
	}

	return true
}

func partTwo() string {
	s, cF := readFile("input.txt")
	defer cF()

	var sum int
	for s.Scan() {
		must(s.Err())
		g := &game{}
		err := g.UnmarshalText(s.Bytes())
		must(err)
		sum += calculatePower(g)
	}

	return strconv.Itoa(sum)
}

func calculatePower(g *game) int {
	mins := findMinSet(g)
	power := 1
	for _, minCount := range mins {
		power *= minCount
	}

	return power
}

func findMinSet(g *game) map[string]int {
	mins := make(map[string]int, len(g.cubes))
	for _, cube := range g.cubes {
		for color, count := range cube {
			if curr, ok := mins[color]; !ok || count > curr {
				mins[color] = count
			}
		}
	}

	return mins
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

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("part one: %s\n", partOne())
	fmt.Printf("part two: %s\n", partTwo())
}

func partOne() string {
	s, cF := readFile("sample-input.txt")
	defer cF()

	for s.Scan() {
		must(s.Err())
		line := s.Text()
		println(line)
	}

	return "todo"
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

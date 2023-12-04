package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Printf("part one: %s\n", partOne())
	fmt.Printf("part two: %s\n", partTwo())
}

func partOne() string {
	c := make(chan string)
	go readInput(c, "input.txt")

	var sum int
	for s := range c {
		num := parseLine(s)
		sum += num
	}

	return strconv.Itoa(sum)
}

func parseLine(s string) int {
	digits := make([]rune, 0, 2)
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		}
	}

	n, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
	if err != nil {
		panic(fmt.Sprintf("not a number: %v", digits))
	}

	fmt.Printf("read number: %d from line: %s\n", n, s)
	return n
}

func partTwo() string {
	c := make(chan string)
	go readInput(c, "sample-input.txt")
	for s := range c {
		println(s)
	}

	return "todo"
}

func readInput(c chan string, fileName string) {
	file, err := os.Open(fileName)
	must(err)
	defer mustf(file.Close)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	close(c)
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

	log.Printf("read number: %d from line: %s", n, s)
	return n
}

func partTwo() string {
	c := make(chan string)
	go readInput(c, "input.txt")
	var sum int
	for s := range c {
		num := parseLineTwo(s)
		sum += num
	}

	return strconv.Itoa(sum)
}

var (
	digits = map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func parseLineTwo(s string) int {
	var (
		fi, li = math.MaxInt, -1
		f, l   int
	)

	for code, value := range digits {
		index := strings.Index(s, code)
		if index != -1 && index < fi {
			fi = index
			f = value
		}

		last := strings.LastIndex(s, code)
		if last != -1 && last > li {
			li = last
			l = value
		}
	}

	fmt.Printf("read number: %d from line: %s\n", 10*f+l, s)
	return 10*f + l
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("part one: %s\n", partOne())
	fmt.Printf("part two: %s\n", partTwo())
}

func partOne() string {
	s, cF := readFile("input.txt")
	defer cF()

	total := 0
	for s.Scan() {
		must(s.Err())
		line := s.Text()
		pts := countCardPoints(line)
		total += pts
	}

	return strconv.Itoa(total)
}

func countWonNumbers(line string) int {
	numbers := strings.SplitN(line, ":", 2)[1]
	split := strings.Split(numbers, " | ")
	winningS, myS := split[0], split[1]
	winning, my := parseNumbers(winningS), parseNumbers(myS)
	won := intersection(winning, my)
	return len(won)
}

func countCardPoints(line string) int {
	won := countWonNumbers(line)
	if won == 0 {
		return 0
	}

	if won == 1 {
		return 1
	}

	return 2 << (won - 2)
}

func intersection(a, b []int) []int {
	var common []int
	for _, n := range a {
		for _, m := range b {
			if n == m {
				common = append(common, n)
			}
		}
	}
	return common
}

func parseNumbers(numbers string) []int {
	var nums []int
	for _, n := range strings.Fields(numbers) {
		n, err := strconv.Atoi(n)
		must(err)
		nums = append(nums, n)
	}
	return nums
}

func partTwo() string {
	s, cF := readFile("input.txt")
	defer cF()

	instances := 0
	boost := make([]int, 10_000) // hack to avoid out-of-bound
	cardNo := 1
	for s.Scan() {
		fmt.Printf("card no: %d\n", cardNo)
		cardNo++
		must(s.Err())
		line := s.Text()
		wonCards := countWonNumbers(line)
		fmt.Printf("won cards: %d\n", wonCards)
		copies := boost[0] + 1
		fmt.Printf("add instances: %d\n", copies)
		instances += copies
		fmt.Printf("copies of current card: %d\n", copies)
		boost = boost[1:]
		for i := 0; i < wonCards; i++ {
			boost[i] += copies
		}
		fmt.Printf("%v\n", boost[:10])
	}

	return strconv.Itoa(instances)
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

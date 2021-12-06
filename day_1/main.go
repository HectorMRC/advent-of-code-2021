package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	InputPath = "./day_1/input.txt"
)

func HowManyIncreases(input []int, from int) int {
	if input == nil || len(input) < 2 || from < 0 || from > len(input)-2 {
		return 0
	}

	if input[from] < input[from+1] {
		return 1 + HowManyIncreases(input, from+1)
	}

	return HowManyIncreases(input, from+1)
}

func WindowedHowManyIncreases(input []int, from int, wind int) int {
	if input == nil || len(input) < 2 || from < 0 || from > len(input)-2 {
		return 0
	}

	before := VecSum(input, from, wind)
	totalIncreases := 0
	for it := from + 1; it < len(input); it++ {
		current := VecSum(input, it, wind)
		if before < current {
			totalIncreases++
		}

		before = current
	}

	return totalIncreases
}

func VecSum(in []int, start int, l int) (total int) {
	if len(in) <= start {
		return 0
	}

	for index, i := range in[start:] {
		if index >= l {
			return
		}

		total += i
	}

	return
}

func VecAtoi(in []string) (out []int, err error) {
	out = make([]int, len(in))
	for i, s := range in {
		out[i], err = strconv.Atoi(s)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	r, err := os.Open(InputPath)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	input, err := VecAtoi(strings.Split(string(data), "\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("how many increases = %v\n", HowManyIncreases(input, 0))
	fmt.Printf("windowed how many increases = %v\n", WindowedHowManyIncreases(input, 0, 3))
}

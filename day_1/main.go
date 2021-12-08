package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	common "github.com/HectorMRC/advent-of-code-2021"
)

const (
	InputPath = "./day_1/input.txt"
	WindSize  = 3
	BufSize   = 5
)

func HowManyIncreases(r io.Reader, base int) int {
	if r == nil {
		return 0
	}

	buf := make([]byte, BufSize)
	l, err := r.Read(buf)
	if l == 0 || err != nil {
		return 0
	}

	current, err := strconv.Atoi(string(buf[:l]))
	if err != nil {
		log.Print(err)
		return 0
	}

	if 0 < base && current > base {
		return 1 + HowManyIncreases(r, current)
	} else {
		return HowManyIncreases(r, current)
	}
}

func WindowedHowManyIncreases(r io.Reader, wind int, base int) (count int) {
	if r == nil {
		return 0
	}

	var window []int
	buf := make([]byte, BufSize)

	for l, err := r.Read(buf); l > 0 && err == nil; l, err = r.Read(buf) {
		current, err := strconv.Atoi(string(buf[:l]))
		if err != nil {
			log.Print(err)
			return 0
		}

		b0 := base
		base += current

		if window = append(window, current); len(window) <= wind {
			continue
		}

		if base -= window[0]; 0 < b0 && base > b0 {
			count++
		}

		window = window[1:]
	}

	return
}

func main() {
	file, err := os.Open(InputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := common.NewReader(file)
	//fmt.Printf("total increases: %v\n", HowManyIncreases(r, 0))
	fmt.Printf("total increases (window size %v): %v\n", WindSize, WindowedHowManyIncreases(r, WindSize, 0))
}

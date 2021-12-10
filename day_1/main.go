package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	inputPath = "./day_1/input.txt"
	windSize  = 3
)

func HowManyIncreases(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	return howManyIncreases(s, 0)
}

func howManyIncreases(s *bufio.Scanner, base int) int {
	if s == nil || !s.Scan() {
		return 0
	}

	current, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Print(err)
		return 0
	}

	if 0 < base && current > base {
		return 1 + howManyIncreases(s, current)
	} else {
		return howManyIncreases(s, current)
	}
}

func WindowedHowManyIncreases(r io.Reader, wind int) (count int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	return windowedHowManyIncreases(s, wind, 0)
}

func windowedHowManyIncreases(s *bufio.Scanner, wind int, base int) (count int) {
	if s == nil {
		return 0
	}

	var window []int
	for s.Scan() {
		current, err := strconv.Atoi(s.Text())
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
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	fmt.Printf("total increases: %v\n", HowManyIncreases(tee))
	fmt.Printf("total increases (window size %v): %v\n", windSize, WindowedHowManyIncreases(&buf, windSize))
}

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
	InputPath = "./day_3/input.txt"
	BufSize   = 12
	BitMask   = 0b111111111111
)

func gammaRate(s *bufio.Scanner, base []int) {
	if !s.Scan() {
		return
	}

	for i, b := range s.Text() {
		if b == '1' {
			base[i] += 1
		} else {
			base[i] -= 1
		}
	}

	gammaRate(s, base)
}

func GammaRate(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	count := make([]int, BufSize)
	gammaRate(s, count)

	var str string
	for _, i := range count {
		if i >= 0 {
			str += "1"
		} else {
			str += "0"
		}
	}

	return strconv.ParseInt(str, 2, BufSize+1)
}

func oxygenGeneratorRating(bytes []string, index int) string {
	if l := len(bytes); l == 1 {
		return bytes[0]
	} else if l < 1 {
		return ""
	}

	var zeros []string
	var ones []string

	for _, item := range bytes {
		if item[index] == '0' {
			zeros = append(zeros, item)
		} else {
			ones = append(ones, item)
		}
	}

	if len(ones) >= len(zeros) {
		return oxygenGeneratorRating(ones, index+1)
	} else {
		return oxygenGeneratorRating(zeros, index+1)
	}
}

func co2ScrubberRating(bytes []string, index int) string {
	if l := len(bytes); l == 1 {
		return bytes[0]
	} else if l < 1 {
		return ""
	}

	var zeros []string
	var ones []string

	for _, item := range bytes {
		if item[index] == '0' {
			zeros = append(zeros, item)
		} else {
			ones = append(ones, item)
		}
	}

	if len(zeros) <= len(ones) {
		return co2ScrubberRating(zeros, index+1)
	} else {
		return co2ScrubberRating(ones, index+1)
	}
}

func LifeSupoprtRating(r io.Reader) (int64, int64, error) {
	var zeros []string
	var ones []string

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if item := s.Text(); item[0] == '0' {
			zeros = append(zeros, item)
		} else {
			ones = append(ones, item)
		}
	}

	var oxygen string
	if len(ones) >= len(zeros) {
		oxygen = oxygenGeneratorRating(ones, 1)
	} else {
		oxygen = oxygenGeneratorRating(zeros, 1)
	}

	var co2 string
	if len(zeros) <= len(ones) {
		co2 = co2ScrubberRating(zeros, 1)
	} else {
		co2 = co2ScrubberRating(ones, 1)
	}

	gen, err := strconv.ParseInt(oxygen, 2, BufSize+1)
	if err != nil {
		return 0, 0, err
	}

	scr, err := strconv.ParseInt(co2, 2, BufSize+1)
	if err != nil {
		return 0, 0, err
	}

	return gen, scr, nil
}

func main() {
	f, err := os.Open(InputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	gamma, err := GammaRate(tee)
	if err != nil {
		log.Fatal(err)
	}

	epsilon := gamma ^ BitMask
	fmt.Printf("%v (gamma rate) x %v (epsilon rate): %v\n", gamma, epsilon, gamma*epsilon)

	oxygen, co2, err := LifeSupoprtRating(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v (oxygen generator rate) x %v (CO2 scrubber rate): %v\n", oxygen, co2, oxygen*co2)
}

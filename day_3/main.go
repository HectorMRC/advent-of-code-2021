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
		} else if b == '0' {
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
}

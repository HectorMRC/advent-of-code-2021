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
	InputPath = "./day_3/input.txt"
	BufSize   = 12
	BitMask   = 0b111111111111
)

func rate(r io.Reader, base []int) {
	buf := make([]byte, BufSize)
	l, err := r.Read(buf)
	if l == 0 || err != nil {
		return
	}

	for i, b := range buf {
		if b == "1"[0] {
			base[i] += 1
		} else if b == "0"[0] {
			base[i] -= 1
		}
	}

	rate(r, base)
}

func GammaRate(r io.Reader) (s int64, err error) {
	count := make([]int, BufSize)
	rate(r, count)

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
	file, err := os.Open(InputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := common.NewReader(file)
	gamma, err := GammaRate(r)
	if err != nil {
		log.Fatal(err)
	}

	epsilon := gamma ^ BitMask
	fmt.Printf("%v (gamma rate) x %v (epsilon rate): %v\n", gamma, epsilon, gamma*epsilon)
}

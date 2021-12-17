package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputPath     = "./day_6/input.txt"
	Separator     = ","
	shortLifetime = 7
	longLifetime  = 9
)

type frame [longLifetime]int

func howManyFishes(init frame, days int) (total int) {
	if days == 0 {
		for _, fishes := range init {
			total += fishes
		}

		return
	}

	var next frame
	for index, fishes := range init {
		if index == 0 {
			next[shortLifetime-1] += fishes
			next[longLifetime-1] += fishes
		} else {
			next[index-1] += fishes
		}
	}

	return howManyFishes(next, days-1)
}

func HowManyFishes(r io.Reader, days int) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	if !s.Scan() {
		return 0
	}

	var init frame
	for _, lifetime := range strings.Split(s.Text(), Separator) {
		index, err := strconv.Atoi(lifetime)
		if err != nil {
			log.Fatal(err)
		}

		init[index] += 1
	}

	return howManyFishes(init, days)
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	fmt.Printf("how many fishes after 80 days: %v\n", HowManyFishes(tee, 80))
	fmt.Printf("how many fishes after 256 days: %v\n", HowManyFishes(&buf, 256))
}

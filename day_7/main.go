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
	"sync"
)

const (
	inputPath = "./day_7/input.txt"
	Separator = ","
)

func fuelRequired(in []int, target int, constant bool) (sum int) {
	for _, origin := range in {
		diff := origin - target
		if diff < 0 {
			diff = -diff
		}

		if constant {
			sum += diff
			continue
		}

		for i := 0; i < diff; i++ {
			sum += i + 1
		}
	}

	return
}

func MinimumFuelRequired(r io.Reader, constant bool) (fuel int, err error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	if !s.Scan() {
		return
	}

	input := strings.Split(s.Text(), Separator)
	positions := make([]int, len(input))
	for i, str := range input {
		if positions[i], err = strconv.Atoi(str); err != nil {
			log.Fatal(err)
		}
	}

	var min int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, target := range positions {
		wg.Add(1)

		go func(wg *sync.WaitGroup, target int) {
			defer wg.Done()
			required := fuelRequired(positions, target, constant)

			if min > required || min == 0 {
				mu.Lock()
				defer mu.Unlock()

				if min > required || min == 0 {
					min = required
				}
			}
		}(&wg, target)
	}

	wg.Wait()
	return min, nil
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	min, err := MinimumFuelRequired(tee, true)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("minimum fuel required (constant consumption): %v\n", min)

	min, err = MinimumFuelRequired(&buf, false)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("minimum fuel required (incremental consumption): %v\n", min)
}

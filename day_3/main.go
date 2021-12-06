package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	InputPath = "./day_3/input.txt"
	BitSize   = 12
	BitMask   = 0b111111111111
)

func GammaRate(input []string, index int) (gamma string) {
	if input == nil || len(input) == 0 || len(input) <= index {
		return
	}

	counter := make(map[uint8]int)
	for _, entry := range input {
		counter[entry[index]]++
	}

	max := -1
	for key, value := range counter {
		if value > max {
			gamma = string(key)
			max = value
		}
	}

	return
}

func BitGammaRate(input []string, bitSize int) (gamma string) {
	buff := make([]string, bitSize)

	var wg sync.WaitGroup
	wg.Add(bitSize)

	for it := 0; it < bitSize; it++ {
		var current int = it
		go func(wg *sync.WaitGroup, index int) {
			buff[index] = GammaRate(input, index)
			wg.Done()
		}(&wg, current)
	}

	wg.Wait()
	return strings.Join(buff, "")
}

func GroupByOccurrences(input []string, index int) (counter map[uint8]int, filter map[uint8][]string) {
	if input == nil || len(input) == 0 || index < 0 || index >= BitSize {
		return
	}

	counter = make(map[uint8]int)
	filter = make(map[uint8][]string)
	for _, item := range input {
		counter[item[index]]++

		if _, exists := filter[item[index]]; exists {
			filter[item[index]] = append(filter[item[index]], item)
		} else {
			filter[item[index]] = []string{item}
		}
	}

	return
}

func OxygenGeneratorRating(input []string, index int) (rate string) {
	if len(input) == 1 {
		return input[0]
	}

	counter, filter := GroupByOccurrences(input, index)

	max := -1
	var selected uint8
	for key, value := range counter {
		if value > max || (value == max && key == "1"[0]) {
			selected = key
			max = value
		}
	}

	return OxygenGeneratorRating(filter[selected], index+1)
}

func Co2ScrubberRating(input []string, index int) (rate string) {
	if len(input) == 1 {
		return input[0]
	}

	counter, filter := GroupByOccurrences(input, index)

	min := len(input) + 1
	var selected uint8
	for key, value := range counter {
		if value < min || (value == min && key == "0"[0]) {
			selected = key
			min = value
		}
	}

	return Co2ScrubberRating(filter[selected], index+1)
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

	input := strings.Split(string(data), "\n")
	gamma, err := strconv.ParseInt(BitGammaRate(input, BitSize), 2, BitSize+1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("gamma x epsilon = %v\n", gamma*(gamma^BitMask))

	gen, err := strconv.ParseInt(OxygenGeneratorRating(input, 0), 2, BitSize+1)
	if err != nil {
		panic(err)
	}

	scrub, err := strconv.ParseInt(Co2ScrubberRating(input, 0), 2, BitSize+1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("oxygen x CO2 = %v\n", gen*scrub)
}

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	InputPath = "./day_2/input.txt"
)

func Location(movements []string, from int) (h, v int) {
	if movements == nil || len(movements) == 0 || from < 0 || len(movements) <= from {
		return
	}

	movement := strings.Split(movements[from], " ")
	if got := len(movement); got != 2 {
		panic(fmt.Errorf("got %v in line %v, want %v", got, from, 2))
	}

	steps, err := strconv.Atoi(movement[1])
	if err != nil {
		panic(err)
	}

	switch movement[0] {
	case "up":
		v = -steps
	case "down":
		v = steps
	case "forward":
		h = steps
	}

	h1, v1 := Location(movements, from+1)
	return h + h1, v + v1
}

func AimedLocation(movements []string, from, aim, h, v int) (int, int) {
	if movements == nil || len(movements) == 0 || from < 0 || len(movements) <= from {
		return h, v
	}

	movement := strings.Split(movements[from], " ")
	if got := len(movement); got != 2 {
		panic(fmt.Errorf("got %v in line %v, want %v", got, from, 2))
	}

	steps, err := strconv.Atoi(movement[1])
	if err != nil {
		panic(err)
	}

	switch movement[0] {
	case "up":
		aim -= steps
	case "down":
		aim += steps
	case "forward":
		h += steps
		v += aim * steps
	}

	return AimedLocation(movements, from+1, aim, h, v)
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
	h0, v0 := Location(input, 0)

	fmt.Printf("location = %v\n", h0*v0)

	h1, v1 := AimedLocation(input, 0, 0, 0, 0)
	fmt.Printf("aimed location = %v\n", h1*v1)
}

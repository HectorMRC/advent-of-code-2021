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
	inputPath = "./day_2/input.txt"
)

func command(s *bufio.Scanner) (string, int) {
	if s == nil {
		return "", 0
	}

	cmd := strings.Split(s.Text(), " ")
	l, err := strconv.Atoi(cmd[1])
	if err != nil {
		panic(err)
	}

	return cmd[0], l
}

func Location(r io.Reader) (int, int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	return location(s)
}

func location(s *bufio.Scanner) (h, v int) {
	if !s.Scan() {
		return
	}

	direction, steps := command(s)
	switch direction {
	case "up":
		v = -steps
	case "down":
		v = steps
	case "forward":
		h = steps
	}

	h1, v1 := location(s)
	return h + h1, v + v1
}

func AimedLocation(r io.Reader) (int, int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	return aimedLocation(s, 0, 0, 0)
}

func aimedLocation(s *bufio.Scanner, aim, h, v int) (int, int) {
	if !s.Scan() {
		return h, v
	}

	direction, steps := command(s)
	switch direction {
	case "up":
		aim -= steps
	case "down":
		aim += steps
	case "forward":
		h += steps
		v += aim * steps
	}

	return aimedLocation(s, aim, h, v)
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	h, v := Location(tee)
	fmt.Printf("location: %v, %v : %v\n", h, v, h*v)

	h, v = AimedLocation(&buf)
	fmt.Printf("aimed location: %v, %v : %v\n", h, v, h*v)
}

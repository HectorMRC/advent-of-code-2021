package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	common "github.com/HectorMRC/advent-of-code-2021"
)

const (
	InputPath = "./day_2/input.txt"
	BufSize   = 9
)

func command(r io.Reader) (d string, s int, ok bool) {
	if ok = r != nil; !ok {
		return
	}

	buf := make([]byte, BufSize)
	l, err := r.Read(buf)
	if ok = l > 0 && err == nil; !ok {
		return
	}

	cmd := strings.Split(string(buf[:l]), " ")

	d = cmd[0]
	s, err = strconv.Atoi(cmd[1])
	if err != nil {
		panic(err)
	}

	return
}

func Location(r io.Reader) (h, v int) {
	direction, steps, ok := command(r)
	if !ok {
		return
	}

	switch direction {
	case "up":
		v = -steps
	case "down":
		v = steps
	case "forward":
		h = steps
	}

	h1, v1 := Location(r)
	return h + h1, v + v1
}

func AimedLocation(r io.Reader, aim, h, v int) (int, int) {
	direction, steps, ok := command(r)
	if !ok {
		return h, v
	}

	switch direction {
	case "up":
		aim -= steps
	case "down":
		aim += steps
	case "forward":
		h += steps
		v += aim * steps
	}

	return AimedLocation(r, aim, h, v)
}

func main() {
	file, err := os.Open(InputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := common.NewReader(file)

	//h, v := Location(r)
	//fmt.Printf("location: %v, %v : %v\n", h, v, h*v)

	h, v := AimedLocation(r, 0, 0, 0)
	fmt.Printf("aimed location: %v, %v : %v\n", h, v, h*v)
}

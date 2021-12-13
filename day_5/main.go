package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputPath            = "./day_5/input.txt"
	PointsSeparator      = "->"
	CoordinatesSeparator = ","
)

var (
	ErrBadFormat = errors.New("bad format")
)

type point [2]int

func lineEquationFrom2Points(p, q *point) (a, b, c int) {
	a = q[1] - p[1]
	b = p[0] - q[0]
	c = -(a*(p[0]) + b*(p[1]))

	if b < 0 {
		return a, -b, -c
	}

	return
}

func isBetween(subject, p, q *point) bool {
	min := minPoint(p, q)
	max := maxPoint(p, q)

	return !(subject[0] < min[0] || subject[0] > max[0] || subject[1] < min[1] || subject[1] > max[1])
}

func minPoint(p, q *point) *point {
	r := &point{p[0], p[1]}
	if q[0] < r[0] {
		r[0] = q[0]
	}

	if q[1] < r[1] {
		r[1] = q[1]
	}

	return r
}

func maxPoint(p, q *point) *point {
	r := &point{p[0], p[1]}
	if q[0] > r[0] {
		r[0] = q[0]
	}

	if q[1] > r[1] {
		r[1] = q[1]
	}

	return r
}

func parsePoint(s string) (*point, error) {
	coordinates := strings.Split(s, CoordinatesSeparator)
	if len(coordinates) != 2 {
		return nil, ErrBadFormat
	}

	x, err := strconv.Atoi(strings.TrimSpace(coordinates[0]))
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(strings.TrimSpace(coordinates[1]))
	if err != nil {
		return nil, err
	}

	return &point{x, y}, nil
}

func parsePoints(s string) (p, q *point, err error) {
	points := strings.Split(s, PointsSeparator)
	if len(points) != 2 {
		err = ErrBadFormat
		return
	}

	p, err = parsePoint(points[0])
	if err != nil {
		return
	}

	q, err = parsePoint(points[1])
	return
}

func HowManyPointsBelongs2MultipleLines(r io.Reader, diagonal bool) (int, error) {
	if r == nil {
		return 0, nil
	}

	equations := [][3]int{}
	bounds := [][2]*point{}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	min := &point{0, 0}
	max := &point{0, 0}
	for s.Scan() {
		p, q, err := parsePoints(s.Text())
		if err != nil {
			return 0, err
		}

		a, b, c := lineEquationFrom2Points(p, q)
		if diagonal || a == 0 || b == 0 {
			min = minPoint(min, p)
			min = minPoint(min, q)

			max = maxPoint(max, p)
			max = maxPoint(max, q)

			equations = append(equations, [3]int{a, b, c})
			bounds = append(bounds, [2]*point{p, q})

		}
	}

	total := 0
	for y := min[1]; y <= max[1]; y++ {
		for x := min[0]; x <= max[0]; x++ {
			count := 0
			for i, eq := range equations {
				if eq[0]*x+eq[1]*y+eq[2] == 0 && isBetween(&point{x, y}, bounds[i][0], bounds[i][1]) {
					count++
				}
				// } else {
				// 	log.Printf("equation: %v(%v) + %v(%v) + %v = %v", eq[0], x, eq[1], y, eq[2], eq[0]*x+eq[1]*y+eq[2])
				// }

				// if count > 1 {
				// 	break
				// }
			}

			c := "."
			if count > 0 {
				c = fmt.Sprint(count)
			}

			fmt.Print(c)

			if count > 1 {
				total++
			}
		}

		fmt.Println()
	}

	return total, nil
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	total, err := HowManyPointsBelongs2MultipleLines(tee, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("how many points belongs to multiple (non diagonal) lines: %v\n", total)

	total, err = HowManyPointsBelongs2MultipleLines(&buf, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("how many points belongs to multiple (all directions) lines: %v\n", total)
}

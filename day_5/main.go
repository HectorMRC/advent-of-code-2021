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

	horizontal = iota
	vertical
	other
)

var (
	ErrBadFormat = errors.New("bad format")
	MaxOverlaps  = 2
)

type point [2]int
type belongsFunc func(p *point) bool

func getLine(p, q *point) (f belongsFunc, orientation int) {
	changeInX := q[0] - p[0]
	changeInY := q[1] - p[1]

	f = func(r *point) bool {
		if changeInX == 0 {
			// is vertical
			return r[0] == p[0]
		} else if changeInY == 0 {
			// is horizontal
			return r[1] == p[1]
		}

		slope := changeInY / changeInX
		return p[1]-r[1] == slope*(p[0]-r[0])
	}

	if changeInX == 0 {
		orientation = vertical
	} else if changeInY == 0 {
		orientation = horizontal
	} else {
		orientation = other
	}

	return
}

func isBetween(subject, p, q *point) bool {
	bounds, ok := getBounds(p, q)
	if !ok {
		return false
	}

	return !(subject[0] < bounds[0][0] || subject[0] > bounds[1][0] || subject[1] < bounds[0][1] || subject[1] > bounds[1][1])
}

func getBounds(p ...*point) ([2]*point, bool) {
	if len(p) == 0 {
		return [2]*point{nil, nil}, false
	}

	var min, max *point
	for i := 0; i < len(p); i++ {
		if p[i] == nil {
			continue
		}

		if min == nil && max == nil {
			min = &point{p[i][0], p[i][1]}
			max = &point{p[i][0], p[i][1]}
			continue
		}

		if p[i][0] < min[0] {
			min[0] = p[i][0]
		}

		if p[i][1] < min[1] {
			min[1] = p[i][1]
		}

		if p[i][0] > max[0] {
			max[0] = p[i][0]
		}

		if p[i][1] > max[1] {
			max[1] = p[i][1]
		}
	}

	return [2]*point{min, max}, true
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

	equations := []belongsFunc{}
	limits := [][2]*point{}
	var bounds [2]*point

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		p, q, err := parsePoints(s.Text())
		if err != nil {
			return 0, err
		}

		belongs, orientation := getLine(p, q)
		if diagonal || orientation < other {
			newBounds, ok := getBounds(bounds[0], bounds[1], p, q)
			if !ok {
				log.Println("failed while geting bounds")
				continue
			}

			bounds = newBounds
			equations = append(equations, belongs)
			limits = append(limits, [2]*point{p, q})
		}
	}

	total := 0
	for y := bounds[0][1]; y <= bounds[1][1]; y++ {
		for x := bounds[0][0]; x <= bounds[1][0]; x++ {
			count := 0
			for i, belongs := range equations {
				if belongs(&point{x, y}) && isBetween(&point{x, y}, limits[i][0], limits[i][1]) {
					count++
				}

				if MaxOverlaps > 0 && count > MaxOverlaps-1 {
					break
				}
			}

			if count > MaxOverlaps-1 {
				total++
			}
		}
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

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
	inputPath = "./day_4/input.txt"
)

func boolMatrix(r int, c int) [][]bool {
	m := make([][]bool, r)
	for i := range m {
		m[i] = make([]bool, r)
	}

	return m
}

type player struct {
	board          [][]string
	marked         [][]bool
	markedByColumn []int
	markedByRow    []int
}

func (player *player) mark(str string) bool {
	for r := 0; r < len(player.board); r++ {
		for c := 0; c < len(player.board[r]); c++ {
			player.marked[r][c] = player.marked[r][c] || player.board[r][c] == str
			if player.board[r][c] == str {
				player.markedByColumn[c]++
				player.markedByRow[r]++
			}

			if player.markedByColumn[c] == len(player.board) {
				return true
			}
		}

		if player.markedByRow[r] == len(player.board[r]) {
			return true
		}
	}

	return false
}

func buildPlayboard(s *bufio.Scanner) (input []string, players []player) {
	if s == nil || !s.Scan() {
		return nil, nil
	}

	input = strings.Split(s.Text(), ",")

	var current player
	for s.Scan() {
		if len(s.Text()) == 0 {
			if len(current.board) > 0 {
				current.marked = boolMatrix(len(current.board), len(current.board[0]))
				current.markedByColumn = make([]int, len(current.board[0]))
				current.markedByRow = make([]int, len(current.board))
				players = append(players, current)
			}

			current = player{}
			continue
		}

		row := strings.Fields(s.Text())
		current.board = append(current.board, row)
	}

	return
}

func getWinner(input []string, players []player) (*player, string) {
	for _, str := range input {
		for _, player := range players {
			if player.mark(str) {
				return &player, str
			}
		}
	}

	return nil, ""
}

func getLoser(input []string, players []player) (loser *player, latest string) {
	for _, str := range input {
		if len(players) == 0 {
			return
		}

		playing := []player{}
		latest = str

		for _, player := range players {
			if loser = &player; !loser.mark(str) {
				playing = append(playing, player)
			}
		}

		players = playing
	}

	return
}

func calcScore(player *player, latest int) int {
	sum := 0
	for r := range player.board {
		for c, str := range player.board[r] {
			if !player.marked[r][c] {
				value, err := strconv.Atoi(str)
				if err != nil {
					log.Fatal(err)
				}

				sum += value
			}
		}
	}

	return sum * latest
}

func Winner(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	input, players := buildPlayboard(s)
	winner, latest := getWinner(input, players)

	if winner == nil {
		log.Fatal("no winner")
	}

	first, err := strconv.Atoi(latest)
	if err != nil {
		log.Fatal(err)
	}

	return calcScore(winner, first), nil
}

func Loser(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	input, players := buildPlayboard(s)
	loser, latest := getLoser(input, players)

	if loser == nil {
		log.Fatal("no loser")
	}

	last, err := strconv.Atoi(latest)
	if err != nil {
		log.Fatal(err)
	}

	return calcScore(loser, last), nil
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	winner, err := Winner(tee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("winner score: %v\n", winner)

	loser, err := Loser(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("loser score: %v\n", loser)
}

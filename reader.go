package adventofcode2021

import (
	"bufio"
	"io"
	"os"
)

type reader struct {
	scanner *bufio.Scanner
}

func (r *reader) Read(buf []byte) (int, error) {
	if !r.scanner.Scan() {
		return 0, r.scanner.Err()
	}

	data := r.scanner.Bytes()
	for i, b := range data {
		if len(buf) <= i {
			return i, io.ErrShortBuffer
		}

		buf[i] = b
	}

	return len(data), nil
}

func NewReader(f *os.File) io.Reader {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	return &reader{scanner}
}

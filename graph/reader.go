package graph

import (
	"io"
	"bufio"
	"strings"
	"strconv"
)

type Reader struct {
	r *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

func (r *Reader) Read() (left, right int) {
	left, right = r.ReadLine()
	return left, right
}

func (r *Reader) ReadLine() (int, int) {
	var line string
	line, _ = r.r.ReadString('\n')
	values := strings.Split(line, " ")
	a := strings.TrimSpace(values[0])
	b := strings.TrimSpace(values[1])
	first, _ := strconv.Atoi(a)
	second, _ := strconv.Atoi(b)
	return first, second
}

package stats

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Statistics struct {
	numbers    []int
	numbersLen uint
}

func NewStat() *Statistics {
	return &Statistics{
		numbers:    nil,
		numbersLen: 0,
	}
}

func (st *Statistics) fillFields() {
	sort.Ints(st.numbers)
	st.numbersLen = uint(len(st.numbers))
}

func (st *Statistics) Init() {
	var input string
	var err error
	var reader = bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input = reader.Text()
		num, errParse := strconv.Atoi(input)
		isRightBounds := num >= -100_000 && num <= 100_000
		if errParse != nil || !isRightBounds {
			fmt.Println("Error: input must be an integer number")
			continue
		}
		if !errors.Is(err, io.EOF) {
			st.numbers = append(st.numbers, num)
		}
	}
	st.fillFields()
}

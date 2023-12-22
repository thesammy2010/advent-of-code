package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Seat struct {
	Row    int
	Column int
	SeatID int
}

func Convert(line string) int {
	line = strings.ReplaceAll(line, "F", "0")
	line = strings.ReplaceAll(line, "B", "1")
	line = strings.ReplaceAll(line, "L", "0")
	line = strings.ReplaceAll(line, "R", "1")
	n, err := strconv.ParseInt(line, 2, 64)

	if err != nil {
		return 0
	}

	return int(n)

}

func CreateSeat(line string) Seat {

	row := Convert(line[:7])
	column := Convert(line[7:])
	seatID := row*8 + column

	s := Seat{row, column, seatID}

	return s

}

func Main(filepath string) int {

	lines, err := helpers.ReadFileLineByLine(filepath)

	if err != nil {
		fmt.Println(err)
	}
	seats := make([]Seat, 0)

	for _, i := range lines {
		seats = append(seats, CreateSeat(i))
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].SeatID > seats[j].SeatID
	})

	return seats[0].SeatID
}

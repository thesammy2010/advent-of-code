package part2

import (
	"advent-of-code/advent_of_code/day5/part1"
	"advent-of-code/advent_of_code/helpers"
	"fmt"
	"sort"
)

func getMissingSeat(seatIds []int) int {

	for i := range seatIds {

		leftSeat := seatIds[i]
		potentialSeat := leftSeat + 1
		rightSeat := seatIds[i+1]

		if potentialSeat+1 == rightSeat {
			return potentialSeat
		}

	}

	return 0

}

func Main(filepath string) int {

	lines, err := helpers.ReadFileLineByLine(filepath)

	if err != nil {
		fmt.Println(err)
	}
	seats := make([]part1.Seat, 0)

	for _, i := range lines {
		seats = append(seats, part1.CreateSeat(i))
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].SeatID < seats[j].SeatID
	})

	seatIds := make([]int, 0)
	for i := range seats {
		seatIds = append(seatIds, seats[i].SeatID)
	}

	return getMissingSeat(seatIds)
}

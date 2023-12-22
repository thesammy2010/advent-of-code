package part2

import (
	"advent-of-code/advent_of_code/day4/part1"
	"advent-of-code/advent_of_code/day4/part2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassportIsValidV2(t *testing.T) {

	p1 := part1.Passport{"1980", "2012", "2030", "74in", "#623a2f", "grn", "087499704", "147"}
	p2 := part1.Passport{"1989", "2014", "2029", "165cm", "#a97842", "blu", "896056539", "129"}
	p3 := part1.Passport{"2001", "2015", "2022", "164cm", "#888785", "hzl", "545766238", "88"}
	p4 := part1.Passport{"1944", "2010", "2021", "158cm", "#b6652a", "blu", "093154719", "88"}
	p5 := part1.Passport{"1926", "2018", "1972", "186cm", "#18171d", "amb", "186cm", "100"}
	p6 := part1.Passport{"1946", "2019", "1967", "170cm", "#602927", "grn", "012533040", ""}
	p7 := part1.Passport{"1992", "2012", "2020", "182cm", "dab227", "brn", "021572410", "277"}
	p8 := part1.Passport{"2007", "2023", "2038", "59cm", "74454a", "zzz", "3556412378", "277"}

	assert.True(t, part2.IsPassportValidV2(p1))
	assert.True(t, part2.IsPassportValidV2(p2))
	assert.True(t, part2.IsPassportValidV2(p3))
	assert.True(t, part2.IsPassportValidV2(p4))
	assert.False(t, part2.IsPassportValidV2(p5))
	assert.False(t, part2.IsPassportValidV2(p6))
	assert.False(t, part2.IsPassportValidV2(p7))
	assert.False(t, part2.IsPassportValidV2(p8))

}

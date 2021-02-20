package part1

import (
	"advent-of-code/advent_of_code/day4/part1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanup(t *testing.T) {
	expectedResult := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
	}
	calculatedResult := part1.CleanupInput("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929")
	assert.Equal(t, expectedResult, calculatedResult, "The two sentences should be the same.")
}

func TestCreatePassport(t *testing.T) {
	expectedResult := part1.Passport{"1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147"}
	calculatedResult := part1.CreatePassport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
	assert.Equal(t, expectedResult, calculatedResult, "The two sentences should be the same.")
}

func TestIsValid(t *testing.T) {
	validPassport := part1.Passport{"1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147"}
	invalidPassport := part1.Passport{"1937", "2017", "2020", "183cm", "#fffffd", "", "860033327", "147"}
	assert.True(t, part1.IsPassportValid(validPassport))
	assert.False(t, part1.IsPassportValid(invalidPassport))
}

func TestMain(t *testing.T) {
	assert.Equal(t, part1.Main("../input_test.txt"), 2)
}

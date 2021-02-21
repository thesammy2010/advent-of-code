package part2

import (
	"advent-of-code/advent_of_code/day4/part1"
	"advent-of-code/advent_of_code/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsPassportValidV2(p part1.Passport) bool {

	// byr
	byr, err := strconv.ParseInt(p.BYR, 10, 64)
	if err != nil {
		return false
	}
	if !(1920 <= byr && byr <= 2002) {
		return false
	}

	// iyr
	iyr, err := strconv.ParseInt(p.IYR, 10, 64)
	if err != nil {
		return false
	}
	if !(2010 <= iyr && iyr <= 2020) {
		return false
	}

	// eyr
	eyr, err := strconv.ParseInt(p.EYR, 10, 64)
	if err != nil {
		return false
	}
	if !(2020 <= eyr && eyr <= 2030) {
		return false
	}

	// hgt
	heightPattern := regexp.MustCompile("[0-9]{2,3}(cm|in)")
	if len(heightPattern.Find([]byte(p.HGT))) == 0 {
		return false
	} else {
		len := len(p.HGT)
		hgt, _ := strconv.ParseInt(string(p.HGT)[0:len-2], 10, 64)
		if strings.Contains(p.HGT, "cm") {
			if !(150 <= hgt && hgt <= 193) {
				return false
			}
		} else if strings.Contains(p.HGT, "in") {
			if !(59 <= hgt && hgt <= 76) {
				return false
			}
		} else {
			return false
		}
	}

	// hcl
	colourPattern := regexp.MustCompile("#[0-9,a-f]{6}")
	if !colourPattern.MatchString(p.HCL) {
		return false
	}
	if colourPattern.FindStringSubmatch(p.HCL)[0] != p.HCL {
		return false
	}

	// ecl
	allowedColours := regexp.MustCompile("[amb,blu,brn,gry,grn,hzl,oth]{3}")
	if !allowedColours.MatchString(p.ECL) {
		return false
	}
	if allowedColours.FindStringSubmatch(p.ECL)[0] != p.ECL {
		return false
	}

	// pid
	passportPattern := regexp.MustCompile("[0-9]{9}")
	if !passportPattern.MatchString(p.PID) {
		return false
	}
	if passportPattern.FindStringSubmatch(p.PID)[0] != p.PID {
		return false
	}

	return true

}

func Main(filepath string) int {
	data, err := helpers.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	var results int
	records := part1.CleanupInput(data)
	for _, line := range records {
		passport := part1.CreatePassport(line)
		if part1.IsPassportValid(passport) {
			if IsPassportValidV2(passport) {
				results++
			}
		}
	}
	return results
}

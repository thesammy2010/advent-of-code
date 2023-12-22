package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"fmt"
	"strings"
)

type Passport struct {
	BYR string
	IYR string
	EYR string
	HGT string
	HCL string
	ECL string
	PID string
	CID string
}

func CleanupInput(input string) []string {
	split := strings.Split(input, "\n\n")
	newSplit := []string{}
	for idx := range split {
		newSplit = append(newSplit, strings.ReplaceAll(split[idx], "\n", " "))
	}
	return newSplit
}

func CreatePassport(input string) Passport {
	p := Passport{}
	keyPairs := strings.Split(input, " ")
	for _, data := range keyPairs {
		splitData := strings.Split(data, ":")
		key, val := splitData[0], splitData[1]
		if key == "byr" {
			p.BYR = val
		} else {
			if key == "iyr" {
				p.IYR = val
			} else {
				if key == "eyr" {
					p.EYR = val
				} else {
					if key == "hgt" {
						p.HGT = val
					} else {
						if key == "hcl" {
							p.HCL = val
						} else {
							if key == "ecl" {
								p.ECL = val
							} else {
								if key == "pid" {
									p.PID = val
								} else {
									if key == "cid" {
										p.CID = val
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return p
}

func IsPassportValid(p Passport) bool {
	if p.BYR == "" {
		return false
	} else {
		if p.IYR == "" {
			return false
		} else {
			if p.EYR == "" {
				return false
			} else {
				if p.HGT == "" {
					return false
				} else {
					if p.HCL == "" {
						return false
					} else {
						if p.ECL == "" {
							return false
						} else {
							if p.PID == "" {
								return false
							}
						}
					}
				}
			}
		}
	}
	return true
}

func Main(filepath string) int {
	data, err := helpers.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	var results int
	records := CleanupInput(data)
	for _, line := range records {
		passport := CreatePassport(line)
		if IsPassportValid(passport) {
			results++
		}
	}
	return results
}

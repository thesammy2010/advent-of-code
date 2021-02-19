from typing import List
from unittest import TestCase

from advent_of_code.day4.part1.main import Passport
from advent_of_code.day4.part2.main import is_passport_valid_v2, main


class TestDay4Part1(TestCase):

    def test_is_passport_valid(self) -> None:

        p1: Passport = Passport(
            ecl="grn", pid="087499704", eyr="2030", hcl="#623a2f", byr="1980", iyr="2012", cid="147", hgt="74in"
        )
        p2: Passport = Passport(
            eyr="2029", ecl="blu", cid="129", byr="1989", iyr="2014", pid="896056539", hcl="#a97842", hgt="165cm"
        )
        p3: Passport = Passport(
            hcl="#888785", hgt="164cm", byr="2001", iyr="2015", cid="88", pid="545766238", ecl="hzl", eyr="2022"
        )
        p4: Passport = Passport(
            iyr="2010", hgt="158cm", hcl="#b6652a", ecl="blu", byr="1944", eyr="2021", pid="093154719"
        )        
        p5: Passport = Passport(
            eyr="1972", cid="100", hcl="#18171d", ecl="amb", hgt="170", pid="186cm", iyr="2018", byr="1926"
        )
        p6: Passport = Passport(
            iyr="2019", hcl="#602927", eyr="1967", hgt="170cm", ecl="grn", pid="012533040", byr="1946"
        )
        p7: Passport = Passport(
            hcl="dab227", iyr="2012", ecl="brn", hgt="182cm", pid="021572410", eyr="2020", byr="1992", cid="277"
        )
        p8: Passport = Passport(
            hgt="59cm", ecl="zzz", eyr="2038", hcl="74454a", iyr="2023", pid="3556412378", byr="2007"
        )
        self.assertEqual(is_passport_valid_v2(p1), True)
        self.assertEqual(is_passport_valid_v2(p2), True)
        self.assertEqual(is_passport_valid_v2(p3), True)
        self.assertEqual(is_passport_valid_v2(p4), True)
        self.assertEqual(is_passport_valid_v2(p5), False)
        self.assertEqual(is_passport_valid_v2(p6), False)
        self.assertEqual(is_passport_valid_v2(p7), False)
        self.assertEqual(is_passport_valid_v2(p8), False)

    def test_main(self) -> None:
        self.assertEqual(2, main("tests/day4/input_test.txt"))
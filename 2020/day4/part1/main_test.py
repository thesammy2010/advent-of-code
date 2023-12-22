from typing import List
from unittest import TestCase

from advent_of_code.day4.part1.main import Passport, main, cleanup_input, create_passport, is_passport_valid


class TestDay4Part1(TestCase):

    def test_cleanup_input(self) -> None:

        input_data: str = (
            "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\n"
            "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929"
        )
        expected_output: List[str] = [
            "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
            "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929"
        ]
        self.assertEqual(expected_output, cleanup_input(file=input_data))

    def test_create_passport(self) -> None:

        input_data: str = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
        expected_output: Passport = Passport(
            ecl="gry", pid="860033327", eyr="2020", hcl="#fffffd", byr="1937", iyr="2017", cid="147", hgt="183cm"
        )
        self.assertEqual(
            expected_output.__dict__, create_passport(line=input_data).__dict__
        )

    def test_is_passport_valid(self) -> None:
        p1: Passport = Passport(
            ecl="gry", pid="860033327", eyr="2020", hcl="#fffffd", byr="1937", iyr="2017", cid="147", hgt="183cm"
        )
        p2: Passport = Passport(
            pid="860033327", eyr="2020", hcl="#fffffd", byr="1937", iyr="2017", cid="147", hgt="183cm"
        )
        self.assertTrue(is_passport_valid(p1))
        self.assertFalse(is_passport_valid(p2))

    def test_main(self) -> None:
        self.assertEqual(2, main("tests/day4/input_test.txt"))
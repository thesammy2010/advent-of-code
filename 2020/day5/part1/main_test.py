from unittest import TestCase

from advent_of_code.day5.part1.main import Seat, convert, create_seat, main

class TestDayPart1(TestCase):

    def test_convert(self) -> None:
        self.assertEqual(convert("L"), 0)
        self.assertEqual(convert("R"), 1)
        self.assertEqual(convert("LRLRLR"), 21)

    def test_create_seat(self) -> None:
        expectedResult: Seat = Seat(row=44, column=5)
        self.assertEqual(create_seat("FBFBBFFRLR").__dict__, expectedResult.__dict__)

    def test_main(self) -> None:
        self.assertEqual(main(filepath="tests/day5/input_test.txt"), 820)

from unittest import TestCase

from advent_of_code.helpers.helpers import read_file, read_file_line_by_line
from advent_of_code.day3.part2.main import slope, main


class TestDay3Part2(TestCase):

    def setUp(self) -> None:
        self.rows = read_file_line_by_line("tests/day3/input_test.txt")

    def test_slope(self) -> None:
        self.assertEqual(
            slope(rows=self.rows, line_skip=1, slide=1), 2
        )
        self.assertEqual(
            slope(rows=self.rows, line_skip=1, slide=3), 7
        )
        self.assertEqual(
            slope(rows=self.rows, line_skip=1, slide=5), 3
        )
        self.assertEqual(
            slope(rows=self.rows, line_skip=1, slide=7), 4
        )
        self.assertEqual(
            slope(rows=self.rows, line_skip=2, slide=1), 2
        )

    def test_main(self) -> None:
        self.assertEqual(
            main("tests/day3/input_test.txt"), 336
        )
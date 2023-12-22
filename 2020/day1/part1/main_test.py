import unittest

from typing import List, Tuple

from advent_of_code.day1.part1.main import check_sum, make_pairs, main


class TestDay1Part1(unittest.TestCase):
    def test_check_sum(self) -> None:
        input_data: Tuple[int, int] = 2019, 1
        expected_output: Tuple[int, int, int] = 2019, 1, 2019
        self.assertEqual(check_sum(*input_data), expected_output)

    def test_make_pairs(self) -> None:
        input_data: List[int] = [1, 2, 3]
        expected_output: List[List[int]] = [[1, 2], [1, 3], [2, 1], [2, 3], [3, 1], [3, 2]]
        self.assertEqual(make_pairs(data=input_data), expected_output)

    def test_main(self) -> None:
        expected_output: Optional[Tuple[int, int, int]] = 1721, 299, 514579
        actual_input: Optional[Tuple[int, int, int]] = main(file="tests/day1/input_test.txt")
        self.assertEqual(actual_input, expected_output)
import unittest

from typing import List, Tuple

from advent_of_code.day1.part2.main import check_sum_v2, make_triples, main


class TestDay1Part2(unittest.TestCase):
    def test_check_sum_v2(self) -> None:
        input_data: Tuple[int, int] = 2018, 1, 1
        expected_output: Tuple[int, int, int] = 2018, 1, 1, 2018
        self.assertEqual(check_sum_v2(*input_data), expected_output)

    def test_make_triples(self) -> None:
        input_data: List[int] = [1, 2, 3, 4]
        expected_output: List[List[int]] = [
            [1, 2, 3], 
            [1, 2, 4], 
            [1, 3, 2], 
            [1, 3, 4], 
            [1, 4, 2], 
            [1, 4, 3], 
            [2, 1, 3], 
            [2, 1, 4], 
            [2, 3, 1], 
            [2, 3, 4], 
            [2, 4, 1], 
            [2, 4, 3], 
            [3, 1, 2], 
            [3, 1, 4], 
            [3, 2, 1], 
            [3, 2, 4], 
            [3, 4, 1], 
            [3, 4, 2], 
            [4, 1, 2], 
            [4, 1, 3], 
            [4, 2, 1], 
            [4, 2, 3], 
            [4, 3, 1], 
            [4, 3, 2]
        ]
        self.assertEqual(make_triples(data=input_data), expected_output)

    def test_main(self) -> None:
        expected_output: Optional[Tuple[int, int, int]] = 979, 366, 675, 241861950
        actual_input: Optional[Tuple[int, int, int]] = main(file="tests/day1/input_test.txt")
        self.assertEqual(actual_input, expected_output)
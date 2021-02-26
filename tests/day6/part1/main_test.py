from typing import List
from unittest import TestCase

from advent_of_code.day6.part1.main import compile_groups, main
from advent_of_code.helpers.helpers import read_file

class TestDay6Part1(TestCase):

    def test_compile_groups(self) -> None:
        input_data: str = read_file("tests/day6/input_test.txt")
        calculated_output: List[str] = compile_groups(input_data)
        expected_output: List[str] = ['abc', 'abc', 'abac', 'aaaa', 'b', 'abcxabcyabcz']
        self.assertEqual(calculated_output, expected_output)

    def test_main(self) -> None:
        self.assertEqual(main(filepath="tests/day6/input_test.txt"), 17)

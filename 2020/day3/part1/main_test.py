import unittest

from typing import Tuple

from advent_of_code.day3.part1.main import loop, main

class TestDay2Part1(unittest.TestCase):

    def test_loop(self):
        input_data1: Tuple[int, int, str] = 11, 0, "..##......."
        expected_output1: bool = False
        self.assertEqual(loop(*input_data1), expected_output1)

        input_data2: Tuple[int, int, str] = 11, 3, "..##......."
        expected_output2: bool = True
        self.assertEqual(loop(*input_data2), expected_output2)

    def test_main(self):
        input_data: int = main("tests/day3/input_test.txt")
        self.assertEqual(input_data, 7)
    
from typing import Tuple
from unittest import TestCase

from advent_of_code.day8.part1.main import action, main

class TestDay8Part1(TestCase):

    def test_action(self) -> None:
        expected_result1: Tuple[int] = 1, 0
        expected_result2: Tuple[int] = -3, 0
        expected_result3: Tuple[int] = 1, 6
        self.assertEqual(expected_result1, action(row="nop +0"))
        self.assertEqual(expected_result2, action(row="jmp -3"))
        self.assertEqual(expected_result3, action(row="acc +6"))

    def test_main(self) -> None:
        self.assertEqual(main(filepath="tests/day8/input_test.txt"), 5)

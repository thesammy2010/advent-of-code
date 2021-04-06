from typing import List
from unittest import TestCase

from advent_of_code.day8.part2.main import change, main
from advent_of_code.helpers.helpers import read_file_line_by_line

class TestDay8Part2(TestCase):

    def test_change(self) -> None:
        raw_result: List[str] = read_file_line_by_line("tests/day8/input_test.txt")
        expected_result: List[str] = [
            "jmp +0", "acc +1", "jmp +4", "acc +3", "jmp -3", "acc -99", "acc +1", "jmp -4", "acc +6"
        ]
        self.assertListEqual(expected_result, change(rs=raw_result, index_to_change=0))

    def test_main(self) -> None:
        self.assertEqual(main(filepath="tests/day8/input_test.txt"), 8)

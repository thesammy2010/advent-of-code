from typing import Dict, List
from unittest import TestCase

from advent_of_code.day6.part2.main import compile_groups, group_data, get_count_from_group, main
from advent_of_code.helpers.helpers import read_file


class TestDay6Part2(TestCase):

    def test_compile_groups(self) -> None:
        input_data: str = read_file("tests/day6/input_test.txt")
        expected_result: List[List[str]] = [
            ["abc"],
            ["a", "b", "c"],
            ["ab", "ac"],
            ["a", "a", "a", "a"],
            ["b"],
            ["abcx", "abcy", "abcz"]
        ]
        self.assertEqual(expected_result, compile_groups(input_data))

    def test_group_data(self) -> None:
        input_data1: List[str] = ["abc"]
        expected_output1: Dict[str, int] = {"a": 1, "b": 1, "c": 1, "count": 1}
        self.assertEqual(expected_output1, group_data(input_data1))
        input_data2: List[str] = ["abcx", "abcy", "abcz"]
        expected_output2: Dict[str, int] = {"a": 3, "b": 3, "c": 3, "x": 1, "y": 1, "z": 1, "count": 3}
        self.assertEqual(expected_output2, group_data(input_data2))
        input_data3: List[str] = ["a", "a", "a", "a"]
        expected_output3: Dict[str, int] = {"a": 4, "count": 4}
        self.assertEqual(expected_output3, group_data(input_data3))

    def test_get_count_from_group(self) -> None:
        input_data1: Dict[str, int] = {"a": 1, "b": 1, "c": 1, "count": 1}
        self.assertEqual(get_count_from_group(input_data1), 3)
        input_data2: Dict[str, int] = {"a": 3, "b": 3, "c": 3, "x": 1, "y": 1, "z": 1, "count": 3}
        self.assertEqual(get_count_from_group(input_data2), 3)
        input_data3: Dict[str, int] = {"a": 4, "count": 4}
        self.assertEqual(get_count_from_group(input_data3), 1)
        input_data4: Dict[str, int] = {"a": 2, "b": 1, "c": 1, "count": 2}
        self.assertEqual(get_count_from_group(input_data4), 1)

    def test_main(self) -> None:
        self.assertEqual(main(filepath="tests/day6/input_test.txt"), 9)

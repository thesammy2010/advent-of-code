import unittest

from advent_of_code.day2.part2.main import Policy
from advent_of_code.day2.part2.main import is_password_valid, main
class TestDay2Part2(unittest.TestCase):

    def test_is_policy_valid(self):
        p1: Policy = Policy(*(1, 3, "a", "abcde"))
        self.assertTrue(is_password_valid(p1))
        p2: Policy = Policy(*(1, 3, "b", "cdefg"))
        self.assertFalse(is_password_valid(p2))
        p3: Policy = Policy(*(2, 9, "c", "ccccccccc"))
        self.assertFalse(is_password_valid(p3))

    def test_main(self):
        input_data: int = main("tests/day2/input_test.txt")
        self.assertEqual(input_data, 1)
    
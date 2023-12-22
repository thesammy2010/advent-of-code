import unittest

from advent_of_code.day2.part1.main import Policy, make_policy, main

class TestDay2Part1(unittest.TestCase):

    def test_make_policy(self):
        input_data: str = "1-4 q: abcqrs"
        expected_output: Policy = Policy(lower_limit=1, upper_limit=4, character="q", password="abcqrs")
        self.assertEqual(make_policy(line=input_data).__dict__, expected_output.__dict__)

    def test_is_policy_valid(self):
        p1: Policy = Policy(*(1, 4, "q", "q"))
        self.assertTrue(p1.is_password_valid())
        p2: Policy = Policy(*(1, 4, "q", ""))
        self.assertFalse(p2.is_password_valid())
        p3: Policy = Policy(*(1, 4, "q", "qqqqq"))
        self.assertFalse(p3.is_password_valid())

    def test_main(self):
        input_data: int = main("tests/day2/input_test.txt")
        self.assertEqual(input_data, 2)
    
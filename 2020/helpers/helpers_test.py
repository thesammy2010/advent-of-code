import unittest

from advent_of_code.helpers.helpers import read_file, read_file_line_by_line

class TestHelpers(unittest.TestCase):
    def test_read_file(self) -> None:
        self.assertEqual(
            read_file(filepath="tests/helpers/helpers_test.txt"),
            "this\nis\na\ntest\n"
        )
    def test_read_file_line_by_line(self) -> None:
        self.assertEqual(
            read_file_line_by_line(filepath="tests/helpers/helpers_test.txt"),
            ["this", "is", "a", "test"]
        )

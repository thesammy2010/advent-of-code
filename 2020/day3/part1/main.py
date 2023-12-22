from advent_of_code.helpers.helpers import read_file_line_by_line

from typing import List


def loop(length_of_row: int, index: int, row: List[str]) -> bool:
    pos: str = row[index % length_of_row]  # maths x mod n for circular

    if pos == "#":
        return True
    return False


def main(file: str = "") -> int:

    input_file: List[str] = read_file_line_by_line(filepath=file)
    count: int = 0
    index: int = 0

    for row in input_file:
        res: int = loop(len(row), index=index, row=row)
        count += 1 if res else 0
        index += 3

    return count

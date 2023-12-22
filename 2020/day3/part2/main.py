from functools import reduce
from typing import List, Tuple


from advent_of_code.helpers.helpers import read_file_line_by_line
from advent_of_code.day3.part1.main import loop


def slope(rows: List[str], slide: int, line_skip: int = 1) -> int:

    count: int = 0
    index: int = 0
    line_number: int = 0

    while line_number < len(rows):
        row: List[str] = rows[line_number]
        res: int = loop(len(row), index=index, row=row)
        count += 1 if res else 0
        index += slide
        line_number += line_skip

    return count


def main(file: str = "") -> int:

    input_file: List[str] = read_file_line_by_line(filepath=file)

    conditions: List[Tuple[int, int]] = [
        (1, 1), (1, 3), (1, 5), (1, 7), (2, 1)
    ]
    counts: List[int] = []

    for down, right in conditions:
        counts.append(
            slope(rows=input_file, slide=right, line_skip=down)  
        )

    return reduce((lambda x, y: x * y), counts)



from typing import List, Optional, Tuple
from advent_of_code.helpers.helpers import read_file_line_by_line


def check_sum(x: int, y: int) -> Optional[Tuple[int, int, int]]:
    if x+y == 2020:
        return x, y, x*y


def cast_to_int(x: str) -> int:
    return int(x)


def make_pairs(data: List[int]) -> List[List[int]]:
    output: List[int] = []
    # i and j are the pairs
    for i in data:
        for j in data:
            if i != j:
                output.append([i, j])
    return output


def main(file: str = "advent_of_code/day1/input.txt") ->  Optional[Tuple[int, int, int]]:
    input_data: List[str] = read_file_line_by_line(filepath=file)
    input_integers: List[int] = [cast_to_int(x=row) for row in input_data]
    number_pairs: List[List[int]] = make_pairs(data=input_integers)

    for pair in number_pairs:
        res: Optional[Tuple[int, int, int]] = check_sum(x=pair[0], y=pair[1])
        if res is not None:
            return(res)

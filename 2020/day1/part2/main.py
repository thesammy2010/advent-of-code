from typing import List, Optional, Tuple

from advent_of_code.helpers.helpers import read_file_line_by_line
from advent_of_code.day1.part1.main import cast_to_int


def check_sum_v2(x: int, y: int, z: int) -> Optional[Tuple[int, int, int, int]]:
    if x+y+z == 2020:
        return x, y, z, x*y*z


def make_triples(data: List[int]) -> List[List[int]]:
    output: List[int] = []
    # i and j are the pairs
    for i in data:
        for j in data:
            for k in data:
                if i != j and i != k and j != k:
                    output.append([i, j, k])
    return output


def main(file: str = "advent_of_code/day1/input.txt") ->  Optional[Tuple[int, int, int]]:
    input_data: List[str] = read_file_line_by_line(filepath=file)
    input_integers: List[int] = [cast_to_int(x=row) for row in input_data]
    number_triples: List[List[int]] = make_triples(data=input_integers)

    for triple in number_triples:
        res: Optional[Tuple[int, int, int]] = check_sum_v2(x=triple[0], y=triple[1], z=triple[2])
        if res is not None:
            return(res)

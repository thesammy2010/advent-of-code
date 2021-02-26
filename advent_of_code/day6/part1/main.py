from typing import List

from advent_of_code.helpers.helpers import read_file


def compile_groups(line: str) -> List[str]:
    return [i.replace("\n", "") for i in line.split("\n\n")]


def main(filepath: str) -> int:
    lines: str = read_file(filepath=filepath)
    groups: List[str] = compile_groups(lines)
    return sum([len(set(i)) for i in groups])

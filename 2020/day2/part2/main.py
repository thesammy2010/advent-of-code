from typing import List

from advent_of_code.day2.part1.main import make_policy, Policy
from advent_of_code.helpers.helpers import read_file_line_by_line


def is_password_valid(p: Policy) -> bool:

    i: List[int] = [p.lower_limit - 1, p.upper_limit - 1]  # indexes

    if (
        (p.password[i[0]] != p.character and p.password[i[1]] == p.character)
        or
        (p.password[i[0]] == p.character and p.password[i[1]] != p.character)
    ):
        return True
    return False


def main(file: str = "") -> int:
    f: List[str] = read_file_line_by_line(filepath=file)

    results: List[bool] = []

    for line in f:
        policy = make_policy(line=line)
        results.append(is_password_valid(p=policy))

    count: int = 0 if not results.count(True) else results.count(True)

    return count
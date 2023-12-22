from typing import List

from advent_of_code.helpers.helpers import read_file_line_by_line

class Policy(object):

    def __init__(
        self, 
        lower_limit: int = 0, 
        upper_limit: int = 0, 
        character: str = "", 
        password: str = ""
    ) -> None:
        self.lower_limit = lower_limit
        self.upper_limit = upper_limit
        self.character = character
        self.password = password

    def is_password_valid(self) -> bool:
        return self.lower_limit <= self.password.count(self.character) <= self.upper_limit



def make_policy(line: str) -> Policy:

    p: Policy = Policy()
    split_line = line.split()

    p.lower_limit = int(split_line[0].split("-")[0])
    p.upper_limit = int(split_line[0].split("-")[1])
    p.character = split_line[1][0]
    p.password = split_line[-1]

    return p


def main(file: str = "") -> int:
    f: List[str] = read_file_line_by_line(filepath=file)

    results: List[bool] = []

    for line in f:
        policy = make_policy(line=line)
        results.append(policy.is_password_valid())

    count = 0 if not results.count(True) else results.count(True)

    return count

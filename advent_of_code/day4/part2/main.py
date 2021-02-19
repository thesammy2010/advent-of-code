import string

from typing import List

from advent_of_code.helpers.helpers import read_file
from advent_of_code.day4.part1.main import Passport, cleanup_input, create_passport, is_passport_valid


def is_passport_valid_v2(p: Passport) -> bool:

    if not (1920 <= int(p.byr) <= 2002):
        return False
    if not (2010 <= int(p.iyr) <= 2020):
        return False
    if not (2020 <= int(p.eyr) <= 2030):
        return False

    # hgt:
    if "cm" in p.hgt:
        if not (150 <= int(p.hgt[:-2]) <= 193):
            return False
    elif "in" in p.hgt:
        if not (59 <= int(p.hgt[:-2]) <= 76):
            return False
    else:
        return False

    if len(p.pid) != 9:
        return False

    # hcl:
    if p.hcl[0] != "#":
        return False
    for char in list(p.hcl[1:]):
        if char not in string.hexdigits:
            return False

    
    if p.ecl not in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]:
        return False

    return True


def main(filepath: str) -> int:

    file: str = read_file(filepath=filepath)
    items: List[str] = cleanup_input(file=file)
    passorts: List[Passport] = [create_passport(i) for i in items]
    count: int = 0

    for p in passorts:
        if is_passport_valid(p):
            count += 1 if is_passport_valid_v2(p) else 0

    return count

from typing import Dict, List

from advent_of_code.helpers.helpers import read_file

class Passport(object):
    def __init__(
        self, 
        byr: str = None, 
        iyr: str = None, 
        eyr: str = None, 
        hgt: str = None, 
        hcl: str = None, 
        ecl: str = None, 
        pid: str = None, 
        cid: str = None
    ) -> None:
        self.byr: str = byr
        self.iyr: str = iyr
        self.eyr: str = eyr
        self.hgt: str = hgt
        self.hcl: str = hcl
        self.ecl: str = ecl
        self.pid: str = pid
        self.cid: str = cid


def cleanup_input(file: str) -> List[str]:
    # take input and return a list of passport strings
    return [i.replace("\n", " ") for i in file.split("\n\n")]


def create_passport(line: str) -> Passport:
    # line = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"

    tags: List[str] = line.split()
    # tags = ["ecl:gry", "pid:860033327", "eyr:2020", "hcl:#fffffd", "byr:1937", "iyr:2017", "cid:147", "hgt:183cm"]

    # the smart python way with dictionary comprehension
    tags_map: Dict[str, str] = {
        t.split(":")[0]: t.split(":")[1] for t in tags
    }
    # tags_map = {"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", 
    # "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"}

    return Passport(**tags_map)  # **kwargs


def is_passport_valid(p: Passport) -> bool:
    # all but cid must be set or have a value

    for key, value in p.__dict__.items():
        if key == "cid":
            continue
        elif not value:
            return False

    return True


def main(filepath: str) -> int:

    file: str = read_file(filepath=filepath)
    items: List[str] = cleanup_input(file=file)
    passorts: List[Passport] = [create_passport(i) for i in items]
    count: int = 0

    for p in passorts:
        count += 1 if is_passport_valid(p) else 0

    return count
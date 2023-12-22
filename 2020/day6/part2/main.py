from collections import defaultdict
from typing import DefaultDict, Dict, List

from advent_of_code.helpers.helpers import read_file


def compile_groups(line: str) -> List[List[str]]:
    return [i.split("\n") for i in line.split("\n\n")] 


def group_data(group: List[str]) -> Dict[str, int]:
    map: DefaultDict[str, int] = defaultdict(int)
    map["count"] = len(group)
    for g in group:
        for letter in list(g):
            map[letter] += 1
    return dict(map)


def get_count_from_group(group: Dict[str, int]) -> int:
    return list(group.values()).count(group.get("count")) - 1


def main(filepath: str) -> int:
    lines: str = read_file(filepath=filepath)
    groups: List[str] = compile_groups(lines)
    grouped_data: List[Dict[str, int]] = [group_data(g) for g in groups]
    return sum([get_count_from_group(gg) for gg in grouped_data])

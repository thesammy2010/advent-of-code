import json

from typing import Dict, List, Optional, Set

from advent_of_code.helpers.helpers import read_file_line_by_line


class Rule(object):
    def __init__(self, colour: str = None, bags: Dict[str, int] = None) -> None:
        self.colour: str = colour
        self.bags: List[Dict[str, int]] = bags
    def __repr__(self) -> str:
        return json.dumps(self.__dict__, indent=4)

def create_rule(line: str) -> Rule:
    # "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags"
    r: Rule = Rule()

    colour: str = line.split(" bags")[0]
    # shiny gold

    bags: Dict[str, int] = {}

    unformatted_bags: List[str] = line.split("contain")[1].replace("bags", "").replace("bag", "").split(",")
    
    if not "contain no other bags" in line:
        for ugly_string in unformatted_bags:
            temp: List[str] = ugly_string.split()
            temp_dict: Dict[str, int] = {}
            temp_dict[" ".join(temp[1:]).replace(" .", "")] = int(temp[0])
            bags.update(temp_dict)

    r.colour = colour
    r.bags = bags

    return r


def find_shiny_gold_bag(r: Rule) -> Optional[bool]:
    return "shiny gold" in r.bags.keys()


def do_recursion_stuff(b: Set[str]) -> Set[str]:
    pass

    
def main(filepath: str) -> int:

    files: List[str] = read_file_line_by_line(filepath=filepath)
    rules: List[Rule] = [create_rule(line) for line in files]
    shiny_gold_rules: List[Rule] = [rule for rule in rules if find_shiny_gold_bag(rule)]
    bags: Set[str] = {rule.colour for rule in shiny_gold_rules}

    for b in rules:
        for base_bag in bags:
            if base_bag in b.bags.keys():
                bags.add(b.colour)
                break

    for b in rules:
        for base_bag in bags:
            if base_bag in b.bags.keys():
                bags.add(b.colour)
                break

    for b in rules:
        for base_bag in bags:
            if base_bag in b.bags.keys():
                bags.add(b.colour)
                break

    for b in rules:
        for base_bag in bags:
            if base_bag in b.bags.keys():
                bags.add(b.colour)
                break

    return len(bags)


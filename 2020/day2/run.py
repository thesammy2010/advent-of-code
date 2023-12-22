import sys

import advent_of_code.day2.part1.main as part1
import advent_of_code.day2.part2.main as part2


def main():
    print("Part 1")
    c: int = part1.main(file="advent_of_code/day2/input.txt")
    print(c)
    
    print("Part 2")
    c: int = part2.main(file="advent_of_code/day2/input.txt")
    print(c)
    
    sys.exit(0)


if __name__ == "__main__":
    main()
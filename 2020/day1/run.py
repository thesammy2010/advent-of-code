import sys

import advent_of_code.day1.part1.main as part1
import advent_of_code.day1.part2.main as part2


def main():
    print("Part 1")
    a, b, c = part1.main(file="advent_of_code/day1/input.txt")
    print(a, b, c)
    
    print("Part 2")
    a, b, c, d = part2.main(file="advent_of_code/day1/input.txt")
    print(a, b, c, d)
    
    sys.exit(0)


if __name__ == "__main__":
    main()
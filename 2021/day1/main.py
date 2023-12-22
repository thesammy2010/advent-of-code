def part1():

    # read file
    with open(file="day1/input.txt", mode="r") as f:
        data = f.readlines()

    count: int = 0
    prev: int = 999999
    for idx, number in enumerate(data):
        if idx == 0:
            prev = int(number)
            continue
        if int(number) > prev:
            count += 1

        prev = int(number)

    print("Part 1: %d" % count)


def part2() -> None:
    # read file
    with open(file="day1/input.txt", mode="r") as f:
        data = f.readlines()

    count: int = 0
    prev1 = prev2 = prev3 = 999999
    for idx, n in enumerate(data):
        number = int(n)
        match idx:
            case 0:
                prev3 = number
                continue
            case 1:
                prev2 = number
                continue
            case 2:
                prev1 = number
                continue

        if (number + prev1 + prev2) > (prev1 + prev2 + prev3):
            count += 1

        prev1, prev2, prev3 = number, prev1, prev2

    print("Part 2: %d" % count)


def main() -> None:
    part1()
    part2()


if __name__ == "__main__":
    main()

def part1() -> None:
    # read file
    with open(file="day2/input.txt", mode="r") as f:
        lines = f.readlines()

    x = y = 0
    for line in lines:
        direction, value = line.split()
        step: int = int(value)
        
        match direction:
            case "forward":
                x += step
            case "down":
                y += step
            case "up":
                y -= step

    print("Part 1: %d" % (x * y))


def part2() -> None:
    # read file
    with open(file="day2/input.txt", mode="r") as f:
        lines = f.readlines()

    # x = horizontal
    # y = aim
    # z = depth
    x = y = z = 0
    for line in lines:
        direction, value = line.split()
        step: int = int(value)

        match direction:
            case "forward":
                x += step
                z += step*y

            case "down":
                y += step
            case "up":
                y -= step

    print("Part 2: %d" % (x * z))


def main() -> None:
    part1()
    part2()


if __name__ == "__main__":
    main()

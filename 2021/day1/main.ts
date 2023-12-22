import * as fs from "fs";


function part1(): void {

    // read file
    const data: string = fs.readFileSync("day1/input.txt", "utf-8")
    const lines: string[] = data.split("\n")

    let count: number = 0
    let prev: number = 999999

    for (let idx in lines) {
        const number = parseInt(lines[idx])
        if (parseInt(idx) == 0) {
            prev = number
            continue
        }
        if (isNaN(number)) {
            break
        }
        if (number > prev) {
            count++
        }

        prev = number

    }

    console.log("Part1: %d", count)

}

function part2(): void {
    // read file
    const data: string = fs.readFileSync("day1/input.txt", "utf-8")
    const lines: string[] = data.split("\n")

    let count: number = 0
    let prev1: number = 999999
    let prev2: number = 999999
    let prev3: number = 999999

    for (let idx in lines) {
        let number: number = parseInt(lines[parseInt(idx)])
        if (isNaN(number)) {
            break
        }
        switch (parseInt(idx)) {
            case 0:
                prev3 = number
                continue
            case 1:
                prev2 = number
                continue
            case 2:
                prev1 = number
                continue
        }

        if ((number + prev1 + prev2) > (prev1 + prev2 + prev3)) {
            count++
        }

        prev3 = prev2
        prev2 = prev1
        prev1 = number

    }

    console.log("Part2: %d", count)

}

function main(): void {
    part1()
    part2()
}

main()

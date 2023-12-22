import * as fs from "fs";


function part1(): void {

    // read file
    const data: string = fs.readFileSync("day2/input.txt", "utf-8")
    const lines: string[] = data.split("\n")

    let x: number = 0
    let y: number = 0

    for (let line of lines) {
        if (line == "") {
            break
        }
        const vals: string[] = line.split(" ")
        const direction: string = vals[0]
        const step: number = parseInt(vals[1])
        // console.log(vals, direction)
        switch (direction) {
            case "forward":
                x += step
                break;
            case "down":
                y += step
                break;
            case "up":
                y -= step
                break;
        }
    }

    console.log("Part1: %d", (x*y))

}

function part2(): void {
    // read file
    const data: string = fs.readFileSync("day2/input.txt", "utf-8")
    const lines: string[] = data.split("\n")

    let x: number = 0
    let y: number = 0
    let z: number = 0

    for (let line of lines) {
        if (line == "") {
            break
        }
        const vals: string[] = line.split(" ")
        const direction: string = vals[0]
        const step: number = parseInt(vals[1])
        switch (direction) {
            case "forward":
                x += step
                z += step*y
                break;
            case "down":
                y += step
                break;
            case "up":
                y -= step
                break;
        }
    }

    console.log("Part2: %d", (x*z))

}

function main(): void {
    part1()
    part2()
}

main()


package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
)

func part_one(delta int) int {
    return 1
}

func part_two(delta int) int {
    if delta >= 3 {
        return -1
    }
    return 1
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    var jumps []int
    for scanner.Scan() {
        n, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Invalid number")
            return
        }
        jumps = append(jumps, n)
    }

    pc := 0
    ticks := 0
    for pc >= 0 && pc < len(jumps) {
        delta := jumps[pc]
        jumps[pc] += part_two(delta)
        pc += delta
        ticks += 1
    }

    fmt.Println("Ticks", ticks)
}

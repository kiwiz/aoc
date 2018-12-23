package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type XYS struct {
    X, Y, S int
}

const GRID_SIZE = 300

func calc(cache map[XYS]int, grid *[GRID_SIZE][GRID_SIZE]int, x, y, s int) int {
    if s == 0 {
        return grid[x][y]
    }
    xys := XYS{x, y, s}
    cached_sum, ok := cache[xys]
    if ok {
        return cached_sum
    }

    sum := calc(cache, grid, x, y, s - 1)
    for i := 0; i < s + 1; i++ {
        sum += grid[x + i][y + s]
    }
    for j := 0; j < s; j++ {
        sum += grid[x + s][y + j]
    }
    cache[xys] = sum

    return sum
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    reader := bufio.NewReader(file)
    line, err := reader.ReadString('\n')

    serial, err := strconv.Atoi(strings.Trim(line, "\n "))
    if err != nil {
        fmt.Println(err)
        return
    }

    var grid [GRID_SIZE][GRID_SIZE]int
    for x := 0; x < GRID_SIZE; x++ {
        for y := 0; y < GRID_SIZE; y++ {
            rack_id := (x + 1) + 10
            power := rack_id * (y + 1)
            power += serial
            power *= rack_id
            power = (power / 100) % 10
            power -= 5
            grid[x][y] = power
        }
    }

    max_x := 0
    max_y := 0
    max_size := 0
    max_power := 0

    cache := make(map[XYS]int)
    for s := 0; s < GRID_SIZE; s++ {
        for x := 0; x < GRID_SIZE - s; x++ {
            for y := 0; y < GRID_SIZE - s; y++ {
                power := calc(cache, &grid, x, y, s)
                if power > max_power {
                    max_x = x
                    max_y = y
                    max_size = s
                    max_power = power
                }
            }
        }
    }

    fmt.Printf("Pos: (%d, %d), Size: %d\n", max_x + 1, max_y + 1, max_size + 1)
}

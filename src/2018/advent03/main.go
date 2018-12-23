package main

import (
    "fmt"
    "os"
    "regexp"
    "bufio"
    "strconv"
)

const SIZE = 1000

type Quad struct {
    ID int
    X int
    Y int
    W int
    H int
}

func part_one(grid [][]int, quads []*Quad) {
    sum := 0
    for x := 0; x < SIZE; x++ {
        for y := 0; y < SIZE; y++ {
            if grid[x][y] >= 2 {
                sum += 1
            }
        }
    }

    fmt.Printf("Sum: %d\n", sum)
}

func part_two(grid [][]int, quads []*Quad) {
    id := 0
    for _, quad := range quads {
        ok := true
        for x := 0; x < quad.W && ok; x++ {
            for y := 0; y < quad.H && ok; y++ {
                if grid[quad.X + x][quad.Y + y] > 1 {
                    ok = false
                    break
                }
            }
        }
        if ok {
            id = quad.ID
            break
        }
    }

    fmt.Printf("ID: %d\n", id)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    line_re := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

    quads := make([]*Quad, 0)
    for scanner.Scan() {
        match := line_re.FindStringSubmatch(scanner.Text())
        id, _ := strconv.Atoi(match[1])
        x, _ := strconv.Atoi(match[2])
        y, _ := strconv.Atoi(match[3])
        w, _ := strconv.Atoi(match[4])
        h, _ := strconv.Atoi(match[5])
        quads = append(quads, &Quad{id, x, y, w, h})
    }

    grid := make([][]int, SIZE)
    for i := 0; i < SIZE; i++ {
        grid[i] = make([]int, SIZE)
    }

    for _, quad := range quads {
        for x := 0; x < quad.W; x++ {
            for y := 0; y < quad.H; y++ {
                grid[quad.X + x][quad.Y + y] += 1
            }
        }
    }

    part_one(grid, quads)
    part_two(grid, quads)
}

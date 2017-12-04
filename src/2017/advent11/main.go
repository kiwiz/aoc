package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "math"
)

type cube_coord struct {
    x int
    y int
    z int
}

var direction_map map[string]cube_coord = map[string]cube_coord{
    //      x   y   z
    "n":  { 0,  1, -1},
    "s":  { 0, -1,  1},
    "nw": {-1,  1,  0},
    "ne": { 1,  0, -1},
    "sw": {-1,  0,  1},
    "se": { 1, -1,  0},
}

func move(pos cube_coord, direction string) cube_coord {
    delta := direction_map[direction]

    return cube_coord{pos.x + delta.x, pos.y + delta.y, pos.z + delta.z}
}

func distance(a, b cube_coord) float64 {
    return (math.Abs(float64(a.x - b.x)) + math.Abs(float64(a.y - b.y)) + math.Abs(float64(a.z - b.z))) / 2
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()

    var curr_pos cube_coord
    var long_pos cube_coord

    for _, direction := range strings.Split(scanner.Text(), ",") {
        curr_pos = move(curr_pos, direction)
        if distance(cube_coord{}, curr_pos) > distance(cube_coord{}, long_pos) {
            long_pos = curr_pos
        }
    }

    moves := distance(cube_coord{}, curr_pos)
    max_moves := distance(cube_coord{}, long_pos)

    fmt.Println("Moves", moves, max_moves)
}

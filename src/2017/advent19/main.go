package main

import (
    "os"
    "fmt"
    "bufio"
    "errors"
)

const (
    DIR_S = iota
    DIR_W = iota
    DIR_N = iota
    DIR_E = iota
)

var dirs [][2]int = [][2]int{
    { 0,  1},
    {-1,  0},
    { 0, -1},
    { 1,  0},
}

type packet struct {
    x int
    y int
    dir int
    visited []rune
    steps int
}

func find_start(world []string) (int, error) {
    for i := 0; i < len(world[0]); i++ {
        if world[0][i] == '|' {
            return i, nil
        }
    }

    return 0, errors.New("Start point not found")
}

func valid_move(world []string, p *packet, delta [2]int) bool {
    nx := p.x + delta[0]
    ny := p.y + delta[1]

    if nx < 0 || nx >= len(world[0]) {
        return false
    }
    if ny < 0 || ny >= len(world) {
        return false
    }

    if world[ny][nx] == ' ' {
        return false
    }

    return true
}

func tick(world []string, p *packet) bool {
    open_dirs := []int{
        p.dir,
        (p.dir + 1) % 4,
        (p.dir + 3) % 4,
    }

    for _, dir := range open_dirs {
        delta := dirs[dir]

        if valid_move(world, p, delta) {
            p.x += delta[0]
            p.y += delta[1]
            p.dir = dir
            p.steps += 1
            b := world[p.y][p.x]

            if b >= 'A' && b <= 'Z' {
                p.visited = append(p.visited, rune(b))
            }
            return true
        }
    }

    return false
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    var world []string

    for scanner.Scan() {
        world = append(world, scanner.Text())
    }

    var p packet

    start_x, err := find_start(world)
    if err != nil {
        return
    }

    p.x = start_x
    p.y = -1

    for tick(world, &p) {
    }

    fmt.Println("Visited", string(p.visited), p.steps)
}

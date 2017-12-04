package main

import (
    "os"
    "fmt"
    "bufio"
)

const (
    ST_CLEAN = iota
    ST_WEAKENED = iota
    ST_INFECTED = iota
    ST_FLAGGED =  iota
    N_STATES = iota
)

type world struct {
    data map[string]int
}

func NewWorld() *world {
    var w world
    w.data = make(map[string]int)
    return &w
}

func hash(x, y int) string {
    return fmt.Sprintf("%d %d", x, y)
}

func (w *world) At(x, y int) int {
    return w.data[hash(x, y)]
}

func (w *world) Set(x, y int, val int) {
    w.data[hash(x, y)] = val
}

func (w *world) Process(x, y int) int {
    h := hash(x, y)
    w.data[h] = (w.data[h] + 1) % N_STATES

    return w.data[h]
}

var dirs [][2]int = [][2]int{
    { 0, -1},
    { 1,  0},
    { 0,  1},
    {-1,  0},
}

type carrier struct {
    x int
    y int
    dir int
}

func burst(w *world, c *carrier) int {
    switch w.At(c.x, c.y) {
    case ST_CLEAN:
        c.dir = (c.dir + 3) % 4
    case ST_WEAKENED:
    case ST_INFECTED:
        c.dir = (c.dir + 1) % 4
    case ST_FLAGGED:
        c.dir = (c.dir + 2) % 4
    }
    s := w.Process(c.x, c.y)

    c.x += dirs[c.dir][0]
    c.y += dirs[c.dir][1]

    return s
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    w := NewWorld()

    y := 0
    for scanner.Scan() {
        for x, i := range scanner.Text() {
            state := ST_CLEAN
            if i == '#' {
                state = ST_INFECTED
            }
            w.Set(x, y, state)
        }
        y += 1
    }

    var c carrier
    c.x = y / 2
    c.y = y / 2

    acc := 0
    for i := 0; i < 10000000; i++ {
        if burst(w, &c) == ST_INFECTED {
            acc += 1
        }
    }

    fmt.Println("Bursts", acc)
}

package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

const (
    DOWN = iota
    UP = iota
)

type layer struct {
    pos int
    size int
    direction int
}

func parse_line(str string) (int, *layer) {
    parts := strings.Split(str, ": ")
    if len(parts) != 2 {
        fmt.Println("Invalid line")
        return 0, &layer{0, 0, DOWN}
    }

    pos, err := strconv.Atoi(parts[0])
    if err != nil {
        fmt.Println("Invalid pos")
    }

    size, err := strconv.Atoi(parts[1])
    if err != nil {
        fmt.Println("Invalid size")
    }

    return pos, &layer{0, size, DOWN}
}

func tick(firewall map[int]*layer) {
    for _, l := range firewall {
        if l.direction == UP && l.pos == 0 {
            l.direction = DOWN
        }
        if l.direction == DOWN && l.pos == l.size - 1 {
            l.direction = UP
        }

        switch(l.direction) {
        case UP:
            l.pos -= 1
            break
        case DOWN:
            l.pos += 1
        }
    }
}

func clone(firewall map[int]*layer) map[int]*layer {
    firewall_copy := make(map[int]*layer)
    for k, v := range firewall {
        l := *v
        firewall_copy[k] = &l
    }

    return firewall_copy
}

func run(firewall map[int]*layer, depth int) bool {
    i := 0
    for i <= depth {
        l, ok := firewall[i]
        if ok && l.pos == 0 {
            return false
        }

        tick(firewall)
        i += 1
    }

    return true
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    firewall := make(map[int]*layer)
    last_depth := 0
    for scanner.Scan() {
        depth, l := parse_line(scanner.Text())
        firewall[depth] = l
        if depth > last_depth {
            last_depth = depth
        }
    }

    i := 0
    for {
        if run(clone(firewall), last_depth) {
            break
        }

        i += 1
        tick(firewall)
    }

    fmt.Println("Delay", i)
}

package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type node struct {
    a int
    b int
    used bool
}

type path []*node

func parse_line(str string) *node {
    parts := strings.Split(str, "/")
    if len(parts) != 2 {
        return nil
    }

    a, err := strconv.Atoi(parts[0])
    if err != nil {
        return nil
    }
    b, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil
    }

    return &node{a, b, false}
}

func start(all map[int]path) (path, int) {
    var p path
    best_path, best_strength, _ := traverse(all, p, 0, 0)
    return best_path, best_strength
}

func clone(p path) path {
    var new_path path
    return append(new_path, p...)
}

func traverse(all map[int]path, p path, seek, strength int) (path, int, int) {
    var best_path path
    best_strength := 0
    best_len := 0

    for _, n := range all[seek] {
        if n.used {
            continue
        }

        next_seek := n.a
        if n.a == seek {
            next_seek = n.b
        }
        next_strength := strength + n.a + n.b

        n.used = true
        p = append(p, n)

        new_path, new_strength, new_len := traverse(all, p, next_seek, next_strength)
        if (new_len == best_len && new_strength > best_strength) || (new_len > best_len) {
            best_len = new_len
            best_path = clone(new_path)
            best_strength = new_strength
        }

        p = p[0:len(p) - 1]
        n.used = false
    }

    if best_path != nil {
        return best_path, best_strength, len(best_path)
    }

    return clone(p), strength, len(p)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    all := make(map[int]path)
    for scanner.Scan() {
        n := parse_line(scanner.Text())

        all[n.a] = append(all[n.a], n)
        if n.a != n.b {
            all[n.b] = append(all[n.b], n)
        }
    }

    _, best_strength := start(all)

    fmt.Println("Strength", best_strength)
}

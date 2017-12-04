package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func group(start int, pipes map[int][]int, visited map[int]bool) int {
    open := []int{start}
    acc := 0
    for len(open) > 0 {
        id := open[0]
        visited[id] = true
        open = append(open[:0], open[1:]...)
        acc += 1

        for _, conn := range pipes[id] {
            _, ok := visited[conn]
            if !ok {
                open = append(open, conn)
            }
            visited[conn] = true
        }
    }

    return acc
}

func parse_line(str string) (int, []int) {
    parts := strings.Split(str, " <-> ")
    if len(parts) != 2 {
        fmt.Println("Invalid line")
        return 0, nil
    }

    id, err := strconv.Atoi(parts[0])
    if err != nil {
        fmt.Println("Invalid id")
    }

    var connections []int
    for _, val := range strings.Split(parts[1], ", ") {
        conn, err := strconv.Atoi(val)
        if err != nil {
            fmt.Println("Invalid conn")
            continue
        }
        connections = append(connections, conn)
    }

    return id, connections
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    pipes := make(map[int][]int)
    for scanner.Scan() {
        id, connections := parse_line(scanner.Text())

        pipes[id] = connections
    }

    visited := make(map[int]bool)
    acc := 0
    for id, _ := range pipes {
        if visited[id] {
            continue
        }
        group(id, pipes, visited)
        acc += 1
    }

    fmt.Println("Groups", acc)
}

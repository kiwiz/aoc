package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "sort"
)

func matches(id string) (bool, bool) {
    counts := make(map[rune]int)

    two := false
    three := false
    for _, r := range id {
        counts[r] += 1
    }
    for _, n := range counts {
        if n == 2 {
            two = true
        }
        if n == 3 {
            three = true
        }
    }

    return two, three
}

func part_one(ids []string) {
    twos := 0
    threes := 0

    for _, id := range ids {
        two, three := matches(id)
        if two {
            twos += 1
        }
        if three {
            threes += 1
        }
    }

    fmt.Printf("Checksum: %d\n", twos * threes)
}

func part_two(ids []string) {
    sort.Strings(ids)

    matches := ""
    for i := 0; i < len(ids) - 1; i++ {
        buf := make([]byte, 0, 10)
        for j := 0; j < len(ids[i]); j++ {
            if ids[i][j] == ids[i + 1][j] {
                buf = append(buf, ids[i][j])
            }
        }
        if len(buf) == len(ids[i]) - 1 {
            matches = string(buf)
            break
        }
    }

    fmt.Printf("Same: %s\n", matches)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    ids := make([]string, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ids = append(ids, strings.Trim(scanner.Text(), "\n"))
    }

    part_one(ids)
    part_two(ids)
}

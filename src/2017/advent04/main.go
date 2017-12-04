package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "sort"
)

func part_one(words []string) bool {
    var seen = make(map[string]bool)

    for _, word := range words {
        _, ok := seen[word]
        if ok {
            return false
        }
        seen[word] = true
    }

    return true
}

func hash(str string) string {
    arr := strings.Split(str, "")
    sort.Strings(arr)
    return strings.Join(arr, "")
}

func part_two(words []string) bool {
    var seen = make(map[string]bool)
    var chars = make(map[string]bool)

    for _, word := range words {
        key := hash(word)
        _, s_ok := seen[word]
        _, c_ok := chars[key]

        if s_ok || c_ok {
            return false
        }
        seen[word] = true
        chars[key] = true
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

    acc := 0
    for scanner.Scan() {
        ok := part_two(strings.Split(scanner.Text(), " "))
        if ok {
            acc += 1
        }
    }

    fmt.Println("Count", acc)
}

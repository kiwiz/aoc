package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

const ITER = 1000000000

func find(line []rune, char rune) int {
    for i, c := range line {
        if c == char {
            return i
        }
    }

    return -1
}

func dance(line []rune, commands []string) []rune {
    for _, command := range commands {
        switch(command[0]) {
        case 's':
            dist, err := strconv.Atoi(command[1:])
            if err != nil {
                fmt.Println("Invalid dist")
            }
            line = append(line[len(line) - dist:], line[:len(line) - dist]...)
            break
        case 'x':
            parts := strings.Split(command[1:], "/")
            if len(parts) != 2 {
                fmt.Println("Invalid indices")
            }
            a, err := strconv.Atoi(parts[0])
            if err != nil {
                fmt.Println("Invalid param")
            }
            b, err := strconv.Atoi(parts[1])
            if err != nil {
                fmt.Println("Invalid param")
            }
            t := line[a]
            line[a] = line[b]
            line[b] = t
            break
        case 'p':
            parts := strings.Split(command[1:], "/")
            if len(parts) != 2 {
                fmt.Println("Invalid keys")
            }
            a := find(line, rune(parts[0][0]))
            b := find(line, rune(parts[1][0]))
            t := line[a]
            line[a] = line[b]
            line[b] = t
            break
        }
    }

    return line
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()

    commands := strings.Split(scanner.Text(), ",")

    line := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}

    seen := make(map[string]int)
    skip := true
    for i := 0; i < ITER; i++ {
        line = dance(line, commands)
        n, ok := seen[string(line)]

        if ok && skip {
            count := i - n
            for i + count < ITER {
                i += count
            }
            skip = false
        }
        seen[string(line)] = i
    }

    fmt.Println(string(line))
}

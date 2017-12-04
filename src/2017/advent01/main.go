package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
)

func eq(a, b byte) bool {
    return a == b && '0' <= a && '9' >= b
}

func part_one(input string, i int) bool {
    j := (len(input) + i - 1) % len(input)
    return eq(input[i], input[j])
}

func part_two(input string, i int) bool {
    j := ((len(input) / 2) + i) % len(input)
    return eq(input[i], input[j])
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    reader := bufio.NewReader(file)
    input, err := reader.ReadString('\n')
    if err != nil {
        return
    }
    input = strings.Trim(input, " \t\n\r")

    var acc int
    for i := 0; i < len(input); i++ {
        pass := part_two(input, i)
        if pass {
            acc += (int)(input[i] - '0')
        }
    }

    fmt.Println("Sum", acc);
}

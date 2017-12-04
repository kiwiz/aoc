package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    reader := bufio.NewReader(file)
    input, _ := reader.ReadString('\n')

    floor := 0
    for _, direction := range input {
        if direction == '(' {
            floor += 1
        }
        if direction == ')' {
            floor -= 1
        }
    }

    fmt.Println("Floor", floor)
}

package main

import (
    "fmt"
    "os"
    "bufio"
)

var board [3][3]int = [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    x := 1
    y := 1
    var output []int
    for scanner.Scan() {
        for _, direction := range scanner.Text() {
            if direction == 'U' {
                y -= 1
            }
            if direction == 'D' {
                y += 1
            }
            if direction == 'L' {
                x -= 1
            }
            if direction == 'R' {
                x += 1
            }

            if x >= 3 {
                x = 2
            }
            if x < 0 {
                x = 0
            }
            if y >= 3 {
                y = 2
            }
            if y < 0 {
                y = 0
            }
        }

        output = append(output, board[y][x])
    }

    fmt.Println("Code", output)
}

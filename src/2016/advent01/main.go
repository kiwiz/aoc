package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math"
)

var direction_map [][2]int = [][2]int{
    // x, y
    { 0,  1}, // North
    { 1,  0}, // East
    { 0, -1}, // South
    {-1,  0}, // West
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
        fmt.Println("No data")
        return
    }

    instructions := strings.Split(input, ",")

    direction := 0
    x := 0
    y := 0
    for _, instruction := range instructions {
        instruction = strings.Trim(instruction, " \t\r\n")
        turn := instruction[0]
        distance, err := strconv.Atoi(instruction[1:])
        if err != nil {
            fmt.Println("Invalid number", distance)
        }

        if turn == 'L' {
            direction -= 1
        }
        if turn == 'R' {
            direction += 1
        }
        direction = (direction + len(direction_map)) % len(direction_map)

        x += direction_map[direction][0] * distance
        y += direction_map[direction][1] * distance
    }

    fmt.Println("Distance", math.Abs(float64(x)) + math.Abs(float64(y)))
}

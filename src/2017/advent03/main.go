package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "math"
)

func num_to_pos(number int) (int, int) {
    if number <= 1 {
        return 0, 0
    }

    length := 1
    for int(math.Pow(float64(length), 2)) < number {
        length += 2
    }

    index := number - int(math.Pow(float64(length - 2), 2)) - 1
    side := index / (length - 1)
    height := index % (length - 1)
    distance := height - (length / 2 - 1)

    if side == 0 {
        return length / 2, distance
    }
    if side == 1 {
        return -distance, length / 2
    }
    if side == 2 {
        return -length / 2, -distance
    }
    if side == 3 {
        return distance, -length / 2
    }

    return 0, 0
}

func pos_to_num(x, y int) int {
    if x == 0 && y == 0 {
        return 1
    }

    length := int(math.Max(math.Abs(float64(x)), math.Abs(float64(y))))
    base := int(math.Pow(float64(length) * 2 - 1, 2))

    if x == length && y > -length {
        return base + length * 0 + (y + length)
    }
    if y == length && x < length {
        return base + length * 2 + length - x
    }
    if x == -length && y < length {
        return base + length * 4 + length - y
    }
    if y == -length && x > -length {
        return base + length * 6 + x + length
    }

    return 1
}

var surrounding [][2]int = [][2]int{
    { 1,  0},
    { 1,  1},
    { 0,  1},
    {-1,  1},
    {-1,  0},
    {-1, -1},
    { 0, -1},
    { 1, -1},
}

func part_one(number int) int {
    x, y := num_to_pos(number)

    return int(math.Max(math.Abs(float64(x)) + math.Abs(float64(y)), 0))
}

func part_two(number int) int {
    var mem []int;
    curr_value := 1
    mem = append(mem, curr_value)

    for number > curr_value {
        x, y := num_to_pos(len(mem) + 1)
        curr_value = 0
        for _, coord := range surrounding {
            n := pos_to_num(x + coord[0], y + coord[1])
            if n <= len(mem) {
                curr_value += mem[n - 1]
            }
        }

        mem = append(mem, curr_value)
    }

    return curr_value
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()

    input := scanner.Text()
    number, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("No data")
        return
    }

    distance := part_two(number)

    fmt.Println("Distance", distance)
}

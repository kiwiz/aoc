package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    acc := 0
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), "x")
        if len(parts) != 3 {
            fmt.Println("Malformed line", parts)
            continue
        }

        w, err := strconv.Atoi(parts[0])
        if err != nil {
            fmt.Println("Invalid w", parts[0])
            continue
        }
        h, err := strconv.Atoi(parts[1])
        if err != nil {
            fmt.Println("Invalid h", parts[1])
            continue
        }
        l, err := strconv.Atoi(parts[2])
        if err != nil {
            fmt.Println("Invalid l", parts[2])
            continue
        }

        acc += 2*l*w + 2*w*h + 2*h*l + int(math.Min(float64(l*w), math.Min(float64(w*h), float64(h*l))))
    }

    fmt.Println("Length", acc)
}

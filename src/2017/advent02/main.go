package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
    "math"
    "strings"
    "sort"
)

func even_div(a, b int) bool {
    div := a / b
    return div * b == a
}

func numerify(numbers []string) []int {
    var ret []int

    for _, number := range numbers {
        n, err := strconv.Atoi(number)
        if err != nil {
            fmt.Println("Invalid number", number)
        }
        ret = append(ret, n)
    }

    return ret
}

func part_one(numbers []int) int {
    lo := math.MaxInt32
    hi := 0

    for _, n := range numbers {
        if n < lo {
            lo = n
        }
        if n > hi {
            hi = n
        }
    }

    return hi - lo
}

func part_two(numbers []int) int {
    for i := 1; i < len(numbers); i++ {
        for j := 0; j < i; j++ {
            if even_div(numbers[i], numbers[j]) {
                return numbers[i] / numbers[j]
            }
        }
    }

    panic("No evenly divisible pair found")
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
    }

    scanner := bufio.NewScanner(file)

    var acc int
    for scanner.Scan() {
        numbers := numerify(strings.Split(scanner.Text(), "\t"))
        sort.Ints(numbers)
        acc += part_two(numbers)
    }

    fmt.Println("Sum", acc)
}

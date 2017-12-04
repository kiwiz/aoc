package main

import (
    "os"
    "strconv"
    "bufio"
    "fmt"
    "strings"
)

func gen_a(x int64) int64 {
    return (x * 16807) % 2147483647
}

func gen_b(x int64) int64 {
    return (x * 48271) % 2147483647
}

func part_one(a, b int64) int {
    acc := 0
    for i := 0; i < 40000000; i++ {
        a = gen_a(a)
        b = gen_b(b)
        if a & 0xffff == b & 0xffff {
            acc += 1
        }
    }

    return acc
}

func part_two(a, b int64) int {
    acc := 0
    for i := 0; i < 5000000; i++ {
        for {
            a = gen_a(a)
            if a % 4 == 0 {
                break
            }
        }
        for {
            b = gen_b(b)
            if b % 8 == 0 {
                break
            }
        }
        if a & 0xffff == b & 0xffff {
            acc += 1
        }
    }

    return acc
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    a_str := strings.Fields(scanner.Text())
    scanner.Scan()
    b_str := strings.Fields(scanner.Text())

    a_, err := strconv.Atoi(a_str[len(a_str) - 1])
    if err != nil {
        fmt.Println("Invalid A")
        return
    }
    b_, err := strconv.Atoi(b_str[len(b_str) - 1])
    if err != nil {
        fmt.Println("Invalid B")
        return
    }

    a, b := int64(a_), int64(b_)
    acc := part_two(a, b)

    fmt.Println(acc)
}

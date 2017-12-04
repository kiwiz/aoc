package main

import (
    "os"
    "io"
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "crypto/sha1"
)

func highest_index(banks []int) int {
    index := 0
    max := 0
    for i, val := range banks {
        if val > max {
            max = val
            index = i
        }
    }

    return index
}

func distribute(index int, banks []int) {
    blocks := banks[index]
    banks[index] = 0

    i := index
    for blocks > 0 {
        i = (i + 1) % len(banks)
        banks[i] += 1
        blocks -=1
    }
}

func hash(banks []int) string {
    h := sha1.New()
    for _, i := range banks {
        io.WriteString(h, string(i))
    }
    return string(h.Sum(nil))
}

func process(banks []int) (int, int) {
    cycles := 0
    var seen map[string]int = make(map[string]int)
    for {
        h := hash(banks)
        _, ok := seen[h]
        if ok {
            return seen[h], cycles
        }
        seen[h] = cycles
        i := highest_index(banks)
        distribute(i, banks)
        cycles += 1
    }
}

func part_one(banks []int) int {
    _, cycles := process(banks)
    return cycles
}
func part_two(banks []int) int {
    last, curr := process(banks)
    return curr - last
}

func main() {
    var banks []int

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

    for _, text := range strings.Split(strings.Trim(input, "\n"), "\t") {
        n, err := strconv.Atoi(text)
        if err != nil {
            fmt.Println("Invalid number")
            return
        }

        banks = append(banks, n)
    }

    cycles := part_two(banks)

    fmt.Println("Cycles", cycles)
}

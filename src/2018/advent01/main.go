package main;

import (
    "fmt"
    "os"
    "regexp"
    "bufio"
    "strconv"
)

func part_one(deltas []int) {
    sum := 0
    for _, n := range deltas {
        sum += n
    }

    fmt.Printf("Sum: %d\n", sum)
}

func part_two(deltas []int) {
    dupe := 0
    cache := make(map[int]bool)
    sum := 0
    i := 0
    for {
        sum += deltas[i]
        if cache[sum] {
            dupe = sum
            break
        }
        cache[sum] = true

        i = (i + 1) % len(deltas)
    }

    fmt.Printf("Dupe: %d\n", dupe)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    var line_re = regexp.MustCompile(`([+-])(\d+)`)

    deltas := make([]int, 0, 10)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        matches := line_re.FindStringSubmatch(scanner.Text())
        n, err := strconv.Atoi(matches[2])
        if err != nil {
            fmt.Println("Invalid number: %s\n", matches[2])
            n = 0
        }

        if matches[1] == "-" {
            n *= -1
        }
        deltas = append(deltas, n)
    }

    part_one(deltas)
    part_two(deltas)
}

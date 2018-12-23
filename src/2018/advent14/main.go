package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
)

func to_seq(num int) []int {
    i := 0
    seq := make([]int, 0)
    for ; math.Pow(10, float64(i + 1)) <= float64(num); i++ {}

    for ; i >= 0; i-- {
        val := (num / int(math.Pow(10, float64(i)))) % 10
        seq = append(seq, val)
    }

    return seq
}

func cmp(a, b []int) bool {
    l := int(math.Min(float64(len(a)), float64(len(b))))

    for i := 0; i < l; i++ {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func proc(scores []int, idx_a, idx_b, n int) ([]int, int, int) {
    for len(scores) < n {
        new_scores := to_seq(scores[idx_a] + scores[idx_b])
        scores = append(scores, new_scores...)

        if len(new_scores) < 2 {
            new_scores = append(new_scores, 0)
        }

        idx_a = (idx_a + scores[idx_a] + 1) % len(scores)
        idx_b = (idx_b + scores[idx_b] + 1) % len(scores)
    }

    return scores, idx_a, idx_b
}

func part_one(scores []int, idx_a, idx_b, num int) {
    fmt.Println(scores[num:num + 10])
}

func part_two(scores []int, idx_a, idx_b, num int) {
    seq := to_seq(num)

    i := 0
    for {
        scores, idx_a, idx_b = proc(scores, idx_a, idx_b, len(seq) + i)
        if cmp(scores[i:], seq) {
            break
        }
        i++
    }

    fmt.Printf("Idx: %d\n", i)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    reader := bufio.NewReader(file)
    text, err := reader.ReadString('\n')
    num, err := strconv.Atoi(strings.Trim(text, "\n "))

    if err != nil {
        fmt.Println("Invalid number")
        return
    }

    var scores = []int{3, 7}
    idx_a := 0
    idx_b := 1

    scores, idx_a, idx_b = proc(scores, idx_a, idx_b, num + 10)
    part_one(scores, idx_a, idx_b, num)
    part_two(scores, idx_a, idx_b, num)
}

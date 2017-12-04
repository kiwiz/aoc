package main

import (
    "os"
    "bufio"
    "fmt"
    "encoding/hex"
)

func round(knot *[256]byte, seq []byte, i, skip int) (int, int) {
    for _, num := range seq {
        for x, y := i, i + int(num) - 1; x < y; x, y = x + 1, y - 1 {
            x_ := x % len(knot)
            y_ := y % len(knot)

            knot[x_], knot[y_] = knot[y_], knot[x_]
        }

        i = (i + int(num) + skip) % len(knot)
        skip += 1
    }

    return i, skip
}

func main() {
    knot := [256]byte{}
    for i := 0; i < len(knot); i++ {
        knot[i] = byte(i)
    }

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    skip := 0
    idx := 0

    scanner := bufio.NewScanner(file)
    scanner.Scan()

    var seq []byte
    for _, num := range scanner.Text() {
        seq = append(seq, byte(num))
    }
    seq = append(seq, []byte{17, 31, 73, 47, 23}...)

    for i := 0; i < 64; i++ {
        idx, skip = round(&knot, seq, idx, skip)
    }

    dense := [16]byte{}
    for i := 0; i < len(knot); i++ {
        dense[i / 16] ^= knot[i]
    }

    hash := hex.EncodeToString(dense[:])
    fmt.Println("Hash", hash)
}

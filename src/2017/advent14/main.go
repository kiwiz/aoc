package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
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

func hash(input string) [16]byte {
    knot := [256]byte{}
    for i := 0; i < len(knot); i++ {
        knot[i] = byte(i)
    }

    skip := 0
    idx := 0

    var seq []byte
    for _, num := range input {
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

    return dense
}

var probes [][2]int = [][2]int{
    { 1,  0},
    { 0,  1},
    { 0, -1},
    {-1,  0},
}

type point struct {
    x int
    y int
}

func consume_region(x, y int, disk *[128][128]bool) {
    open := []point{point{x, y}}

    for len(open) > 0 {
        p := open[len(open) - 1]
        open = open[:len(open) - 1]

        if disk[p.y][p.x] {
            disk[p.y][p.x] = false
            for i := 0; i < len(probes); i++ {
                p_ := point{p.x + probes[i][1], p.y + probes[i][0]}
                if p_.x < 128 && p_.x >= 0 && p_.y < 128 && p_.y >= 0 {
                    open = append(open, p_)
                }
            }
        }
    }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    if !scanner.Scan() {
        fmt.Println("Invalid data")
        return
    }
    input := scanner.Text()

    disk := [128][128]bool{}
    for i := 0; i < len(disk); i++ {
        for j, chunk := range hash(input + "-" + strconv.Itoa(i)) {
            var bit byte = 0x80
            var k byte = 0x0
            for ; k < 8; k++ {
                if chunk & (bit >> k) == 0 {
                    continue
                }

                disk[i][j * 8 + int(k)] = true
            }
        }
    }

    blocks := 0
    for i := 0; i < len(disk); i++ {
        for j := 0; j < len(disk[i]); j++ {
            if disk[i][j] {
                blocks += 1
            }
        }
    }

    regions := 0
    for i := 0; i < len(disk); i++ {
        for j := 0; j < len(disk[i]); j++ {
            if disk[i][j] {
                consume_region(j, i, &disk)
                regions += 1
            }
        }
    }

    fmt.Println("Usage", blocks, regions)
}

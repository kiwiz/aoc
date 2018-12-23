package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strconv"
)

type Marble struct {
    Val int
    Left *Marble
    Right *Marble
}

func (m *Marble) Append(o *Marble) *Marble {
    o.Left = m
    o.Right = m.Right
    o.Left.Right = o
    o.Right.Left = o
    return o
}

func (m *Marble) Remove() *Marble {
    m.Left.Right = m.Right
    m.Right.Left = m.Left
    o := m.Right
    m.Left = nil
    m.Right = nil
    return o
}

func NewRing() *Marble {
    marble := Marble{}
    marble.Left = &marble
    marble.Right = &marble

    return &marble
}

func compute(players, max_val int) {
    scores := make([]int, players)
    ring := NewRing()
    curr_player := 0
    for val := 1; val <= max_val; val++ {
        if val % 23 == 0 {
            marble := ring.Left.Left.Left.Left.Left.Left.Left
            scores[curr_player] += val + marble.Val
            ring = marble.Remove()
        } else {
            ring = ring.Right.Append(&Marble{Val: val})
        }

        curr_player = (curr_player + 1) % players
    }

    max_score := 0
    for _, score := range scores {
        if score > max_score {
            max_score = score
        }
    }

    fmt.Printf("Score: %d\n", max_score)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    line_re := regexp.MustCompile(`(\d+) players; last marble is worth (\d+) points`)

    reader := bufio.NewReader(file)
    line, err := reader.ReadString('\n')
    if len(line) == 0 {
        fmt.Println(err)
        return
    }

    matches := line_re.FindStringSubmatch(line)
    if matches == nil {
        fmt.Println("Invalid input")
        return
    }

    players, err := strconv.Atoi(matches[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    max_val, err := strconv.Atoi(matches[2])
    // max_val *= 100
    if err != nil {
        fmt.Println(err)
        return
    }

    compute(players, max_val)
    compute(players, max_val * 100)
}

package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "math"
    "strconv"
)

type Point struct {
    X int
    Y int
}

type Bounds struct {
    MinX int
    MinY int
    MaxX int
    MaxY int
    Set bool
}

func (b *Bounds) Add(x, y int) {
    if !b.Set || b.MinX > x {
        b.MinX = x
    }
    if !b.Set || b.MaxX < x {
        b.MaxX = x
    }
    if !b.Set || b.MinY > y {
        b.MinY = y
    }
    if !b.Set || b.MaxY < y {
        b.MaxY = y
    }
    b.Set = true
}

func (b *Bounds) W() int {
    return b.MaxX - b.MinX
}
func (b *Bounds) H() int {
    return b.MaxY - b.MinY
}

func (p *Point) Dist(x, y int) int {
    return int(math.Abs(float64(p.X) - float64(x)) + math.Abs(float64(p.Y) - float64(y)))
}

func closest(points []*Point, x, y int) int {
    min_dist := math.MaxInt32
    count := 0
    idx := 0
    for i, point := range points {
        dist := point.Dist(x, y)
        if dist < min_dist {
            idx = i
            min_dist = dist
            count = 1
        } else if dist == min_dist {
            count += 1
        }
    }

    if count > 1 {
        return -1
    }
    return idx
}

func part_one(bounds *Bounds, points []*Point) {
    scores := make([]int, len(points))
    for x := bounds.MinX; x < bounds.MaxX; x++ {
        for y := bounds.MinY; y < bounds.MaxY; y++ {
            i := closest(points, x, y)
            if i >= 0 {
                scores[i] += 1
            }
        }
    }

    max_score := 0
    for _, score := range scores {
        if score > max_score {
            max_score = score
        }
    }

    fmt.Printf("Size: %d\n", max_score)
}

func part_two(bounds *Bounds, points []*Point, limit int) {
    deltaX := limit * len(points) / bounds.W()
    deltaY := limit * len(points) / bounds.H()
    new_bounds := Bounds{
        MinX: bounds.MaxX - deltaX,
        MaxX: bounds.MinX + deltaX,
        MinY: bounds.MaxY - deltaY,
        MaxY: bounds.MinY + deltaY,
    }

    score := 0
    for x := new_bounds.MinX; x < new_bounds.MaxX; x++ {
        for y := new_bounds.MinY; y < new_bounds.MaxY; y++ {
            dist := 0
            for _, point := range points {
                dist += point.Dist(x, y)
            }

            if dist < limit {
                score += 1
            }
        }
    }

    fmt.Printf("Size: %d\n", score)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    points := make([]*Point, 0)
    scanner := bufio.NewScanner(file)
    bounds := &Bounds{}
    for scanner.Scan() {
        line := strings.Trim(scanner.Text(), "\n ")

        parts := strings.Split(line, ", ")
        if len(parts) != 2 {
            continue
        }

        x, err := strconv.Atoi(parts[0])
        if err != nil {
            continue
        }
        y, err := strconv.Atoi(parts[1])
        if err != nil {
            continue
        }

        bounds.Add(x, y)
        points = append(points, &Point{x, y})
    }

    part_one(bounds, points)
    part_two(bounds, points, 10000)
}


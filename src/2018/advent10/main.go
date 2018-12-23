package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strconv"
    "math"
)

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
func (b *Bounds) Size() int {
    return b.W() * b.H()
}

type Point struct {
    X int
    Y int
    VelX int
    VelY int
}

func (p *Point) Dist(x, y int) int {
    return int(math.Abs(float64(p.X) - float64(x)) + math.Abs(float64(p.Y) - float64(y)))
}

func tick(points []*Point) *Bounds {
    bounds := Bounds{}
    for _, point := range points {
        point.X += point.VelX
        point.Y += point.VelY
        bounds.Add(point.X, point.Y)
    }

    return &bounds
}

func display(points []*Point, bounds *Bounds) {
    w := bounds.W() + 1
    h := bounds.H() + 1
    grid := make([][]rune, h)
    for i := 0; i < h; i++ {
        grid[i] = make([]rune, w)
        for j := 0; j < w; j++ {
            grid[i][j] = ' '
        }
    }
    for _, point := range points {
        grid[point.Y - bounds.MinY][point.X - bounds.MinX] = '#'
    }

    for _, row := range grid {
        fmt.Println(string(row))
    }
    fmt.Println("---")
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    line_re := regexp.MustCompile(`position=<\s*(\-?\d+),\s*(\-?\d+)> velocity=<\s*(\-?\d+),\s*(\-?\d+)>`)

    points := make([]*Point, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        parts := line_re.FindStringSubmatch(scanner.Text())
        x, err := strconv.Atoi(parts[1])
        if err != nil {
            continue
        }
        y, err := strconv.Atoi(parts[2])
        if err != nil {
            continue
        }
        vel_x, err := strconv.Atoi(parts[3])
        if err != nil {
            continue
        }
        vel_y, err := strconv.Atoi(parts[4])
        if err != nil {
            continue
        }

        points = append(points, &Point{x, y, vel_x, vel_y})
    }

    min_size := math.MaxInt32
    for i := 1; true; i++ {
        bounds := tick(points)
        size := bounds.Size()
        if min_size <= size {
            continue
        }

        min_size = bounds.Size()
        if size > 1000 {
            continue
        }
        fmt.Printf("Secs: %d\n", i)
        display(points, bounds)
    }
}

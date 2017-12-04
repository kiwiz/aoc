package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "errors"
    "math"
)

type particle struct {
    x int
    y int
    z int
    vx int
    vy int
    vz int
    ax int
    ay int
    az int
    dead bool
    cnt int
}

func dist(p *particle) float64 {
    return math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z))
}

func update(p *particle) {
    p.vx += p.ax
    p.vy += p.ay
    p.vz += p.az
    p.x += p.vx
    p.y += p.vy
    p.z += p.vz
}

func parse_vector(str string) (int, int, int, error) {
    if len(str) < 2 || str[0] != '<' || str[len(str) - 1] != '>' {
        return 0, 0, 0, errors.New("Invalid vector")
    }

    parts := strings.Split(str[1:len(str) - 1], ",")
    if len(parts) != 3 {
        return 0, 0, 0, errors.New("Invalid vector")
    }

    a, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, 0, 0, err
    }
    b, err := strconv.Atoi(parts[1])
    if err != nil {
        return 0, 0, 0, err
    }
    c, err := strconv.Atoi(parts[2])
    if err != nil {
        return 0, 0, 0, err
    }

    return a, b, c, nil
}

func parse_line(str string) *particle {
    parts := strings.Split(str, ", ")
    if len(parts) != 3 {
        return nil
    }

    if len(parts[0]) < 2 || parts[0][0:2] != "p=" {
        return nil
    }
    if len(parts[1]) < 2 || parts[1][0:2] != "v=" {
        return nil
    }
    if len(parts[2]) < 2 || parts[2][0:2] != "a=" {
        return nil
    }

    x, y, z, err := parse_vector(parts[0][2:])
    if err != nil {
        return nil
    }
    vx, vy, vz, err := parse_vector(parts[1][2:])
    if err != nil {
        return nil
    }
    ax, ay, az, err := parse_vector(parts[2][2:])
    if err != nil {
        return nil
    }

    var p particle
    p.x, p.y, p.z = x, y, z
    p.vx, p.vy, p.vz = vx, vy, vz
    p.ax, p.ay, p.az = ax, ay, az

    return &p
}

func hash(p *particle) int64 {
    return int64(p.x) + 10001 * int64(p.y) + 100000001 * int64(p.z)
}

func part_one(particles []*particle) int {
    for i := 0; i < 10000; i++ {
        closest_i := 0
        closest_dist := math.MaxFloat32
        for n, p := range particles {
            update(p)

            d := dist(p)
            if d < closest_dist {
                closest_dist = d
                closest_i = n
            }
        }

        particles[closest_i].cnt += 1
    }

    max_i := 0
    max_cnt := 0
    for n, p := range particles {
        if p.cnt > max_cnt {
            max_cnt = p.cnt
            max_i = n
        }
    }

    return max_i
}

func part_two(particles []*particle) int {
    for i := 0; i < 10000; i++ {
        positions := make(map[int64][]int)

        for n, p := range particles {
            if p.dead {
                continue
            }

            h := hash(p)
            positions[h] = append(positions[h], n)
            update(p)
        }

        for _, vals := range positions {
            if len(vals) < 2 {
                continue
            }

            for _, j := range vals {
                particles[j].dead = true
            }
        }
    }

    acc := 0
    for _, p := range particles {
        if !p.dead {
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

    var particles []*particle
    for scanner.Scan() {
        p := parse_line(scanner.Text())
        if p == nil {
            fmt.Println("Invalid particle")
            continue
        }

        particles = append(particles, p)
    }

    cnt := part_two(particles)

    fmt.Println("Index", cnt)
}

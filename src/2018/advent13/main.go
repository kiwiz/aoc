package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "sort"
)

const (
    NORTH = '^'
    EAST = '>'
    SOUTH = 'v'
    WEST = '<'
)

const (
    UD = '|'
    LR = '-'
    TLBR = '\\'
    BLTR = '/'
    INTR = '+'
)

type CartList []*Cart
func (c CartList) Len() int {
    return len(c)
}
func (c CartList) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}
func (c CartList) Less(i, j int) bool {
    if c[i].Y < c[j].Y {
        return true
    }
    return c[i].X < c[j].X
}

type Cart struct {
    X, Y int
    Dir rune
    Next int
}

type Tracks struct {
    Data [][]rune
    Carts CartList
}

var tlbr = map[rune]rune{
    NORTH: WEST,
    WEST:  NORTH,
    SOUTH: EAST,
    EAST:  SOUTH,
}

var bltr = map[rune]rune{
    NORTH: EAST,
    EAST:  NORTH,
    SOUTH: WEST,
    WEST:  SOUTH,
}

var left = map[rune]rune{
    NORTH: WEST,
    WEST:  SOUTH,
    SOUTH: EAST,
    EAST:  NORTH,
}

var right = map[rune]rune{
    NORTH: EAST,
    EAST:  SOUTH,
    SOUTH: WEST,
    WEST:  NORTH,
}

type Pair struct{
    X, Y int
}

var delta = map[rune]Pair{
    NORTH: Pair{ 0, -1},
    EAST:  Pair{ 1,  0},
    SOUTH: Pair{ 0,  1},
    WEST:  Pair{-1,  0},
}

func (t *Tracks) Tick() []*Pair {
    sort.Sort(t.Carts)

    check := make(map[Pair]int)
    collisions := make([]*Pair, 0)
    for _, cart := range t.Carts {
        p1 := Pair{cart.X, cart.Y}
        if check[p1] > 0 {
            if check[p1] == 1 {
                collisions = append(collisions, &p1)
            }
            continue
        }

        d := delta[cart.Dir]
        cart.X += d.X
        cart.Y += d.Y

        track := t.Data[cart.Y][cart.X]
        if track == TLBR {
            cart.Dir = tlbr[cart.Dir]
        }
        if track == BLTR {
            cart.Dir = bltr[cart.Dir]
        }
        if track == INTR {
            switch(cart.Next) {
            case 0:
                cart.Dir = left[cart.Dir]
            case 2:
                cart.Dir = right[cart.Dir]
            }
            cart.Next = (cart.Next + 1) % 3
        }

        p2 := Pair{cart.X, cart.Y}
        if check[p2] == 1 {
            collisions = append(collisions, &p2)
        }
        check[p2] += 1
    }

    return collisions
}

func (t *Tracks) Print() {
    disp := make([][]rune, 0)
    for _, row := range t.Data {
        new_row := make([]rune, len(row))
        copy(new_row, row)
        disp = append(disp, new_row)
    }

    for _, cart := range t.Carts {
        val := disp[cart.Y][cart.X]
        if val == NORTH || val == SOUTH || val == EAST || val == WEST || val == 'X' {
            disp[cart.Y][cart.X] = 'X'
        } else {
            disp[cart.Y][cart.X] = cart.Dir
        }
    }

    for _, row := range disp {
        fmt.Println(string(row))
    }
}

func part_one(world *Tracks) {
    var collisions []*Pair
    for {
        collisions = world.Tick()
        if len(collisions) > 0 {
            break
        }
    }

    fmt.Printf("Pos: (%d, %d)\n", collisions[0].X, collisions[0].Y)
}

func part_two(world *Tracks) {
    for len(world.Carts) > 1 {
        var collisions []*Pair
        for {
            collisions = world.Tick()
            if len(collisions) > 0 {
                break
            }
        }

        for _, collision := range collisions {
            for i := len(world.Carts) - 1; i >= 0; i-- {
                if world.Carts[i].X != collision.X || world.Carts[i].Y != collision.Y {
                    continue
                }

                sz := len(world.Carts)
                world.Carts[i] = world.Carts[sz - 1]
                world.Carts = world.Carts[:sz - 1]
            }
        }
    }

    cart := world.Carts[0]
    fmt.Printf("Pos: (%d, %d)\n", cart.X, cart.Y)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    scanner := bufio.NewScanner(file)
    track_data := make([][]rune, 0)
    carts := make(CartList, 0)
    y := 0
    for scanner.Scan() {
        line := []rune(strings.Trim(scanner.Text(), "\n"))
        for x, val := range line {
            if val == NORTH || val == SOUTH {
                line[x] = UD
                carts = append(carts, &Cart{x, y, val, 0})
            }
            if val == WEST || val == EAST {
                line[x] = LR
                carts = append(carts, &Cart{x, y, val, 0})
            }
        }
        track_data = append(track_data, line)
        y++
    }

    world := &Tracks{track_data, carts}

    part_two(world)
}

package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "sort"
    "math"
    "container/heap"
)

const (
    NEW = 0
    OPEN = 1
    DONE = 2
)

const MAX_HP = 200

type UnitList []*Unit
func (u UnitList) Len() int {
    return len(u)
}
func (u UnitList) Swap(i, j int) {
    u[i], u[j] = u[j], u[i]
}
func (u UnitList) Less(i, j int) bool {
    if !u[i].Dead() && u[j].Dead() {
        return true
    }

    return u[i].Cell.Id < u[j].Cell.Id
}

type Unit struct {
    Friendly bool
    HP int
    Cell *Cell
}

func (u *Unit) Dead() bool {
    return u.HP == 0
}

func (u *Unit) Move(c *Cell) {
    if u.Cell != nil {
        u.Cell.Unit = nil
    }
    u.Cell = c
    c.Unit = u
}
func (u *Unit) Attack(o *Unit, dmg int) {
    o.HP -= dmg
    if o.HP < 0 {
        o.HP = 0
    }
}

func NewUnit(friendly bool) *Unit {
    return &Unit{friendly, MAX_HP, nil}
}

type Path struct {
    Next *Cell
    End *Cell
}
type PathList []*Path
func (p PathList) Len() int {
    return len(p)
}
func (p PathList) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p PathList) Less(i, j int) bool {
    if p[i].End != p[j].End {
        return p[i].End.Id < p[j].End.Id
    }

    return p[i].Next.Id < p[j].Next.Id
}

type Cell struct {
    Id int
    Unit *Unit
    Wall bool
    neighbors []*Cell

    from []*Cell
    currDistance int
    index int
}
func (c *Cell) Neighbors() []*Cell {
    return c.neighbors
}
func (c *Cell) Distance(o *Cell) int {
    return 1
}

type PQueue []*Cell
func (p PQueue) Len() int {
    return len(p)
}
func (p PQueue) Less(i, j int) bool {
    return p[i].currDistance < p[j].currDistance
}
func (p PQueue) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
    p[i].index = i
    p[j].index = j
}
func (p *PQueue) Push(x interface{}) {
    n := len(*p)
    item := x.(*Cell)
    item.index = n
    *p = append(*p, item)
}
func (p *PQueue) Pop() interface{} {
    old := *p
    n := len(old)
    item := old[n - 1]
    item.index = -1
    *p = old[0:n - 1]

    return item
}
func (p *PQueue) Update(x *Cell) {
    heap.Fix(p, x.index)
}

func clearCell(c *Cell, dist int) {
    c.currDistance = dist
    c.from = make([]*Cell, 0)
    c.index = 0
}

func ExpandPaths(end *Cell, start *Cell) PathList {
    paths := make(PathList, 0)
    q := make([]*Cell, 0)
    q = append(q, end)

    visited := make(map[int]bool)
    visited[end.Id] = true
    for len(q) > 0 {
        n := len(q)
        cell := q[n - 1]
        q = q[:n - 1]
        visited[cell.Id] = true
        for _, from := range cell.from {
            if from == start {
                paths = append(paths, &Path{cell, end})
                continue
            }
            if from == nil {
                continue
            }
            if !visited[from.Id] {
                q = append(q, from)
            }
        }
    }

    sort.Sort(paths)
    return paths
}

func NextCell(start *Cell, ends []*Cell) *Cell {
    for _, cell := range ends {
        clearCell(cell, 0)
    }

    open_cells := make(PQueue, 0)
    cell_map := make(map[int]int)

    clearCell(start, 0)
    heap.Push(&open_cells, start)
    cell_map[start.Id] = OPEN

    for open_cells.Len() > 0 {
        curr_cell := heap.Pop(&open_cells).(*Cell)
        cell_map[curr_cell.Id] = DONE

        for _, cell := range curr_cell.Neighbors() {
            // Check if cell closed or blocked
            state := cell_map[cell.Id]
            if state == DONE || cell.Unit != nil {
                continue
            }
            dist := curr_cell.currDistance + curr_cell.Distance(cell)
            if state == NEW {
                clearCell(cell, 0)
                cell_map[cell.Id] = OPEN
                cell.currDistance = dist
                cell.from = append(cell.from, curr_cell)
                open_cells.Push(cell)
            }
            if dist < cell.currDistance {
                cell.currDistance = dist
                cell.from = append(cell.from, curr_cell)
                open_cells.Update(cell)
            }
            if dist == cell.currDistance {
                cell.from = append(cell.from, curr_cell)
            }
        }
    }

    var next_cell *Cell = nil
    min_dist := math.MaxInt32
    for _, end := range ends {
        // No path found
        if len(end.from) == 0 {
            continue
        }
        // Path too long
        if end.currDistance > min_dist {
            continue
        }

        var paths []*Path
        if end.currDistance < min_dist {
            paths = ExpandPaths(end, start)
            min_dist = end.currDistance
        } else {
            paths = append(paths, ExpandPaths(end, start)...)
        }

        // Pick best path
        min_end_id := math.MaxInt32
        min_next_id := math.MaxInt32
        for _, path := range paths {
            if path.End.Id > min_end_id {
                continue
            }
            if path.End.Id < min_end_id {
                next_cell = path.Next
                min_end_id = path.End.Id
                min_next_id = math.MaxInt32
                paths = []*Path{path}
                continue
            }

            if path.Next.Id < min_next_id {
                next_cell = path.Next
                min_next_id = path.Next.Id
                paths = []*Path{path}
                continue
            }
        }
    }

    return next_cell
}

type Map struct {
    Data [][]*Cell
    W int
    H int
    elfDamage int
    Units UnitList
}

func NewMap(data [][]*Cell, dmg int) *Map {
    h := len(data)
    w := 0
    if h > 0 {
        w = len(data[0])
    }

    units := make(UnitList, 0)
    new_data := make([][]*Cell, h)
    for y := 0; y < h; y++ {
        new_row := make([]*Cell, w)
        for x := 0; x < w; x++ {
            new_cell := *data[y][x]
            if new_cell.Unit != nil {
                new_unit := *new_cell.Unit
                new_unit.Cell = &new_cell
                new_cell.Unit = &new_unit
                units = append(units, &new_unit)
            }
            new_row[x] = &new_cell
        }
        new_data[y] = new_row
    }

    deltas := [][]int{
        { 0, -1},
        {-1,  0},
        { 1,  0},
        { 0,  1},
    }
    i := 0
    for y := 0; y < h; y++ {
        for x := 0; x < w; x++ {
            new_data[y][x].Id = i
            new_data[y][x].neighbors = make([]*Cell, 0)
            i++

            if new_data[y][x].Wall {
                continue
            }

            for _, delta := range deltas {
                nx, ny := x + delta[0], y + delta[1]
                if nx < 0 || ny < 0 || nx >= w || ny >= h {
                    continue
                }

                neighbor := new_data[ny][nx]
                if neighbor.Wall {
                    continue
                }

                new_data[y][x].neighbors = append(new_data[y][x].neighbors, neighbor)
            }
        }
    }


    return &Map{new_data, w, h, dmg, units}
}

func (m *Map) Tick() bool {
    sort.Sort(m.Units)

    for _, unit := range m.Units {
        if unit.Dead() {
            continue
        }

        target := m.PickTarget(unit)
        if target == nil {
            enemies := m.FindEnemies(unit.Friendly)
            if len(enemies) == 0 {
                return false
            }
            target_cells := m.FindOpenCells(enemies)
            cell := NextCell(unit.Cell, target_cells)
            if cell == nil {
                continue
            }
            unit.Move(cell)
            target = m.PickTarget(unit)
        }

        if target == nil {
            continue
        }

        if unit.Friendly {
            unit.Attack(target, m.elfDamage)
        } else {
            unit.Attack(target, 3)
        }
        if target.Dead() {
            m.Remove(target)
        }
    }
    return true
}

func (m *Map) FindEnemies(friendly bool) []*Unit {
    enemies := make([]*Unit, 0)
    for _, unit := range m.Units {
        if unit.Friendly != friendly && !unit.Dead() {
            enemies = append(enemies, unit)
        }
    }

    return enemies
}

func (m *Map) Remove(u *Unit) {
    u.Cell.Unit = nil
}

func (m *Map) FindOpenCells(units []*Unit) []*Cell {
    open := make([]*Cell, 0)
    for _, unit := range units {
        for _, neighbor := range unit.Cell.Neighbors() {
            if neighbor.Unit != nil {
                continue
            }

            open = append(open, neighbor)
        }
    }

    return open
}

func (m *Map) PickTarget(u *Unit) *Unit {
    var selected_target *Unit
    min_hp := MAX_HP
    id := math.MaxInt32

    for _, neighbor := range u.Cell.Neighbors() {
        target := neighbor.Unit
        if target == nil || u.Friendly == target.Friendly {
            continue
        }

        if !(selected_target == nil || target.HP < min_hp || (target.HP == min_hp && target.Cell.Id < id)) {
            continue
        }

        selected_target = target
        min_hp = target.HP
        id = target.Cell.Id
    }

    return selected_target
}

func (m *Map) Print() {
    for _, row := range m.Data {
        var sb strings.Builder
        for _, cell := range row {
            val := 'X'
            if cell.Wall {
                val = '#'
            } else if cell.Unit == nil {
                val = '.'
            } else if cell.Unit.Friendly {
                val = 'E'
            } else {
                val = 'G'
            }
            sb.WriteRune(val)
        }
        fmt.Println(sb.String())
    }
}

func part_one(data [][]*Cell) {
    m := NewMap(data, 3)
    i := 0
    m.Print()
    for ; m.Tick(); i++ {
        fmt.Println("TURN")
        m.Print()
    }
    m.Print()
    total_hp := 0
    for _, unit := range m.Units {
        total_hp += unit.HP
    }
    fmt.Printf("Score: %d\n", total_hp * i)
}

func part_two(data [][]*Cell) {
    i := 3
    var m *Map
    for ; true; i++ {
        m = NewMap(data, i)
        for m.Tick() {}
        dead := false
        for _, unit := range m.Units {
            if unit.Friendly && unit.HP == 0 {
                dead = true
                break
            }
        }
        if !dead {
            break
        }
    }

    m.Print()
    total_hp := 0
    for _, unit := range m.Units {
        total_hp += unit.HP
    }
    fmt.Printf("Score: %d\n", total_hp * i)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    // FIXME: This solution is broken.

    scanner := bufio.NewScanner(file)
    data := make([][]*Cell, 0)
    units := make(UnitList, 0)
    for scanner.Scan() {
        cells := make([]*Cell, 0)
        for _, val := range strings.Trim(scanner.Text(), "\n") {
            cell := &Cell{}
            switch(val) {
            case 'E':
                unit := NewUnit(true)
                unit.Move(cell)
                units = append(units, unit)
            case 'G':
                unit := NewUnit(false)
                unit.Move(cell)
                units = append(units, unit)
            case '.':
            case '#':
                fallthrough
            default:
                cell.Wall = true
            }
            cells = append(cells, cell)
        }
        data = append(data, cells)
    }

    part_one(data)
    part_two(data)
}

package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

type program struct {
    weight int
    children []*program
    parent *program
    key string
}

func (prog *program) TotalWeight() int {
    acc := 0
    acc += prog.weight
    for _, child := range prog.children {
        acc += child.TotalWeight()
    }

    return acc
}

func (prog *program) IsBalanced() bool {
    weight := 0
    for i, child := range prog.children {
        child_weight := child.TotalWeight()
        if i == 0 {
            weight = child_weight
        }

        if weight != child_weight {
            return false
        }
    }

    return true
}

type prog_map struct {
    prog *program
    children_keys []string
}

func build_tree(programs map[string]prog_map) *program {
    var prog *program
    for _, pm := range programs {
        prog = pm.prog
        for _, child_key := range pm.children_keys {
            child := programs[child_key].prog
            prog.children = append(prog.children, child)
            child.parent = prog
        }
    }

    for prog.parent != nil {
        prog = prog.parent
    }

    return prog
}

func parse_line(text string) (*program, []string) {
    var prog program
    var children_keys []string
    parts := strings.Split(text, " -> ")

    if len(parts) > 1 {
        children_keys = strings.Split(parts[1], ", ")
    }

    parts = strings.Split(parts[0], " ")
    if len(parts) != 2 {
        fmt.Println("Invalid line")
        return nil, nil
    }

    prog.key = parts[0]
    weight, err := strconv.Atoi(strings.Trim(parts[1], "()"))
    if err != nil {
        fmt.Println("Invalid line")
        return nil, nil
    }
    prog.weight = weight

    return &prog, children_keys
}

func mode(nums []int) int {
    var counts map[int]int = make(map[int]int)
    for _, i := range nums {
        _, ok := counts[i]
        if !ok {
            counts[i] = 0
        }
        counts[i] += 1
    }

    max := 0
    num := 0
    for i, cnt := range counts {
        if cnt > max {
            max = cnt
            num = i
        }
    }

    return num
}

func find_invalid_node(node *program) *program {
    for _, child := range node.children {
        if !child.IsBalanced() {
            return find_invalid_node(child)
        }
    }

    return node
}

func main() {
    var programs map[string]prog_map = make(map[string]prog_map)

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        prog, children_keys := parse_line(scanner.Text())
        programs[prog.key] = prog_map{prog, children_keys}
    }

    node := build_tree(programs)

    invalid_node := find_invalid_node(node)

    var weights []int
    for _, child := range invalid_node.children {
        weights = append(weights, child.TotalWeight())
    }
    expected_weight := mode(weights)

    for _, child := range invalid_node.children {
        weight := child.TotalWeight()
        if weight != expected_weight {
            delta := expected_weight - weight
            fmt.Println(child.weight + delta)
            break
        }
    }
}

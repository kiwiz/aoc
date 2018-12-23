package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Node struct {
    children []*Node
    values []int
}

func NewNode() *Node {
    return &Node{}
}

func build(stream []int, i int) (*Node, int) {
    node := NewNode()
    n_children := stream[i]
    i++
    n_values := stream[i]
    i++
    for c := 0; c < n_children; c++ {
        var child *Node
        child, i = build(stream, i)
        node.children = append(node.children, child)
    }
    for c := 0; c < n_values; c++ {
        node.values = append(node.values, stream[i])
        i++
    }

    return node, i
}

func sum(node *Node) int {
    c := 0
    for i := 0; i < len(node.values); i++ {
        c += node.values[i]
    }
    for _, child := range node.children {
        c += sum(child)
    }
    return c
}

func resolve(node *Node) int {
    c := 0
    if len(node.children) == 0 {
        for i := 0; i < len(node.values); i++ {
            c += node.values[i]
        }
    } else {
        for _, i := range node.values {
            if i == 0 || i > len(node.children) {
                continue
            }
            c += resolve(node.children[i - 1])
        }
    }
    return c
}

func part_one(root *Node) {
    fmt.Printf("Sum: %d\n", sum(root))
}

func part_two(root *Node) {
    fmt.Printf("Sum: %d\n", resolve(root))
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    reader := bufio.NewReader(file)
    stream := make([]int, 0)
    for {
        chunk, err := reader.ReadString(' ')
        if chunk == "" {
            break
        }
        i, err := strconv.Atoi(strings.Trim(chunk, "\n "))
        if err != nil {
            fmt.Println(err)
            continue
        }
        stream = append(stream, i)
    }

    root, _ := build(stream, 0)
    part_one(root)
    part_two(root)
}

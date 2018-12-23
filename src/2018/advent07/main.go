package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "sort"
)

type Node struct {
    val rune
    prev []*Node
    next []*Node
    finished bool
    work int
}

func (n *Node) Ready() bool {
    for _, node := range n.prev {
        if !node.finished {
            return false
        }
    }
    return true
}

type RuneSlice []rune
func (s RuneSlice) Len() int {
    return len(s)
}
func (s RuneSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s RuneSlice) Less(i, j int) bool {
    return s[i] < s[j]
}

func NewNode(val rune) *Node {
    return &Node{
        val,
        make([]*Node, 0),
        make([]*Node, 0),
        false,
        int(rune(60) + val - 'A' + 1),
    }
}

func part_one(nodes map[rune]*Node) {
    traverse(nodes, 1)
}

func part_two(nodes map[rune]*Node) {
    traverse(nodes, 5)
}

func traverse(nodes map[rune]*Node, workers int) {
    secs := 0
    ordering := make([]rune, 0)

    for len(nodes) > 0 {
        ready_vals := make(RuneSlice, 0)
        min_work := 0xff
        for _, node := range nodes {
            if !node.Ready() {
                continue
            }

            ready_vals = append(ready_vals, node.val)
            if min_work > node.work {
                min_work = node.work
            }
        }

        sort.Sort(ready_vals)
        secs += min_work

        for i := 0; i < workers && i < len(ready_vals); i++ {
            val := ready_vals[i]
            nodes[val].work -= min_work
            if nodes[val].work == 0 {
                ordering = append(ordering, val)
                nodes[val].finished = true
                delete(nodes, val)
            }
        }
    }

    fmt.Println(string(ordering))
    fmt.Printf("Time: %d\n", secs)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    line_re := regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin.`)

    nodes := make(map[rune]*Node)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        matches := line_re.FindStringSubmatch(scanner.Text())
        if matches == nil {
            continue
        }

        from := rune(matches[1][0])
        from_node, ok := nodes[from]
        if !ok {
            from_node = NewNode(from)
            nodes[from] = from_node
        }

        to := rune(matches[2][0])
        to_node, ok := nodes[to]
        if !ok {
            to_node = NewNode(to)
            nodes[to] = to_node
        }

        to_node.prev = append(to_node.prev, from_node)
        from_node.next = append(from_node.next, to_node)
    }

    part_two(nodes)
}

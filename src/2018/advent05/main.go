package main

import (
    "fmt"
    "os"
    "bufio"
    "unicode"
    "strings"
    "container/list"
)

func react(input *list.List) *list.List {
    ok := true
    for ok {
        ok = false
        e := input.Front()
        for e != nil {
            n := e.Next()
            if n == nil {
                break
            }

            if unicode.ToLower(e.Value.(rune)) == unicode.ToLower(n.Value.(rune)) && e.Value != n.Value {
                t := e.Prev()
                if t == nil {
                    t = n.Next()
                }
                input.Remove(n)
                input.Remove(e)
                e = t
                ok = true
            } else {
                e = n
            }
        }
    }

    return input
}

func part_one(input *list.List) {
    output := react(copy_list(input))
    fmt.Printf("Len: %d\n", output.Len())
}

func copy_list(input *list.List) *list.List {
    list := list.New()
    for e := input.Front(); e != nil; e = e.Next() {
        list.PushBack(e.Value)
    }

    return list
}

func part_two(input *list.List) {
    min_size := input.Len()
    for i := 'a'; i <= 'z'; i++ {
        j := unicode.ToUpper(i)
        local := copy_list(input)
        e := local.Front();
        for e != nil {
            t := e.Next()
            if e.Value == i || e.Value == j {
                local.Remove(e)
            }
            e = t
        }
        size := react(local).Len()
        if size < min_size {
            min_size = size
        }
    }

    fmt.Printf("Len: %d\n", min_size)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    reader := bufio.NewReader(file)
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }

    input = strings.Trim(input, "\n ")
    list := list.New()
    for _, r := range input {
        list.PushBack(r)
    }

    part_one(list)
    part_two(list)
}

package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
)

type node struct {
    val int
    next *node
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()

    skip, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Invalid input")
        return
    }

    curr := &node{0, nil}
    curr.next = curr

    for n := 1; n <= 50000000; n++ {
        for i := 0; i < skip; i++ {
            curr = curr.next
        }

        t := curr.next
        curr.next = &node{n, t}
        curr = curr.next

        if n % 100000 == 0 {
            fmt.Println(n)
        }
    }

    for curr.val != 0 {
        curr = curr.next
    }

    fmt.Println("Val", curr.next.val)
}

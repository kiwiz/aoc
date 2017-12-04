package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    reader := bufio.NewReader(file)

    var val byte
    depth := 0
    acc := 0
    junk := 0

    garbage := false
    cancel := false
    for ; err == nil; val, err = reader.ReadByte() {
        if cancel {
            cancel = false
            continue
        }

        switch val {
        case '{':
            if !garbage {
                depth += 1
            } else {
                junk += 1
            }
        case '}':
            if !garbage {
                acc += depth
                depth -= 1
            } else {
                junk += 1
            }
        case '<':
            if !garbage {
                garbage = true
            } else {
                junk += 1
            }
        case '>':
            if garbage {
                garbage = false
            }
        case '!':
            if !cancel {
                cancel = true
            }
        default:
            if garbage {
                junk += 1
            }
        }
    }

    fmt.Println("Score", acc, junk)
}

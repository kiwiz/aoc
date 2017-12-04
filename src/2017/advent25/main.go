package main

import (
    "os"
    "io"
    "fmt"
    "bufio"
    "regexp"
    "strconv"
)

type operation struct {
    write int
    move int
    dest rune
}

type turing struct {
    tape map[int]int
    cursor int
    state rune
    states map[rune]map[int]operation
    steps int
    terminate int
}

func NewTuring() *turing {
    var t turing
    t.tape = make(map[int]int)
    t.states = make(map[rune]map[int]operation)

    return &t
}

func generate(file io.Reader) *turing {
    scanner := bufio.NewScanner(file)

    begin_re := regexp.MustCompile(`Begin in state ([A-Z])\.`)
    diag_re := regexp.MustCompile(`Perform a diagnostic checksum after (\d+) steps\.`)

    state_re := regexp.MustCompile(`In state ([A-Z]):`)
    state_if_re := regexp.MustCompile(`  If the current value is (\d+):`)
    state_wr_re := regexp.MustCompile(`    - Write the value (\d+)\.`)
    state_mv_re := regexp.MustCompile(`    - Move one slot to the (left|right)\.`)
    state_ch_re := regexp.MustCompile(`    - Continue with state ([A-Z])\.`)

    t := NewTuring()

    scanner.Scan()
    match := begin_re.FindStringSubmatch(scanner.Text())
    if len(match) < 2 {
        return nil
    }
    t.state = rune(match[1][0])

    scanner.Scan()
    match = diag_re.FindStringSubmatch(scanner.Text())
    if len(match) < 2 {
        return nil
    }
    t.terminate, _ = strconv.Atoi(match[1])

    scanner.Scan()

    for scanner.Scan() {

        match = state_re.FindStringSubmatch(scanner.Text())
        if len(match) < 2 {
            break
        }
        state := rune(match[1][0])

        ops := make(map[int]operation)

        for scanner.Scan() {
            match = state_if_re.FindStringSubmatch(scanner.Text())
            if len(match) < 2 {
                break
            }
            val, _ := strconv.Atoi(match[1])

            scanner.Scan()
            match = state_wr_re.FindStringSubmatch(scanner.Text())
            if len(match) < 2 {
                return nil
            }
            write, _ := strconv.Atoi(match[1])

            scanner.Scan()
            match = state_mv_re.FindStringSubmatch(scanner.Text())
            if len(match) < 2 {
                return nil
            }
            move := -1
            if match[1] == "right" {
                move = 1
            }

            scanner.Scan()
            match = state_ch_re.FindStringSubmatch(scanner.Text())
            if len(match) < 2 {
                return nil
            }
            dest := rune(match[1][0])
            ops[val] = operation{write, move, dest}
        }

        t.states[state] = ops
    }

    return t
}

func (t *turing) Tick() {
    val := t.tape[t.cursor]
    op := t.states[t.state][val]

    t.tape[t.cursor] = op.write
    t.cursor += op.move
    t.state = op.dest

    t.steps += 1
}

func (t *turing) Checksum() int {
    acc := 0
    for _, val := range t.tape {
        if val == 1 {
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

    t := generate(file)

    for t.terminate > t.steps {
        t.Tick()
    }

    fmt.Println("Checksum", t.Checksum())
}

package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type state struct {
    pc int
    regs map[string]int64
    outq chan int64
    blocked bool
    out_cnt int
    other *state
}

func NewState(p int64) *state {
    var st state
    st.regs = make(map[string]int64)
    st.outq = make(chan int64, 1000)
    st.regs["p"] = p
    return &st
}

func is_reg(str string) bool {
    if len(str) > 1 {
        return false
    }

    return str[0] >= 'a' && str[0] <= 'z'
}

func resolve(st *state, str string) int64 {
    if is_reg(str) {
        return st.regs[str]
    } else {
        n, err := strconv.Atoi(str)
        if err != nil {
            fmt.Println(err)
        }
        return int64(n)
    }
}

type op_info struct {
    fn func(st *state, args []string) (int, bool)
    cnt int
}
var ops map[string]op_info = map[string]op_info{
    "snd": {op_snd, 1},
    "set": {op_set, 2},
    "add": {op_add, 2},
    "mul": {op_mul, 2},
    "mod": {op_mod, 2},
    "rcv": {op_rcv, 1},
    "jgz": {op_jgz, 2},
}

func op_snd(st *state, args []string) (int, bool) {
    st.outq <- resolve(st, args[0])
    st.out_cnt += 1

    return 1, true
}

func op_set(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] = resolve(st, args[1])

    return 1, true
}

func op_add(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] += resolve(st, args[1])

    return 1, true
}

func op_mul(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] *= resolve(st, args[1])

    return 1, true
}

func op_mod(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] %= resolve(st, args[1])

    return 1, true
}

func op_rcv(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }

    st.blocked = true

    if len(st.other.outq) == 0 {
        if st.other.blocked {
            return 0, false
        }
        return 0, true
    }

    st.blocked = false
    st.regs[args[0]] = <-st.other.outq

    return 1, true
}

func op_jgz(st *state, args []string) (int, bool) {
    if resolve(st, args[0]) > 0 {
        return int(resolve(st, args[1])), true
    }

    return 1, true
}

func tick(st *state, instructions [][]string) bool {
    if st.pc < 0 || st.pc >= len(instructions) {
        return false
    }
    instruction := instructions[st.pc]

    op, ok := ops[instruction[0]]
    if !ok {
        fmt.Println("Invalid instruction")
        return false
    }

    delta, ok := op.fn(st, instruction[1:])
    if !ok {
        return false
    }

    st.pc += delta
    return true
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    var instructions [][]string

    for scanner.Scan() {
        parts := strings.Fields(scanner.Text())
        instructions = append(instructions, parts)
    }

    sta := NewState(0)
    stb := NewState(1)
    sta.other = stb
    stb.other = sta

    a := true
    b := true
    for (a || b) {
        if a {
            a = tick(sta, instructions)
        }
        if b {
            b = tick(stb, instructions)
        }
    }

    fmt.Println("Count", stb.out_cnt, stb.out_cnt)
}

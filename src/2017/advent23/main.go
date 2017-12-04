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
    mul_cnt int
}

func NewState() *state {
    var st state
    st.regs = make(map[string]int64)
    st.outq = make(chan int64, 1000)
    st.regs["a"] = 1
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
    "set": {op_set, 2},
    "sub": {op_sub, 2},
    "mul": {op_mul, 2},
    "div": {op_div, 2},
    "jnz": {op_jnz, 2},
    "mod": {op_mod, 2},
    "gcd": {op_gcd, 3},
    "pnc": {op_pnc, 0},
    "prt": {op_prt, 0},
    "nop": {op_nop, 0},
}

func op_xxa(st *state, args []string) (int, bool) {
    return 1, true
}

func op_gcd(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }

    a := resolve(st, args[1])
    b := resolve(st, args[2])

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
    st.regs[args[0]] = a

    return 1, true
}

func op_mod(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] %= resolve(st, args[1])

    return 1, true
}

func op_set(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] = resolve(st, args[1])

    return 1, true
}

func op_sub(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] -= resolve(st, args[1])

    return 1, true
}

func op_mul(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.mul_cnt += 1
    st.regs[args[0]] *= resolve(st, args[1])

    return 1, true
}

func op_div(st *state, args []string) (int, bool) {
    if !is_reg(args[0]) {
        return 0, false
    }
    st.regs[args[0]] /= resolve(st, args[1])

    return 1, true
}

func op_jnz(st *state, args []string) (int, bool) {
    if resolve(st, args[0]) != 0 {
        return int(resolve(st, args[1])), true
    }

    return 1, true
}

func op_prt(st *state, args []string) (int, bool) {
    fmt.Println(st.regs)

    return 1, true
}

func op_pnc(st *state, args []string) (int, bool) {
    panic("")

    return 1, true
}

func op_nop(st *state, args []string) (int, bool) {
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

    st := NewState()

    for tick(st, instructions) {}

    fmt.Println("Val", st.regs["h"])
}

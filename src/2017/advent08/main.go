package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

const (
    IOP_NULL = iota
    IOP_INC = iota
    IOP_DEC = iota
)

type instruction struct {
    reg string
    op int
    imm int
    cond *condition
}

const (
    CCMP_NULL = iota
    CCMP_LT = iota
    CCMP_GT = iota
    CCMP_GTE = iota
    CCMP_LTE = iota
    CCMP_EQ = iota
    CCMP_NE = iota
)

type condition struct {
    reg string
    cmp int
    imm int
}

type value struct {
    curr int
    max int
}

func parse_condition(str string) *condition {
    parts := strings.Fields(str)
    if len(parts) != 3 {
        fmt.Println("Invalid condition")
        return &condition{"", CCMP_NULL, 0}
    }

    imm, err := strconv.Atoi(parts[2])
    if err != nil {
        fmt.Println("Invalid number")
    }
    cmp := CCMP_NULL
    switch parts[1] {
    case ">":
        cmp = CCMP_GT
    case "<":
        cmp = CCMP_LT
    case ">=":
        cmp = CCMP_GTE
    case "<=":
        cmp = CCMP_LTE
    case "==":
        cmp = CCMP_EQ
    case "!=":
        cmp = CCMP_NE
    }

    return &condition{parts[0], cmp, imm}
}

func parse_op(str string) (string, int, int) {
    parts := strings.Fields(str)
    if len(parts) != 3 {
        fmt.Println("Invalid operation")
        return "", IOP_NULL, 0
    }

    imm, err := strconv.Atoi(parts[2])
    if err != nil {
        fmt.Println("Invalid number")
    }

    op := IOP_NULL
    switch parts[1] {
    case "inc":
        op = IOP_INC
    case "dec":
        op = IOP_DEC
    }

    return parts[0], op, imm
}

func parse_instruction(str string) *instruction {
    parts := strings.Split(str, " if ")
    if len(parts) != 2 {
        fmt.Println("Invalid instruction")
        return nil
    }

    reg, op, imm := parse_op(parts[0])
    cond := parse_condition(parts[1])
    return &instruction{reg, op, imm, cond}
}

func eval_condition(registers map[string]*value, cond *condition) bool {
    _, ok := registers[cond.reg]
    if !ok {
        registers[cond.reg] = &value{}
    }

    val := registers[cond.reg].curr
    switch cond.cmp {
    case CCMP_GT:
        return val > cond.imm
    case CCMP_LT:
        return val < cond.imm
    case CCMP_GTE:
        return val >= cond.imm
    case CCMP_LTE:
        return val <= cond.imm
    case CCMP_EQ:
        return val == cond.imm
    case CCMP_NE:
        return val != cond.imm
    }

    return false
}

func eval_instruction(registers map[string]*value, instr *instruction) {
    _, ok := registers[instr.reg]
    if !ok {
        registers[instr.reg] = &value{}
    }

    if !eval_condition(registers, instr.cond) {
        return
    }
    switch instr.op {
    case IOP_INC:
        registers[instr.reg].curr += instr.imm
    case IOP_DEC:
        registers[instr.reg].curr -= instr.imm
    }

    if registers[instr.reg].max < registers[instr.reg].curr {
        registers[instr.reg].max = registers[instr.reg].curr
    }
}

func process(registers map[string]*value, instructions []*instruction) {
    for _, instr := range instructions {
        eval_instruction(registers, instr)
    }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    registers := map[string]*value{}
    var instructions []*instruction
    for scanner.Scan() {
        instructions = append(instructions, parse_instruction(scanner.Text()))
    }

    process(registers, instructions)

    max := 0
    global_max := 0
    for _, val := range registers {
        if val.curr > max {
            max = val.curr
        }
        if val.max > global_max {
            global_max = val.max
        }
    }

    fmt.Println("Max", max, global_max)
}

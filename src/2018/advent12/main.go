package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
)

const INIT_GENERATIONS = 20
const CHUNK_GENERATIONS = 500
const TOTAL_GENERATIONS = 50000000000

type Seq struct {
    data []bool
    Padding int
}

func NewSeq(input string, padding int) *Seq {
    seq := &Seq{
        make([]bool, len(input) + padding * 2),
        padding,
    }

    for i := 0; i < len(input); i++ {
        seq.Set(i, input[i] == '#')
    }

    return seq
}

func (s *Seq) CopyEmpty() *Seq {
    seq := &Seq{
        make([]bool, len(s.data)),
        s.Padding,
    }

    return seq
}

func (s *Seq) String() string {
    data := make([]rune, len(s.data))
    for i, val := range s.data {
        if val {
            data[i] = '#'
        } else {
            data[i] = ' '
        }
    }

    return string(data)
}

func (s *Seq) Range(delta int) (int, int) {
    return -s.Padding + delta, len(s.data) - s.Padding - delta
}

func (s *Seq) Set(i int, val bool) {
    s.data[i + s.Padding] = val
}

func (s *Seq) Get(i int) bool {
    return s.data[i + s.Padding]
}

type Rule struct {
    Pattern *Seq
    Result bool
}

func (r *Rule) Match(garden *Seq, i int) bool {
    return (
        garden.Get(i - 2) == r.Pattern.Get(0) &&
        garden.Get(i - 1) == r.Pattern.Get(1) &&
        garden.Get(i + 0) == r.Pattern.Get(2) &&
        garden.Get(i + 1) == r.Pattern.Get(3) &&
        garden.Get(i + 2) == r.Pattern.Get(4))
}

func score(garden *Seq) int {
    sum := 0
    lo, hi := garden.Range(2)
    for x := lo; x < hi; x++ {
        if garden.Get(x) {
            sum += x
        }
    }

    return sum
}

func tick(garden *Seq, rules []*Rule) *Seq {
    next := garden.CopyEmpty()
    lo, hi := garden.Range(2)
    for x := lo; x < hi; x++ {
        for _, rule := range rules {
            if rule.Match(garden, x) {
                next.Set(x, rule.Result)
                break
            }
        }
    }

    return next
}

func part_one(input string, rules []*Rule) {
    garden := NewSeq(input, INIT_GENERATIONS * len(rules))

    for i := 0; i < INIT_GENERATIONS; i++ {
        garden = tick(garden, rules)
    }

    fmt.Printf("Score: %d\n", score(garden))
}

func part_two(input string, rules []*Rule) {
    garden := NewSeq(input, CHUNK_GENERATIONS * len(rules))

    for i := 0; i < CHUNK_GENERATIONS; i++ {
        garden = tick(garden, rules)
    }
    a_score := score(garden)

    for i := 0; i < CHUNK_GENERATIONS; i++ {
        garden = tick(garden, rules)
    }
    b_score := score(garden)

    chunk_score := b_score - a_score
    num_chunks := (TOTAL_GENERATIONS - (CHUNK_GENERATIONS * 2)) / CHUNK_GENERATIONS
    total_score := b_score + chunk_score * num_chunks

    fmt.Printf("Score: %d\n", total_score)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    state_re := regexp.MustCompile(`initial state: (.+)`)
    rule_re := regexp.MustCompile(`(.+) => (.)`)

    scanner := bufio.NewScanner(file)
    if !scanner.Scan() {
        fmt.Println("Invalid format")
        return
    }

    state_matches := state_re.FindStringSubmatch(scanner.Text())
    if state_matches == nil {
        fmt.Println("Initial state not found")
        return
    }

    scanner.Scan()

    rules := make([]*Rule, 0)
    for scanner.Scan() {
        rule_matches := rule_re.FindStringSubmatch(scanner.Text())
        if rule_matches == nil {
            continue
        }

        if rule_matches[2][0] != '#' {
            continue
        }

        rules = append(rules, &Rule{NewSeq(rule_matches[1], 0), true})
    }

    part_one(state_matches[1], rules)
    part_two(state_matches[1], rules)
}

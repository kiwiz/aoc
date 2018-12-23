package main

import (
    "fmt"
    "bufio"
    "regexp"
    "os"
    "time"
    "sort"
    "strconv"
    "strings"
)

type Action int
const (
    Begin Action = iota
    Sleep
    Wake
)

type Entry struct {
    date time.Time
    action Action
    id int
}

type TimeSheet struct {
    minsAsleep int
    record [60]int
}

type EntryList []*Entry
func (r EntryList) Len() int {
    return len(r)
}
func (r EntryList) Less(i, j int) bool {
    return r[i].date.Before(r[j].date)
}
func (r EntryList) Swap(i, j int) {
    r[i], r[j] = r[j], r[i]
}

func parse_line(parts []string) *Entry {
    year, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil
    }
    month, err := strconv.Atoi(parts[2])
    if err != nil {
        return nil
    }
    day, err := strconv.Atoi(parts[3])
    if err != nil {
        return nil
    }
    hour, err := strconv.Atoi(parts[4])
    if err != nil {
        return nil
    }
    min, err := strconv.Atoi(parts[5])
    if err != nil {
        return nil
    }

    id := 0
    action := Begin
    if len(parts[7]) > 0 {
        id, err = strconv.Atoi(parts[7])
        if err != nil {
            return nil
        }
    } else {
        if strings.HasPrefix(parts[6], "falls") {
            action = Sleep
        } else {
            action = Wake
        }
    }
    date := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.UTC)
    return &Entry{date, action, id}
}

func seek(sheet *TimeSheet, start int, end int) {
    for i := start; i < end; i++ {
        sheet.record[i] += 1
        sheet.minsAsleep += 1
    }
}

func part_one(sheets map[int]*TimeSheet) {
    guard_id := 0
    max_mins := 0
    for i, sheet := range sheets {
        if max_mins < sheet.minsAsleep {
            max_mins = sheet.minsAsleep
            guard_id = i
        }
    }

    min := 0
    max_times := 0
    for i, times := range sheets[guard_id].record {
        if max_times < times {
            max_times = times
            min = i
        }
    }

    fmt.Printf("Val: %d\n", guard_id * min)
}

func part_two(sheets map[int]*TimeSheet) {
    guard_id := 0
    min := 0
    max_mins := 0
    for i, sheet := range sheets {
        for j, times := range sheet.record {
            if max_mins < times {
                max_mins = times
                guard_id = i
                min = j
            }
        }
    }

    fmt.Printf("Val: %d\n", guard_id * min)
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file", err)
        return
    }

    line_re := regexp.MustCompile(`\[(\d+)\-(\d+)\-(\d+) (\d+):(\d+)\] (Guard #(\d+)|falls|wakes)`)

    records := make(EntryList, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        parts := line_re.FindStringSubmatch(scanner.Text())

        record := parse_line(parts)
        if record == nil {
            continue
        }
        records = append(records, record)
    }
    sort.Sort(records)

    sheets := make(map[int]*TimeSheet, 0)
    var curr_sheet *TimeSheet
    min := 0
    awake := true
    for _, record := range records {
        if record.action == Begin {
            if curr_sheet != nil && !awake {
                seek(curr_sheet, min, 60)
            }

            _, ok := sheets[record.id]
            if !ok {
                sheets[record.id] = &TimeSheet{}
            }
            curr_sheet = sheets[record.id]
            min = 0
            awake = true
        }
        if record.action == Sleep {
            min = record.date.Minute()
            awake = false
        }
        if record.action == Wake {
            seek(curr_sheet, min, record.date.Minute())
            min = record.date.Minute()
            awake = true
        }
    }
    seek(curr_sheet, min, 60)

    part_one(sheets)
    part_two(sheets)
}

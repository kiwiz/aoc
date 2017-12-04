package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

type grid [][]bool

func parse_line(str string) (grid, grid) {
    parts := strings.Split(str, " => ")
    if len(parts) != 2 {
        return nil, nil
    }

    a := parse_chunk(strings.Split(parts[0], "/"))
    b := parse_chunk(strings.Split(parts[1], "/"))

    return a, b
}

func parse_chunk(rows []string) grid {
    var chunk grid

    for _, str := range rows {
        var row []bool
        for _, pix := range str {
            row = append(row, pix == '#')
        }
        chunk = append(chunk, row)
    }

    return chunk
}

func hash(chunk grid) int {
    i := 0
    if len(chunk) == 3 {
        i = 1
    }
    for _, row := range chunk {
        for _, b := range row {
            i <<= 1
            if b {
                i |= 1
            }
        }
    }

    return i
}

func alloc(size int) grid {
    var chunk grid

    for j := 0; j < size; j++ {
        var row []bool
        for i := 0; i < size; i++ {
            row = append(row, false)
        }

        chunk = append(chunk, row)
    }

    return chunk
}

func rotate(chunk grid) grid {
    new_chunk := alloc(len(chunk))

    for i := 0; i < len(chunk); i++ {
        for j := 0; j < len(chunk); j++ {
            new_chunk[i][j] = chunk[len(chunk) - j - 1][i];
        }
    }

    return new_chunk
}

func flip(chunk grid) grid {
    new_chunk := alloc(len(chunk))

    mid := (len(chunk) + 1) / 2
    for i := 0; i < mid; i++ {
        j := len(chunk) - i - 1
        new_chunk[i] = chunk[j]
        new_chunk[j] = chunk[i]
    }

    return new_chunk
}

func permute(chunk grid) []grid {
    var chunks []grid

    for i := 0; i < 4; i++ {
        b := flip(chunk)
        chunks = append(chunks, chunk, b)
        chunk = rotate(chunk)
    }

    return chunks
}

func extract(img grid, x, y, block_size int) grid {
    var new_chunk grid

    for i := 0; i < block_size; i++ {
        new_chunk = append(new_chunk, img[y * block_size + i][x * block_size:(x + 1) * block_size])
    }

    return new_chunk
}

func write(img grid, x, y, block_size int, chunk grid) {
    for j := 0; j < block_size; j++ {
        for i := 0; i < block_size; i++ {
            img[y * block_size + j][x * block_size + i]= chunk[j][i]
        }
    }
}

func print_grid(img grid) {
    for j := 0; j < len(img); j++ {
        for i := 0; i < len(img[j]); i++ {
            b := "."
            if img[j][i] {
                b = "#"
            }
            fmt.Print(b)
        }
        fmt.Println("")
    }
    fmt.Println("-----")
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Unable to open file")
        return
    }

    scanner := bufio.NewScanner(file)

    img := grid{
        {false,  true, false},
        {false, false,  true},
        { true,  true,  true},
    }

    rules := make(map[int]grid)

    for scanner.Scan() {
        in, out := parse_line(scanner.Text())

        for _, variant := range permute(in) {
            rules[hash(variant)] = out
        }
    }

    for i := 0; i < 18; i++ {
        block_size := 3
        if len(img) % 2 == 0 {
            block_size = 2
        }
        new_block_size := block_size + 1

        new_img := alloc((len(img) / block_size) * new_block_size)
        for y := 0; y < len(img) / block_size; y += 1 {
            for x := 0; x < len(img[y]) / block_size; x += 1 {
                orig_chunk := extract(img, x, y, block_size)
                new_chunk := rules[hash(orig_chunk)]
                write(new_img, x, y, new_block_size, new_chunk)
            }
        }
        img = new_img
    }

    acc := 0
    for _, row := range img {
        for _, pix := range row {
            if pix {
                acc += 1
            }
        }
    }

    fmt.Println("Count", acc)
}

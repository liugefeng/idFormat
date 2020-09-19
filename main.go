package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
    "strings"
)

func format_line(line *string) {
    re, _ := regexp.Compile(`\[([a-fA-F\d])\]`)
    result := re.ReplaceAllString(*line, `[0${1}]`)

    res, _ := regexp.Compile(`\]\s*\[|[\[\]]`)
    results := res.ReplaceAllString(result, ` `)
    results = strings.Trim(results, " ")

    i := 0
    lst_ids := strings.Split(results, " ")
    for _, value := range lst_ids {
        value = strings.Trim(value, " ")
        fmt.Print(value + " ")

        i++
        if i % 16 == 0 {
            fmt.Println()
        }
    }
}

func format_id(id_file string) error {
    f, err := os.Open(id_file)
    defer f.Close()

    if err != nil {
        return err
    }

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        format_line(&line)
    }

    return nil
}

func main() {
    length := len(os.Args)
    if length < 2 {
       fmt.Println("no id file found.")
        return
    }

    id_file := os.Args[1]
    fmt.Println("id file: " + id_file)

    format_id(id_file)
}


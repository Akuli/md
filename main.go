package main

import (
    "fmt"
    "strings"
    "os"
    "os/exec"
    "github.com/charmbracelet/glamour"
)

func readInput() string {
    if len(os.Args) == 2 && !strings.HasPrefix(os.Args[1], "-") {
        inputBytes, err := os.ReadFile(os.Args[1])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Can't read file: %s\n", os.Args[1])
            panic(err)
        }
        // TODO: utf8 error handling?
        return string(inputBytes[:])
    } else {
        fmt.Fprintf(os.Stderr, "Usage: %s [FILE]\n", os.Args[0])
        os.Exit(2)
        return "dummy"
    }
}

// https://stackoverflow.com/a/28706113/5806943
func pipeToPager(text string) {
    // TODO: read $PAGER if exists?
    cmd := exec.Command("less", "-R")
    cmd.Stdin = strings.NewReader(text)
    cmd.Stdout = os.Stdout

    err := cmd.Run()
    if err != nil {
        fmt.Println("Running \"less\" failed")
        panic(err)
    }
}

func main() {
    text := readInput()

    renderedText, err := glamour.Render(text, "dark")
    if err != nil {
        fmt.Println("Rendering failed")
        panic(err)
    }

    pipeToPager(renderedText)
}

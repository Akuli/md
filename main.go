package main

import (
	"fmt"
	"strings"
	"os"
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

func main() {
	text := readInput()
	out, err := glamour.Render(text, "dark")
	if err != nil {
		fmt.Println("Rendering failed")
		panic(err)
	}
	fmt.Print(out)
}

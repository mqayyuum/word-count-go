package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	byteCount bool
	wordCount bool
	lineCount bool
	charCount bool
)

func init() {
	flag.BoolVarP(&byteCount, "size", "c", false, "Get the byte count")
	flag.BoolVarP(&wordCount, "words", "w", false, "Get number of words")
	flag.BoolVarP(&lineCount, "line", "l", false, "Get number of lines")
	flag.BoolVarP(&charCount, "chars", "m", false, "Get number of chars")
}

func main() {

	// Get flags set
	flag.Parse()

	paths := flag.CommandLine.Args()

	if len(paths) == 0 {
		fmt.Printf("Usage: %s [options] filepath\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	path := paths[0]
	if path == "" {
		fmt.Println("File is not provided")
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// By default, if no flags is passed, then provide all information
	if !byteCount && !wordCount && !lineCount && !charCount {
		byteCount = true
		wordCount = true
		lineCount = true
		charCount = true
	}

	if stat, err := os.Stat(path); err == nil && byteCount {
		fmt.Printf("Filesize: %d bytes\n", stat.Size())
	}

	if wordCount {
		words := Count("words", file)
		fmt.Printf("Number of words: %d\n", words)
	}

	if lineCount {
		lines := Count("lines", file)
		fmt.Printf("Number of lines: %d\n", lines)

	}

	if charCount {
		chars := Count("chars", file)
		fmt.Printf("Number of chars: %d\n", chars)

	}
}

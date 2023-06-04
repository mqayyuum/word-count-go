package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	path      string
	byteCount bool
	wordCount bool
	lineCount bool
)

func init() {
	flag.StringVarP(&path, "file", "f", "", "File path")
	flag.BoolVarP(&byteCount, "size", "c", false, "Get the byte count")
	flag.BoolVarP(&wordCount, "words", "w", false, "Get number of words")
	flag.BoolVarP(&lineCount, "line", "l", false, "Get number of lines")
}

func main() {

	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if path == "" {
		fmt.Println("File is not provided")
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if stat, err := os.Stat(path); err == nil && byteCount {
		fmt.Printf("Filesize: %d bytes\n", stat.Size())
	}

	if wordCount {
		words := CountWords(file)
		fmt.Printf("Number of words: %d\n", words)
	}

	if lineCount {
		lines := CountLines(file)
		fmt.Printf("Number of lines: %d\n", lines)

	}
}

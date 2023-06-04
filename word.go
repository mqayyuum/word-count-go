package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// This is a wrapper function to count the data based on mode selected
func Count(op string, file *os.File) int {
	var count int
	if op == "words" {
		count = fileCounter(file, bufio.ScanWords)
	} else if op == "lines" {
		count = fileCounter(file, bufio.ScanLines)
	} else if op == "chars" {
		count = fileCounter(file, bufio.ScanRunes)
	} else {
		panic("Invalid mode selected")
	}

	return count
}

func fileCounter(file *os.File, bufFunc bufio.SplitFunc) int {
	file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufFunc)
	counter := 0
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error", err)
	}

	return counter
}

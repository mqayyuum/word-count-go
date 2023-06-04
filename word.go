package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Read the file by size provided
func ReadByChunk(file *os.File, size int) {
	b := make([]byte, size)

	for {
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(b[:readTotal]))
	}
}

// This is a wrapper function to count the data based on mode selected
func Count(op string, file *os.File) int {
	var count int
	if op == "words" {
		count = fileCounter(file, bufio.ScanWords)
	} else if op == "lines" {
		count = fileCounter(file, bufio.ScanLines)
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

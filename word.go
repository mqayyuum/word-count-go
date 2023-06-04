package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

func CountWords(file *os.File) int {
	file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	counter := 0
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error", err)
	}

	return counter
}

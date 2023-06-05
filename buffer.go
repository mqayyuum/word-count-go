package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Buffer struct {
	buffer *os.File
}

type OSReader interface {
	Lines() int
	Chars() int
	Words() int
	Size() int
}

func (b Buffer) Lines() int {
	return getCount(b, bufio.ScanLines)
}

func (b Buffer) Chars() int {
	return getCount(b, bufio.ScanRunes)
}

func (b Buffer) Words() int {
	return getCount(b, bufio.ScanWords)
}

func (b Buffer) Size() int {
	stat, _ := b.buffer.Stat()

	return int(stat.Size())
}

func Print(f *Flags, g OSReader) {
	if f.size {
		fmt.Printf("Size:\t %d bytes\n", g.Size())
	}
	if f.chars {
		fmt.Printf("Chars:\t %d\n", g.Chars())
	}
	if f.words {
		fmt.Printf("Words:\t %d\n", g.Words())
	}
	if f.lines {
		fmt.Printf("Lines:\t %d\n", g.Lines())
	}
}

func getCount(b Buffer, fn bufio.SplitFunc) int {
	b.buffer.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(b.buffer)
	scanner.Split(fn)

	counter := 0
	for scanner.Scan() {
		counter++
	}

	return counter
}

package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var flags = Flags{
	size:  false,
	words: false,
	lines: false,
	chars: false,
}

func init() {
	flag.BoolVarP(&flags.size, "size", "c", false, "Get the byte count")
	flag.BoolVarP(&flags.words, "words", "w", false, "Get number of words")
	flag.BoolVarP(&flags.lines, "line", "l", false, "Get number of lines")
	flag.BoolVarP(&flags.chars, "chars", "m", false, "Get number of chars")
}

func main() {

	// Get flags set
	flag.Parse()

	paths := flag.CommandLine.Args()

	buffer, _ := os.Stdin.Stat()

	if len(paths) == 0 && buffer.Size() == 0 {
		fmt.Printf("Usage: %s [options] filepath\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// By default, if no flags is passed, then provide all information
	if flags.IsNoneSet() {
		flags.SetAllTrue()
	}

	// Read from buffer
	if buffer.Size() != 0 {
		b := Buffer{
			buffer: os.Stdin,
		}

		Print(&flags, b)
		os.Exit(0)
	}

	// Read from file
	file, err := os.Open(paths[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f := Buffer{
		buffer: file,
	}

	Print(&flags, f)

}

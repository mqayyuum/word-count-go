package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	path      string
	byteCount bool
)

func init() {
	flag.StringVarP(&path, "file", "f", "", "File")
	flag.BoolVar(&byteCount, "c", false, "Return the byte count of file")
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
		fmt.Println("Filesize", stat.Size())
		os.Exit(1)
	}

	words := CountWords(file)
	fmt.Printf("Number of words: %d\n", words)
}

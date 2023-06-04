package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("File path: ")
	scanner.Scan()

	filepath := scanner.Text()

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	words := CountWords(file)
	fmt.Printf("Word count: %d", words)

}

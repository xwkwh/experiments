package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "Spicy jalapeno pastrami ut ham turducken."
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	fmt.Println(scanner.Text())

	scanner.Scan()
	// Count the words.
	count := 0
	for scanner.Scan() {
		fmt.Println("====", scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)

	fmt.Println(4 % 3)
}

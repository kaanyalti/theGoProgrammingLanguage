package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// input.Scan()
	// fmt.Println(input.Text())
	// press CTRL + D to end the scanner from the console
	for input.Scan() {
		fmt.Println("HERE")
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

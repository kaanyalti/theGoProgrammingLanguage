package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	* A variable can be initialized as part of its declaration.
	* If it is not explicitly initialized, it is implicitly initialized
	* to the zero value for its type, which is 0 for numeric types and the
	* empty string "" for strings
	 */
	var s, sep string

	fmt.Printf("=========================\nThis is the value of \"sep\" variable initially -> %s <-\n", sep)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("\n=========================\nThis is our string at the end\n%s\n", s)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Formatted")
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Non-formatted")
	fmt.Println(os.Args[1:])
}

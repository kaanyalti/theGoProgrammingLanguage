package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("adsfasdf/as/fdasdf/asf/filename.png"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

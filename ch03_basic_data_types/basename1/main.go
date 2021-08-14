package main

import (
	"fmt"
)

func main() {
	fmt.Println(basename("adsfasdf/as/fdasdf/asf/filename.png"))
}

func basename(s string) string {
	// Discard last "/" and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last "."
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// func basename(s string) string {
// 	slash := strings.LastIndex(s, "/") // -1 if "/" not found
// 	s = s[slash+1:]
// 	if dot := strings.LastIndex(s, "."); dot >= 0 {
// 		s = s[:dot]
// 	}
// 	return s
// }

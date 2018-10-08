package main

import (
	"fmt"
)

func basename(s string) string {
	// Remove everything preceding the last /
	for i:= len(s)-1; i >= 0; i--{
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last .
	for i:= len(s) -1; i>=0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("aded/deded/ge.got.gtogk.bonjour"))
}

	

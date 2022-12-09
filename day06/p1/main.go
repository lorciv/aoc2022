package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func start(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i:i+1] == s[j:j+1] {
				return false
			}
		}
	}
	return true
}

func detect(s string) int {
	if len(s) < 4 {
		return -1
	}
	for i := 4; i < len(s); i++ {
		if start(s[i-4 : i]) {
			return i
		}
	}
	return -1
}

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	pos := detect(string(buf))
	fmt.Println(pos)
}

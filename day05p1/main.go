package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

var stacks = make(map[int][]string)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	var lines []string
	for scan.Scan() {
		l := scan.Text()
		if l == "" {
			break
		}
		lines = append(lines, l)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// Build empty stacks
	l := lines[len(lines)-1]
	count := len(strings.Fields(l))
	fmt.Println(count)

	// Initialize the stacks
	for i := 1; i < len(lines); i++ {
		l := lines[len(lines)-1-i]
		for j := 0; j < count; j++ {
			start := 4*j + 1
			el := l[start : start+1]
			if el == " " {
				continue
			}
			stacks[j+1] = append(stacks[j+1], el)
		}
	}

	// Execute instructions
	for scan.Scan() {
		var n, src, dest int

		line := scan.Text()
		fmt.Sscanf(line, "move %d from %d to %d", &n, &src, &dest)

		for i := 0; i < n; i++ {
			e := stacks[src][len(stacks[src])-1]
			stacks[src] = stacks[src][:len(stacks[src])-1]
			stacks[dest] = append(stacks[dest], e)
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// Extract message
	var msg bytes.Buffer
	for i := 0; i < count; i++ {
		s := stacks[i+1]
		last := s[len(s)-1]
		fmt.Fprint(&msg, last)
	}
	fmt.Println(msg.String())
}

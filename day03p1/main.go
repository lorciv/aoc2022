package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func common(s1, s2 string) rune {
	for _, c := range s1 {
		if strings.ContainsRune(s2, c) {
			return c
		}
	}
	return 0
}

func priority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}

func main() {
	sum := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		c1, c2 := line[0:len(line)/2], line[len(line)/2:]
		comm := common(c1, c2)
		prio := priority(comm)
		fmt.Printf("%d (%c)\n", prio, comm)

		sum += prio
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sum: %d\n", sum)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type interval struct {
	start, end int
}

func main() {
	count := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, ",")
		if len(split) != 2 {
			log.Fatalf("expected 2 fields, got %d", len(split))
		}
		var t1, t2 interval
		fmt.Sscanf(split[0], "%d-%d", &t1.start, &t1.end)
		fmt.Sscanf(split[1], "%d-%d", &t2.start, &t2.end)

		if (t1.start >= t2.start && t1.end <= t2.end) || (t2.start >= t1.start && t2.end <= t1.end) {
			count++
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

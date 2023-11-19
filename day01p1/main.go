package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	max, sum := 0, 0
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}

		n, _ := strconv.Atoi(line)
		sum += n
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// Last group
	if sum > max {
		max = sum
	}

	fmt.Println(max)
}

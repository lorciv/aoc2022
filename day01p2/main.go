package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var sums []int

	scan := bufio.NewScanner(os.Stdin)
	cur := 0
	for scan.Scan() {
		line := scan.Text()

		if line == "" {
			sums = append(sums, cur)
			cur = 0
			continue
		}

		n, _ := strconv.Atoi(line)
		cur += n
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	if cur > 0 {
		sums = append(sums, cur)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	top3 := 0
	for _, s := range sums[0:3] {
		top3 += s
	}
	fmt.Println(top3)
}

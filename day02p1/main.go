package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock     = 0
	paper    = 1
	scissors = 2
)

func main() {
	tot := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		fields := strings.Fields(line)

		var opp int
		switch fields[0] {
		case "A":
			opp = rock
		case "B":
			opp = paper
		case "C":
			opp = scissors
		}

		var move int
		switch fields[1] {
		case "X":
			move = rock
		case "Y":
			move = paper
		case "Z":
			move = scissors
		}

		score := 0
		if move == (opp+1)%3 {
			score += 6
		} else if opp == move {
			score += 3
		}

		score += move + 1

		fmt.Println(score)
		tot += score
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("tot:", tot)
}

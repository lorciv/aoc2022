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

		var score int

		var move int
		switch fields[1] {
		case "X":
			// lose
			move = (opp + 2) % 3
		case "Y":
			// draw
			move = opp
			score += 3
		case "Z":
			// win
			move = (opp + 1) % 3
			score += 6
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func read_inputs() ([]string, [][][]string) {
	var today_number int = 1
	_, program_name, _, ok := runtime.Caller(1) // /some/path/adventofcode2021/day_1.go
	if ok {
		fields := strings.Split(program_name, "/")
		base_name := fields[len(fields)-1]                 // day_1.go
		drop_extension := strings.Split(base_name, ".")[0] // day_1
		today_number, _ = strconv.Atoi(strings.Split(drop_extension, "_")[1])
	}
	input_file := fmt.Sprintf("inputs/2021_%d_test.in", today_number)

	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var linecounter int
	var cardcounter int
	var cards [][][]string
	//cards := [][][]string{}
	var lastcard [][]string

	var drawnnums []string

	for scanner.Scan() {
		linecounter += 1
		line := scanner.Text()
		if linecounter == 1 {
			drawnnums = strings.Split(line, ",")
			//fmt.Println(drawnnums)
			continue
		}
		if linecounter == 2 {
			continue
		}

		//fmt.Println(linecounter)
		if line != "" {
			//fmt.Println(cardcounter)
			lastcard = append(lastcard, strings.Split(line, " "))
			//fmt.Println(lastcard)
		} else {
			//fmt.Println("cards before", cards)
			//fmt.Println("lastcard before", lastcard)
			cards = append(cards, lastcard)
			//fmt.Println(lastcard)
			lastcard = make([][]string, 0)
			//fmt.Println("lastcard after", lastcard)
			//fmt.Println("cards after", cards)
			cardcounter += 1
		}
	}
	fmt.Println("cards", cards)
	return drawnnums, cards
}

func testcard(card [][]string) bool {
	for i := 0; i < 5; i++ {
		if card[i][0] == card[i][1] && card[i][1] == card[i][2] && card[i][2] == card[i][3] && card[i][3] == card[i][4] && card[i][4] == "X" {
			fmt.Println("winning here", card[i][0])
			return true
		}
		if card[0][i] == card[1][i] && card[1][i] == card[2][i] && card[2][i] == card[3][i] && card[3][i] == card[4][i] && card[4][i] == "X" {
			fmt.Println("winning there", i)
			fmt.Println(card[0][i], card[1][i], card[2][i], card[3][i], card[4][i])
			fmt.Println(card)
			return true
		}
	}
	return false
}

func calculatecard(card [][]string, lastnum string) int {
	var sum int
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card[i][j] != "" {
				num, _ := strconv.Atoi(card[i][j])
				sum += num
			}
		}
	}
	lastnumdec, _ := strconv.Atoi(lastnum)
	score = sum * lastnumdec
	//fmt.Println(score)
	return score
}

func step_1(drawnnums []string, cards [][][]string) int {
	//fmt.Println(drawnnums)
	//fmt.Println(cards[0][0])

	for _, number := range drawnnums {
		fmt.Println("number", number)
		for i := 0; i < len(cards); i++ {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if cards[i][j][k] == number {
						cards[i][j][k] = "X"
						if testcard(cards[i]) {
							fmt.Println(cards[i], number)
							return calculatecard(cards[i], number)
						}
						//fmt.Println(cards[i][j][k])
					}
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("Score", step_1(read_inputs()))
}

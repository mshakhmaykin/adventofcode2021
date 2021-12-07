package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func read_inputs() ([]int, [][][]int) {
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
	var cards [][][]int
	//cards := [][][]string{}
	var lastcard [][]int

	var drawnnums []int

	for scanner.Scan() {
		linecounter += 1
		line := scanner.Text()
		if linecounter == 1 {
			drawnnumsstr := strings.Split(line, ",")
			for _, v := range drawnnumsstr {
				n, _ := strconv.Atoi(v)
				drawnnums = append(drawnnums, n)
			}
			//fmt.Println(drawnnums)
			continue
		}
		if linecounter == 2 {
			continue
		}

		//fmt.Println(linecounter)
		if line != "" {
			//fmt.Println(cardcounter)
			tmp := strings.Fields(line)
			lineint := []int{}
			for _, v := range tmp {
				n, _ := strconv.Atoi(v)
				lineint = append(lineint, n)
			}
			lastcard = append(lastcard, lineint)
			//fmt.Println(lastcard)
		} else {
			//fmt.Println("cards before", cards)
			//fmt.Println("lastcard before", lastcard)
			cards = append(cards, lastcard)
			//fmt.Println(lastcard)
			lastcard = make([][]int, 0)
			//fmt.Println("lastcard after", lastcard)
			//fmt.Println("cards after", cards)
			cardcounter += 1
		}
	}
	//fmt.Println("cards", cards)
	return drawnnums, cards
}

func testcard(card [][]int) bool {
	for i := 0; i < 5; i++ {
		if card[i][0] == card[i][1] && card[i][1] == card[i][2] && card[i][2] == card[i][3] && card[i][3] == card[i][4] && card[i][4] == -1 {
			//fmt.Println("winning here", card[i][0])
			return true
		}
		if card[0][i] == card[1][i] && card[1][i] == card[2][i] && card[2][i] == card[3][i] && card[3][i] == card[4][i] && card[4][i] == -1 {
			//fmt.Println("winning there", i)
			//fmt.Println(card[0][i], card[1][i], card[2][i], card[3][i], card[4][i])
			//fmt.Println(card)
			return true
		}
	}
	return false
}

func calculatecard(card [][]int, lastnum int) int {
	var sum int
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card[i][j] != -1 {
				sum += card[i][j]
			}
		}
	}

	score = sum * lastnum
	//fmt.Println(score)
	return score
}

func step_1(drawnnums []int, cards [][][]int) int {
	//fmt.Println(drawnnums)
	//fmt.Println(cards[0][0])

	for _, number := range drawnnums {
		fmt.Println("number", number)
		for i := 0; i < len(cards); i++ {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if cards[i][j][k] == number {
						cards[i][j][k] = -1
						//fmt.Println(cards[i])
						if testcard(cards[i]) {
							//fmt.Println(cards[i], number)
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_inputs() []int {
	day := 1 // redo by capturing program filename
	input_file := fmt.Sprintf("inputs/2021_%d.in", day)
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines)
	var inputs []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return inputs
}

func step_1(values []int) {
	increases := 0
	previous := 1000000000000000

	for _, num := range values {
		if num > previous {
			increases += 1
		}
		previous = num
	}

	fmt.Println("Total increases", increases)
}

func step_2(values []int) {
	increases := 0
	var sums []int
	for i := 0; i < len(values)-2; i++ {
		sum := values[i] + values[i+1] + values[i+2]
		sums = append(sums, sum)
	}
	previous := 10000000000
	for _, sum := range sums {
		if sum > previous {
			increases += 1
		}
		previous = sum
	}
	fmt.Println("Sum increases", increases)
}

func main() {
	values := read_inputs()
	step_1(values)
	step_2(values)
}

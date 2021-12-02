package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func read_inputs() map[string]int {
	var today_number int = 1
	_, program_name, _, ok := runtime.Caller(1) // /some/path/adventofcode2021/day_1.go
	if ok {
		fields := strings.Split(program_name, "/")
		base_name := fields[len(fields)-1]                 // day_1.go
		drop_extension := strings.Split(base_name, ".")[0] // day_1
		today_number, _ = strconv.Atoi(strings.Split(drop_extension, "_")[1])
	}

	input_file := fmt.Sprintf("inputs/2021_%d.in", today_number)
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputs = make(map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction := line[0]
		units, _ := strconv.Atoi(line[1])
		inputs[direction] += units
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return inputs
}

// duplicating read inputs as I liked how outputs are done in the first part
func read_inputs_2() []string {
	var today_number int = 1
	_, program_name, _, ok := runtime.Caller(1) // /some/path/adventofcode2021/day_1.go
	if ok {
		fields := strings.Split(program_name, "/")
		base_name := fields[len(fields)-1]                 // day_1.go
		drop_extension := strings.Split(base_name, ".")[0] // day_1
		today_number, _ = strconv.Atoi(strings.Split(drop_extension, "_")[1])
	}

	input_file := fmt.Sprintf("inputs/2021_%d.in", today_number)
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputs []string

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return inputs
}

func step_1(values map[string]int) {
	depth := values["down"] - values["up"]
	horizontal := values["forward"]
	product := depth * horizontal
	fmt.Println("Product 1:", product)
}

func step_2(values []string) {
	horizontal, depth, aim := 0, 0, 0
	for _, command := range values {
		line := strings.Split(command, " ")
		direction := line[0]
		units, _ := strconv.Atoi(line[1])
		switch direction {
		case "down":
			aim += units
		case "up":
			aim -= units
		case "forward":
			horizontal += units
			depth += aim * units
		}
	}

	product := depth * horizontal
	fmt.Println("Product 2:", product)
}

func main() {
	values1 := read_inputs()
	values2 := read_inputs_2()
	step_1(values1)
	step_2(values2)
}

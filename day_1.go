package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func read_inputs() []int {
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

	for i, num := range values {
		if i > 0 {
			if num > values[i-1] {
				increases += 1
			}
		}
	}

	fmt.Println("Total increases", increases)
}

func step_2(values []int) {
	increases := 0
	var sums []int
	for i := 0; i < len(values)-2; i++ {
		sum := values[i] + values[i+1] + values[i+2]
		sums = append(sums, sum)
		if i > 0 {
			if sum > sums[len(sums)-2] {
				increases += 1
			}
		}
	}

	fmt.Println("Sum increases", increases)
}

func main() {
	values := read_inputs()
	step_1(values)
	step_2(values)
}

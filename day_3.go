package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func read_inputs() map[int]map[string]int {
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
	var inputs = make(map[int]map[string]int)

	for i := 0; i < 12; i++ { // initialize with the length of the string; cannot do in the next loop as it will recreate the map after each processed line
		inputs[i] = make(map[string]int)
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for idx, elem := range line {
			inputs[idx][elem] += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return inputs
}

func step_1(values map[int]map[string]int) {
	var gamma []string
	var epsilon []string
	var keys []int

	// sorting keys
	for idx, _ := range values {
		keys = append(keys, idx)
	}
	sort.Ints(keys)

	for idx := 0; idx < len(keys); idx++ {
		elem := values[idx]
		if elem["0"] >= elem["1"] {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		} else {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		}
	}
	//fmt.Println(gamma)
	//fmt.Println(epsilon)
	gammabin := strings.Join(gamma, "")
	epsilonbin := strings.Join(epsilon, "")
	gammadec, _ := strconv.ParseInt(gammabin, 2, 16)
	epsilondec, _ := strconv.ParseInt(epsilonbin, 2, 16)
	powercons := gammadec * epsilondec
	fmt.Println(powercons)
}

func main() {
	values := read_inputs()
	//fmt.Println(values)
	step_1(values)
}

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

func read_inputs1() map[int]map[string]int {
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
		//for i := 0; i < 5; i++ { // initialize with the length of the string; cannot do in the next loop as it will recreate the map after each processed line
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

func read_inputs2() []string {
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
		if elem["0"] > elem["1"] {
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
	//return gamma, epsilon
}

func step_2(values []string) {
	//fmt.Println(values)
	var zeros_list []string
	var ones_list []string
	var zeros_count int
	var ones_count int
	//	var co2list []string
	//	var oxlist []string

	var oxvalues []string = values

	for idx := 0; idx < 12; idx++ {
		//fmt.Println(zeros_list)
		//fmt.Println(ones_list)
		zeros_list = nil
		ones_list = nil
		zeros_count = 0
		ones_count = 0
		oxlist := oxvalues
		//	fmt.Println("oxlist", oxlist)
		for _, v := range oxlist {
			if strings.Split(v, "")[idx] == "0" {
				zeros_list = append(zeros_list, v)
				zeros_count += 1
				//fmt.Println(zeros_count)
			} else {
				ones_list = append(ones_list, v)
				ones_count += 1
				//fmt.Println(ones_count)
			}
		}

		if ones_count >= zeros_count {
			//	oxlist = ones_list
			//	co2list = zeros_list
			//	fmt.Println("more ones", ones_count)
			oxvalues = ones_list
		} else {
			//	co2list = ones_list
			//	oxlist = zeros_list
			//	fmt.Println("more zeros", zeros_count)
			oxvalues = zeros_list
		}

		if len(oxvalues) == 1 {
			break
		}
		//fmt.Println(oxlist)
	}
	oxdec, _ := strconv.ParseInt(oxvalues[0], 2, 16)
	fmt.Println(oxdec)

	var co2values []string = values
	for idx := 0; idx < 12; idx++ {
		zeros_list = nil
		ones_list = nil
		zeros_count = 0
		ones_count = 0
		co2list := co2values
		fmt.Println("co2list", co2list)
		for _, v := range co2list {
			if strings.Split(v, "")[idx] == "0" {
				zeros_list = append(zeros_list, v)
				zeros_count += 1
			} else {
				ones_list = append(ones_list, v)
				ones_count += 1
			}
		}

		if ones_count >= zeros_count {
			//	oxlist = ones_list
			//	co2list = zeros_list
			co2values = zeros_list
			fmt.Println("more ones", ones_count)
		} else {
			//	co2list = ones_list
			//	oxlist = zeros_list
			co2values = ones_list
			fmt.Println("more zeros", zeros_count)
		}
		if len(co2values) == 1 {
			break
		}
		//fmt.Println(oxlist)
	}
	co2dec, _ := strconv.ParseInt(co2values[0], 2, 16)
	fmt.Println(co2dec)
	fmt.Println("product", co2dec*oxdec)
}

func main() {
	//values1 := read_inputs1()
	values2 := read_inputs2()
	//fmt.Println(values1)
	//step_1(values1)
	//fmt.Println(g, e)
	step_2(values2)
}

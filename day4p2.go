package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strconv"
	"strings"
)

func readData(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, scanner.Err()
}

func displayData(data []string) {
	fmt.Println("Here is the input data")
	fmt.Println("----------------------------------------------------------")

	for _, value := range data {
		fmt.Println(value)
	}

	fmt.Println("----------------------------------------------------------\n")
}

func processData(data []string) {
	start, _ := strconv.Atoi(data[0])
	end, _ := strconv.Atoi(data[1])

	count := 0

	for i := start; i <= end; i++ {
		ok := isValid(i)
		if ok {
			fmt.Println("Good: ", i)
			count += 1
		}
	}

	fmt.Println("There are ", count, " possible passwords.")
}

func isValid(input int) bool {
	tostring := strconv.Itoa(input)
	data := strings.Split(tostring, "")
	length := len(data)

	if length != 6 {
		return false
	}

	hasPair := false
	for i := 0; i < length - 1; i++ {
		v1, _ := strconv.Atoi(data[i])
		v2, _ := strconv.Atoi(data[i + 1])
		if v2 < v1 {
			return false
		}

		if v1 == v2 {
			left, right := true, true

			if i == 0 { // Base case - far left
			} else if i > 0 && data[i] != data[i - 1] {
			} else {
				left = false
			}

			if i == length - 2 { // Base case - far right
			} else if i < length - 2 && data[i + 1] != data[i + 2] {
			} else {
				right = false
			}

			if left && right {
				hasPair = true
			}
		}
	}

	if ! hasPair {
		return false
	}
	return true
}

func main() {
	displayPtr := flag.Bool("d", false, "Display input text")
	testPtr := flag.Bool("t", false, "Test a single number")

	flag.Parse()

	argv := flag.Args()

	data, err := readData(argv[0])
	if err != nil {
		panic(err)
	}

	if *displayPtr {
		displayData(data)
	}

	if *testPtr {
		for _, v := range data {
			input, _ := strconv.Atoi(v)
			ok := isValid(input)
			if ok {
				fmt.Println("Yes: ", v, "\n")
			} else {
				fmt.Println("No: ", v, "\n")
			}
		}
	} else {
		input := strings.Split(data[0], "-")
		processData(input)
	}
}

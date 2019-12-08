package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strings"
	"strconv"
	"errors"
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

func testData(data []string) {

	for _, line := range data {
		raw := strings.Split(line, " ")

		result, err := processLine(raw[0])
		if err != nil {
			fmt.Println(err)
		}

		expected := raw[1]

		if result == expected {
			fmt.Println("Sucess: ", result, " == ", expected)
		} else {
			fmt.Println("Failure: ", result, " != ", expected)
		}
	}
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
	for _, line := range data {
		pline, err := processLine(line)
		if err != nil {
			fmt.Println(err)
		} else {
			results := strings.Split(pline, ",")
			fmt.Println(results)
			fmt.Println(results[0])
		}
	}
}

func processLine(line string) (string, error) {
	tape := strings.Split(line, ",")
	length := len(tape)

	for index := 0; index <= length; index += 4 {
		var i1, i2, v1, v2, loc, sum int
		var err error

		if index + 4 <= length {
			i1, err = strconv.Atoi(tape[index + 1])
			i2, err = strconv.Atoi(tape[index + 2])
			v1, err = strconv.Atoi(tape[i1])
			v2, err = strconv.Atoi(tape[i2])
			loc, err = strconv.Atoi(tape[index +3])

			if err != nil {
				return "", err
			}

			switch op := tape[index]; op {
			case "1":
				sum = v1 + v2

			case "2":
				sum = v1 * v2

			case "99":
				return strings.Join(tape, ","), nil

			default:
				return "", errors.New(fmt.Sprintf("Invalid op: %s", op))
			}
		} else {
			if tape[index] == "99" {
				return strings.Join(tape, ","), nil
			} else {
				return "", errors.New("Invalid operation")
			}
		}

		if loc < length {
			tape[loc] = strconv.Itoa(sum)
		} else {
			return "", errors.New("Index out of bounds")
		}
	}

	return strings.Join(tape, ","), nil
}

func main() {
	displayPtr := flag.Bool("d", false, "Display input text")
	testPtr := flag.Bool("t", false, "Run testing data")

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
		testData(data)
	} else {
		processData(data)
	}
}

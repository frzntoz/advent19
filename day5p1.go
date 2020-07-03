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

		tmp := strings.Split(raw[0], ",")
		i, err := strconv.Atoi(tmp[1])
		j, err := strconv.Atoi(tmp[2])
		if err != nil {
			continue
		}

		result, err := processLine(raw[0], raw[1], i, j)
		if err == nil {
			fmt.Println("Success!", result, "==", raw[2], "\n")
		} else {
			fmt.Println(err, "\n")
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

func processLine(line, expected string, i, j int) (string, error) {
	tape := strings.Split(line, ",")
	length := len(tape)

	// modify tape as per instructions
	tape[1] = strconv.Itoa(i)
	tape[2] = strconv.Itoa(j)

	index := 0
	for index <= length {
		var op string
		var i1, i2, v1, v2, loc, sum int
		var err error

		if index + 4 <= length {
			op = tape[index]
			i1, err = strconv.Atoi(tape[index + 1])
			i2, err = strconv.Atoi(tape[index + 2])
			v1, err = strconv.Atoi(tape[i1])
			v2, err = strconv.Atoi(tape[i2])
			loc, err = strconv.Atoi(tape[index +3])

			if err != nil {
				return "", err
			}

			switch op {
			case "1":
				sum = v1 + v2

			case "2":
				sum = v1 * v2

			case "3":
				//

			case "4":
				//

			case "99":
				return "", errors.New("Unable to find solution")

			default:
				return "", errors.New(fmt.Sprintf("Invalid op: %s", op))
			}

			result := strconv.Itoa(sum)
			if result == expected {
				return "Success!", nil
			}
		} else {
			if tape[index] == "99" {
				return "", errors.New("Unable to find solution")
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

	return "", errors.New("Unable to find solution")
}

func main() {
	displayPtr := flag.Bool("d", false, "Display input text")
	testPtr := flag.Bool("t", false, "Run testing data")
	searchPtr := flag.String("s", "", "Expected sum")

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
		if *searchPtr == "" {
			panic("Search value required")
		}
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				_, err := processLine(data[0], *searchPtr, i, j)
				if err == nil {
					v := 100 * i + j
					fmt.Println(v)
				}
			}
		}
	}
}

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
			hasPair = true
		}
	}

	if ! hasPair {
		return false
	}
	return true
}

func main() {
	displayPtr := flag.Bool("d", false, "Display input text")

	flag.Parse()

	argv := flag.Args()

	data, err := readData(argv[0])
	if err != nil {
		panic(err)
	}

	if *displayPtr {
		displayData(data)
	}

	input := strings.Split(data[0], "-")
	processData(input)
}

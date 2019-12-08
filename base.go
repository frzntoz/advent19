package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
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

	processData(data)
}

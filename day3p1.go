package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strings"
	"strconv"
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
	one := strings.Split(data[0], ",")
	two := strings.Split(data[1], ",")

	wireOnePath := make(map[string]int)

	processWireOne(one, wireOnePath)

	minDistance := processWireTwo(two, wireOnePath)

	if minDistance > 0 {
		fmt.Println("Closest intersection distance: ", minDistance)
	} else {
		fmt.Println("No wires crossed.")
	}
}

func processWireOne(data []string, path map[string]int) {

	x, y := 0, 0

	for _, v := range data {
		var coord string
		direction := v[:1]
		distance, err := strconv.Atoi(v[1:])
		if err != nil {
			continue
		}

		switch direction {
		case "U":
			j := y + 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", x, j)
				path[coord] = 1
				j += 1
			}
			y = j - 1

		case "D":
			j := y - 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", x, j)
				path[coord] = 1
				j -= 1
			}
			y = j + 1

		case "R":
			i := x + 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", i, y)
				path[coord] = 1
				i += 1
			}
			x = i - 1

		case "L":
			i := x - 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", i, y)
				path[coord] = 1
				i -= 1
			}
			x = i + 1
		}
	}
}

func processWireTwo(data []string, wireOne map[string]int) int {

	path := make(map[string]int)
	x, y, min := 0, 0, 0

	for _, v := range data {
		var coord string
		direction := v[:1]
		distance, err := strconv.Atoi(v[1:])
		if err != nil {
			continue
		}

		switch direction {
		case "U":
			j := y + 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", x, j)
				_, ok1 := wireOne[coord]
				_, ok2 := path[coord]
				if ok1 && !ok2 {
					dist := myAbs(x) + myAbs(j)
					if min > 0 && dist < min {
						min = dist
					}

					if min == 0 {
						min = dist
					}
				}
				path[coord] = 1
				j += 1
			}
			y = j - 1

		case "D":
			j := y - 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", x, j)
				_, ok1 := wireOne[coord]
				_, ok2 := path[coord]
				if ok1 && !ok2 {
					dist := myAbs(x) + myAbs(j)
					if min > 0 && dist < min {
						min = dist
					}

					if min == 0 {
						min = dist
					}
				}
				path[coord] = 1
				j -= 1
			}
			y = j + 1

		case "R":
			i := x + 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", i, y)
				_, ok1 := wireOne[coord]
				_, ok2 := path[coord]
				if ok1 && !ok2 {
					dist := myAbs(i) + myAbs(y)
					if min > 0 && dist < min {
						min = dist
					}

					if min == 0 {
						min = dist
					}
				}
				path[coord] = 1
				i += 1
			}
			x = i - 1

		case "L":
			i := x - 1
			for count := 0; count < distance; count++ {
				coord = fmt.Sprintf("%d,%d", i, y)
				_, ok1 := wireOne[coord]
				_, ok2 := path[coord]
				if ok1 && !ok2 {
					dist := myAbs(i) + myAbs(y)
					if min > 0 && dist < min {
						min = dist
					}

					if min == 0 {
						min = dist
					}
				}
				path[coord] = 1
				i -= 1
			}
			x = i + 1
		}
	}

	return min
}

func myAbs(value int) int {
	if value < 0 {
		return value * -1
	} else {
		return value
	}
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

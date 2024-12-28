package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define a Coordinate struct
type Coordinate struct {
	X, Y int
}

func main() {
	// Open the file
	file, err := os.Open("coordinates/altis.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Map to store Coordinate -> z mapping
	lookup := make(map[Coordinate]float64)

	// Buffered file reading
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read line and split into x, y, z
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		z, err3 := strconv.ParseFloat(parts[2], 32)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Invalid data:", line)
			continue
		}

		// Create Coordinate struct and store in map
		key := Coordinate{X: x, Y: y}
		lookup[key] = z
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Interactive querying
	var x, y int
	fmt.Println("Enter x and y coordinates (or type 'exit' to quit):")
	for {
		fmt.Print("x: ")
		if _, err := fmt.Scan(&x); err != nil {
			break
		}
		fmt.Print("y: ")
		if _, err := fmt.Scan(&y); err != nil {
			break
		}

		// Create Coordinate struct and lookup z
		key := Coordinate{X: x, Y: y}
		if z, found := lookup[key]; found {
			fmt.Printf("z value for (%d, %d) is %f\n", x, y, z)
		} else {
			fmt.Printf("No z value found for (%d, %d)\n", x, y)
		}
	}
}

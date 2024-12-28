package parser

import (
	"a3-go-coordinate-server/models"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var CoordinatesMap = make(map[models.Coordinate]float64)

func GetZ(c models.Coordinate, lookup map[models.Coordinate]float64) float64 {
	if z, found := lookup[c]; found {
		return math.Round(z*100) / 100
	}
	return 0
}

func ReadCoordinatesFromFile() {
	start := time.Now()
	fmt.Println("Reading coordinates from file...")

	// Open the file
	file, err := os.Open("coordinates/altis.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

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
		key := models.Coordinate{X: x, Y: y}
		CoordinatesMap[key] = z
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Coordinates read from file:", len(CoordinatesMap))
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}

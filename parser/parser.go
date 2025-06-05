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

var MapsMap = make(map[string]map[models.Coordinate]float32)

func GetZ(c models.Coordinate, lookup map[models.Coordinate]float32) float32 {
	if z, found := lookup[c]; found {
		return float32(math.Round(float64(z)*100) / 100)
	}
	return 0
}

func ReadCoordinatesFromFile(mapName string) {
	start := time.Now()
	fmt.Println("Reading coordinates from file... ", mapName)

	// Open the file
	filePath := fmt.Sprintf("coordinates/%s.txt", mapName)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	MapsMap[mapName] = make(map[models.Coordinate]float32)
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
		MapsMap[mapName][key] = float32(z)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Coordinates read from file:", len(MapsMap[mapName]))
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	//"strings"
	//"strconv"
	//"math"
	//"regexp"
	//"sort"
	//"encoding/csv"
)

type point struct {
	row,column int
}

type path struct {
	stepCount int
	currentPoint point
	intValue int
	byteValue byte

}

var heightMap = make(map[string]int)


// I miss arry = arry.remove()
func RemoveIndex(s []path, index int) []path {
	return append(s[:index], s[index+1:]...)
}


func ReadFile(r string) []string{

	var arry []string

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		arry = append(arry, scanner.Text())
	}
		
	return arry
}


func getNeighbours(p path, width, height int) []point {
	row, column := p.currentPoint.row, p.currentPoint.column

	neighbours := []point{}

	if row > 0 {
		neighbours = append(neighbours, point{row - 1, column})
	}

	if column > 0 {
		neighbours = append(neighbours, point{row, column - 1})
	}

	if row < width-1 {
		neighbours = append(neighbours, point{row + 1, column})
	}

	if column < height-1 {
		neighbours = append(neighbours, point{row, column + 1})
	}

	return neighbours
}


func getCompeletedPaths(paths []path, puzzle_arry []string) []path{

	width := len(puzzle_arry[0])
	height := len(puzzle_arry)

	stepCounts := make(map[point]int)

	compeletedPaths := []path{}

	// No while loops in golang and can't alter the size of paths inside a normal for loop
	// So keep i at 0 and go round until we run out of paths
	for i := 0; i < len(paths); i = 0 {

		currentPointentPath:= paths[0]
		paths = RemoveIndex(paths, 0)

		// Using just the int value we end up with short paths (confussion between z and E)
		if currentPointentPath.byteValue == 'E' {
			compeletedPaths = append(compeletedPaths, currentPointentPath)
		} else {

			neighbours := getNeighbours(currentPointentPath, width, height)

			for _, neighbour := range neighbours {
				neighbourByteValue := puzzle_arry[neighbour.column][neighbour.row]

				neighbourIntValue := heightMap[string(neighbourByteValue)]
				if neighbourByteValue == 'E' {
					neighbourIntValue = heightMap[string("z")]
				}

				if (neighbourIntValue-currentPointentPath.intValue) <= 1 {

					currentPointentCount, exists := stepCounts[neighbour]

					if ! exists || currentPointentCount > currentPointentPath.stepCount+1 {

						stepCounts[neighbour] = currentPointentPath.stepCount + 1

						paths = append(paths, path{
							stepCount: currentPointentPath.stepCount + 1,
							currentPoint:      neighbour,
							intValue: neighbourIntValue,
							byteValue: neighbourByteValue,
						})
					} 
					
				}
			}
		}
	}

	return compeletedPaths
}


func partOneandTwo(puzzle_arry []string) {

	var startingPath path

	puzzle_one_paths := []path{}

	puzzle_two_paths := []path{}

	for row := 0; row < len(puzzle_arry[0]); row++ {
		for coloum := 0; coloum < len(puzzle_arry); coloum++ {
			if puzzle_arry[coloum][row] == 'S' {
				startingPath = path{
					currentPoint: point{row, coloum},
					stepCount: 0,
					intValue: 1,
					byteValue: 'a',
				}

				puzzle_one_paths = append(puzzle_one_paths, startingPath)
				puzzle_two_paths = append(puzzle_one_paths, startingPath)

			}

			if puzzle_arry[coloum][row] == 'a' {
				startingPath = path{
					currentPoint: point{row, coloum},
					stepCount: 0,
					intValue: 1,
					byteValue: 'a',
				}

				puzzle_two_paths = append(puzzle_two_paths, startingPath)

			}
		}
	}

	compeletedPaths := getCompeletedPaths(puzzle_one_paths, puzzle_arry)

	min := compeletedPaths[0].stepCount
	for _, value := range compeletedPaths {
		if value.stepCount < min {
			min = value.stepCount
		}
	}

	fmt.Printf("Part 1 answer is: %d\n", min)

	// Lazy copy + pasta strikes again
	compeletedPaths = getCompeletedPaths(puzzle_two_paths, puzzle_arry)

	min = compeletedPaths[0].stepCount
	for _, value := range compeletedPaths {
		if value.stepCount < min {
			min = value.stepCount
		}
	}

	fmt.Printf("Part 2 answer is: %d\n", min)
}



func main() {

	count := 1
	for char := 'a'; char <= 'z'; char++ {  
		heightMap[string(char)] = count
		count += 1
	}

	puzzle_arry := ReadFile("puzzle_12.txt")

	partOneandTwo(puzzle_arry)


}



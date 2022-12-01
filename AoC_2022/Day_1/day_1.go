package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
)

func ReadInts(r string) ([]int, error) {
	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result []int
	var count int
	count = 0
	for scanner.Scan() {
		if scanner.Text() != ""{
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return result, err
			}
			count += x
		} else {
			result = append(result, count)
			count = 0
		}
	}
	return result, scanner.Err()
}

// I miss arry = arry.remove()
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}


func main() {

	puzzle_one_array, err := ReadInts("puzzle_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	var max, sum int
	var max_array = []int{0,0,0}

	max = puzzle_one_array[0]
	for _, value := range puzzle_one_array {
		if value > max {
			max = value
		}
	}

	for outer_idx,outer_value := range max_array {
		var temp_value, indx_to_remove int
		temp_value = outer_value
		for inner_idx, inner_value := range puzzle_one_array {
			if inner_value > temp_value {
				temp_value = inner_value
				indx_to_remove = inner_idx
			}
		}
		max_array[outer_idx] = temp_value
		sum += temp_value
		puzzle_one_array = RemoveIndex(puzzle_one_array, indx_to_remove)
	}

	fmt.Printf( "Puzzle 1 answer is: %d\n", max );

	fmt.Printf( "Puzzle 2 answer is: %d\n", sum );

}
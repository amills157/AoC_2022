package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	//"sort"
)

func strToInts(input string) ( []int) {

	elf_min_max := strings.Split(input, "-")

	elf_ints := make([]int, len(elf_min_max))

	for idx, str := range elf_min_max{
		elf_ints[idx], _ = strconv.Atoi(str)
	}

	return elf_ints
}


func partOne( arry []string) int{

	var sum int

	for _, value := range arry {

		split_value := strings.Split(value, ",")
		
		elf_1_ints := strToInts(split_value[0])

		elf_2_ints := strToInts(split_value[1])

		if elf_1_ints[0] >= elf_2_ints[0]{
			if elf_1_ints[1] <= elf_2_ints[1]{
				sum += 1
				// Sneaky uses cases where both if checks true so we if reach here skip next check
				continue
			}
			
		}

		if elf_2_ints[0] >= elf_1_ints[0]{
			if elf_2_ints[1] <= elf_1_ints[1]{
				sum += 1
			}
		}

	}
		
	return sum
}

func partTwo( arry []string) int{

	var sum int

	for _, value := range arry {

		split_value := strings.Split(value, ",")
		
		elf_1_ints := strToInts(split_value[0])

		elf_2_ints := strToInts(split_value[1])

		if elf_1_ints[1] < elf_2_ints[0] {
			// Max less than min - No overlap
			continue
		} else if elf_2_ints[1] < elf_1_ints[0] {
			// Max less than min - No overlap
			continue
		} else {
			sum += 1
		}

	}
		
	return sum
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


func main() {

	puzzle_arry := ReadFile("puzzle_4.txt")

	pt1_answer := partOne(puzzle_arry)

	pt2_answer := partTwo(puzzle_arry)


	fmt.Printf( "Puzzle 1 answer is: %d\n", pt1_answer);

	fmt.Printf( "Puzzle 2 answer is: %d\n", pt2_answer);
	

}
package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
)

var priorities = make(map[string]int)


// I miss set() / numpy arrrays
func intersection(s1, s2 []string) (inter []string) {
    hash := make(map[string]bool)
    for _, e := range s1 {
        hash[e] = true
    }
    for _, e := range s2 {
        // If elements present in the hashmap then append intersection list.
        if hash[e] {
            inter = append(inter, e)
        }
    }
    //Remove dups from slice.
    inter = removeDups(inter)
    return
}

//Remove dups from slice.
func removeDups(elements []string)(nodups []string) {
    encountered := make(map[string]bool)
    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return
}

func partOne( arry []string) int{

	var sum int

	for _, value := range arry {

		half_len := (len(value) / 2)

		first_half_priorities := value[:half_len]
		second_half_priorities := value[half_len:]
		
		common_items := intersection(strings.Split(first_half_priorities, ""), strings.Split(second_half_priorities, ""))

		for _, value := range common_items {
			sum += priorities[value]
		}

	}
		
	return sum
}

func partTwo( arry []string) int{

	var sum, count int

	sum = 0
	count = 1

	var grouped_lines []string

	for _, value := range arry {

		grouped_lines = append(grouped_lines, value)

		// can't use idk as we start at 0 which buggers the calculation
		if count % 3 == 0{

			temp := intersection(strings.Split(grouped_lines[0], ""), strings.Split(grouped_lines[1], ""))

			common_items := intersection(temp, strings.Split(grouped_lines[2], ""))

			for _, value := range common_items {
				sum += priorities[value]
			}

			grouped_lines = nil
		}

		count += 1
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

	// bugger typing that out manually
	var count int
	count = 1
	for char := 'a'; char <= 'z'; char++ {  
		priorities[string(char)] = count
		count += 1
	}

	for char := 'A'; char <= 'Z'; char++ {  
		priorities[string(char)] = count
		count += 1
	}

	puzzle_arry := ReadFile("puzzle_3.txt")

	pt1_score := partOne(puzzle_arry)

	pt2_score := partTwo(puzzle_arry)


	fmt.Printf( "Puzzle 1 answer is: %d\n", pt1_score);

	fmt.Printf( "Puzzle 2 answer is: %d\n", pt2_score)
	

}
package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	//"strings"
	"strconv"
	"math"
	//"regexp"
	//"sort"
	//"encoding/csv"
)


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

func toDecimal(str string) int {
	result := 0
	m := 1
	for i := len(str) - 1; i >= 0; i-- {
		digit := 0
		switch str[i] {
		case '1':
			digit = 1
		case '2':
			digit = 2
		case '-':
			digit = -1
		case '=':
			digit = -2
		}
		result += digit * m
		m *= 5
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func partOne(puzzle_arry []string) string {

	sum := 0
	result := ""
	by_5_value := 5

	for _, value := range puzzle_arry{
		sum += toDecimal(value)
	}
	
	current_value := sum
	
	// We need a *5 per digit - as is I kept getting a wrong answer, so the +5 works / fixed it
	for i := len(strconv.Itoa(sum)) + 5; i >= 0; i-- {
		by_5_value *= 5 
	}

	// Loop down through the digit places / by 5 value
	for ; by_5_value > 0; by_5_value /= 5 {
		snafu_value := 0
		max_value := math.MaxInt
		// SNAFU values go from -2 through to 2
		for j := -2; j <= 2; j++ {
			if Abs(current_value-(j * by_5_value)) < max_value {
				snafu_value = j
				max_value = Abs(current_value - (j * by_5_value))
			}
		}
		
		current_value -= bestt * by_5_value

		switch snafu_value {
		case -1:
			result += "-"
		case -2:
			result += "="
		default:
			result += strconv.Itoa(snafu_value)
		}
	}

	return result
}

func main() {

	puzzle_arry := ReadFile("puzzle_25.txt")
	
	pt1_answer :=  partOne(puzzle_arry)

	fmt.Printf("Part 1 answer is: %s\n", pt1_answer)
}
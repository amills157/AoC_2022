package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	//"math"
	"regexp"
	//"sort"
	//"encoding/csv"
)

type monkeyStruct struct {
	items []int
	operator string
	operand	int
	testValue int
	onTrue int
	onFalse int
	inspected int
}

var modulo_value = 1

func playRound(i int, moneky []monkeyStruct, worryModifier int, modulo_value int ) {
	worryLevel, nextMonkey := 0, 0

	for _, item := range moneky[i].items {

		switch moneky[i].operator {
		case "+":
			worryLevel = (item + moneky[i].operand) / worryModifier
		case "-":
			worryLevel = (item - moneky[i].operand) / worryModifier
			
		case "*":
			worryLevel = (item * moneky[i].operand) / worryModifier
			
		case "/":
			worryLevel = (item / moneky[i].operand) / worryModifier
			
		case "square":
			worryLevel = (item * item) / worryModifier
		}

		// credit where it's due I was totally stuck on day 11 part 2
		// Jon / plumpmonkey mentioned modulo and that all the testVaules are primes
		// <insert d'oh here>

		if modulo_value > 0{
			worryLevel = worryLevel % modulo_value 
		}

		if worryLevel % moneky[i].testValue == 0{
			nextMonkey = moneky[i].onTrue
		} else {
			nextMonkey = moneky[i].onFalse
		}

		moneky[nextMonkey].items = append(moneky[nextMonkey].items, worryLevel)

		moneky[i].inspected++
	}

	moneky[i].items = moneky[i].items[:0]
}

func createMonkey(monkeyData []string) monkeyStruct {

	var monkey monkeyStruct

	for _, value := range monkeyData {

		if strings.Contains(value, "Starting items"){
			re := regexp.MustCompile(`\d.`)
			submatchall := re.FindAllString(value, -1)
			for _, element := range submatchall {
				item, _ := strconv.Atoi(element)
				monkey.items = append(monkey.items, item)
			}
		}

		if strings.Contains(value, "Operation"){
			splt_str_1 := strings.Split(value, "=")
			splt_str_2 := strings.Split(splt_str_1[1], " ")

			// splt_str_2[0] == " "
			if splt_str_2[3] == "old" {
				monkey.operator = "square"
			} else {
				monkey.operator = splt_str_2[2]
				monkey.operand, _ = strconv.Atoi(splt_str_2[3])
			}
			
		}

		if strings.Contains(value, "Test"){
			re := regexp.MustCompile(`\d.*`)
			submatchall := re.FindAllString(value, -1)
			for _, element := range submatchall {
				monkey.testValue, _ = strconv.Atoi(element)
			}
			
		}

		if strings.Contains(value, "true"){
			re := regexp.MustCompile(`\d.*`)
			submatchall := re.FindAllString(value, -1)
			for _, element := range submatchall {
				monkey.onTrue, _ = strconv.Atoi(element)
			}
			
		}

		if strings.Contains(value, "false"){
			re := regexp.MustCompile(`\d.*`)
			submatchall := re.FindAllString(value, -1)
			for _, element := range submatchall {
				monkey.onFalse, _ = strconv.Atoi(element)
			}
			
		}

	}

	return monkey
}

func ReadFile(r string) []monkeyStruct{

	var temp_arry []string
	var monkey_arry [] monkeyStruct

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if ! (scanner.Text() == ""){
			temp_arry = append(temp_arry, scanner.Text())
		} else {
			monkey := createMonkey(temp_arry)
			temp_arry = nil
			monkey_arry = append(monkey_arry, monkey)

		}
	}

	// Because we call this twice we need to make sure we don't re-run the modulo value twice
	if modulo_value == 1{
		for _,monkey :=range monkey_arry{
			modulo_value *= monkey.testValue
		}
	}

	return monkey_arry
}


func partOne(puzzle_arry []monkeyStruct) int {

	for round := 0; round < 20; round++ {
		for idx, _ := range puzzle_arry {
			playRound(idx, puzzle_arry, 3, 0)
		}
	}

	max_1 := puzzle_arry[0].inspected
	max_2 := puzzle_arry[1].inspected
	for _, value := range puzzle_arry {
		if value.inspected > max_1 && value.inspected != max_2 {
			max_1 = value.inspected
		}

		if value.inspected > max_2 && value.inspected != max_1 {
			max_2 = value.inspected
		}
	}

	return max_1 * max_2
}

// We need to start from scratch - SO just copy + pasta the function
func partTwo(puzzle_arry []monkeyStruct) int {

	fmt.Println(modulo_value)

	for round := 0; round < 10000; round++ {
		for idx, _ := range puzzle_arry {
			playRound(idx, puzzle_arry, 1, modulo_value)
		}
	}

	max_1 := puzzle_arry[0].inspected
	max_2 := puzzle_arry[1].inspected
	for _, value := range puzzle_arry {
		if value.inspected > max_1 && value.inspected != max_2 {
			max_1 = value.inspected
		}

		if value.inspected > max_2 && value.inspected != max_1 {
			max_2 = value.inspected
		}
	}

	return max_1 * max_2
}


func main() {

	puzzle_arry_one := ReadFile("puzzle_11.txt")
	
	pt1_answer :=  partOne(puzzle_arry_one)

	fmt.Printf("Part 1 answer is: %d\n", pt1_answer)

	puzzle_arry_two := ReadFile("puzzle_11.txt")

	pt2_answer :=  partTwo(puzzle_arry_two)

	fmt.Printf("Part 2 answer is: %d\n", pt2_answer)

}



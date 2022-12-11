package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	//"math"
	//"regexp"
	//"sort"
	//"encoding/csv"
)

var signalStrengths = make(map[int]int)

var crtScreen = make(map[int][]string)

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

func checkSignalStrength(cycle_count int, x_value int) {

	signal_strength := 0

	switch cycle_count{
	case 20, 60, 100, 140, 180, 220:
		signal_strength += cycle_count * x_value
		// Use of struct so we have the correct / latest value in case of an overwrite
		// (before v after x_value update)
		signalStrengths[cycle_count] = signal_strength
	}

}

func drawPixels(cycle_count int, x_value int) {

	switch cycle_count%40{
		case x_value,x_value+1,x_value+2:
			crtScreen[(cycle_count/40)][(cycle_count%40)] = "#"
	}

}

func partOneandTwo(puzzle_arry []string) int {

	x_value := 1
	cycle_count := 1

	cycle_sums := 0

	for row := 6; row >= 0; row --{
		for idx := 39; idx >= 0; idx--{
			crtScreen[row] = append(crtScreen[row], ".")
		}
	}

	for _, value := range puzzle_arry{

		splt_str := strings.Split(value, " ")

		// It will be addx
		if len(splt_str) > 1{
			
			add_x_value, _ := strconv.Atoi(splt_str[1])
			
			// Need multiple checks to catch cycle_count at the right points

			for i := 1; i >= 0; i -- {
				cycle_count += 1

				drawPixels((cycle_count-1),x_value)
				checkSignalStrength(cycle_count,x_value)
			}

			x_value += add_x_value

			checkSignalStrength(cycle_count,x_value)

		} else {

			cycle_count += 1

			checkSignalStrength(cycle_count,x_value)
			drawPixels((cycle_count-1),x_value)

		}

		


	}

	for _, value := range signalStrengths{

		cycle_sums += value
	}

	// Bloody golang _,value := range kept printing this in random idx order which messed with me for
	// FAR longer than it should have
	for i:=0; i < len(crtScreen)-1; i++ {
		fmt.Println(crtScreen[i])
	}

	return cycle_sums
}


func main() {

	puzzle_arry := ReadFile("puzzle_10.txt")

	
	pt1_answer :=  partOneandTwo(puzzle_arry)

	fmt.Printf("Part 1 answer is: %d\n", pt1_answer)


}



package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	//"regexp"
	//"sort"
	//"encoding/csv"
)

type point struct {
	idx,row int
}

var tailLocations = make(map[point]bool)


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

func adjustTail(tail point, head point) point{
	
	newTail := tail
	
	idx_diff := (head.idx-tail.idx)
	row_diff := (head.row-tail.row)

	//Difference Case maping
	// (+/-)2, (+/-)1 = Two idx, One row
	// (+/-)1, (+/-)2 = One idx, Two rows 
	//      0, (+/-)2 = Centre, Two rows     

	// Up and Down - Using row_diff >= on it's own not working - Need to use idx and row
	switch (point{idx_diff,row_diff}){
		case point{2,1},point{1,2}, point{0,2}, point{-1,2}, point{-2,1}:
			newTail.row += 1
		
		case point{2,-1}, point{1,-2}, point{0,-2}, point{-1,-2}, point{-2,-1} :
			newTail.row -= 1
	}

	//Difference Case maping
	// (+/-)2, (+/-)1 = Two idx, One row 
	// (+/-)2,      0 = Two idx, Same row
	// (+/-)1, (+/-)2 = One idx, Two rows

	// Left and Right
	switch (point{idx_diff,row_diff}){
		case point{2,1}, point{2,0}, point{2,-1}, point{1,2}, point{1,-2}:
			newTail.idx += 1

		case point{-2,-1}, point{-2,0}, point{-2,1}, point{-1,2},point{-1,-2}:
			newTail.idx -= 1
	}

	return newTail
}

func partOne(puzzle_arry []string) int {

	// originally used a make(map[int][]string]) to map the locations as per example
	// Many out of bounds errors later switch to track idx and row itself

	head := point{0,0}
	tail := point{0,0}

	tailLocations[tail]=true

	for _, value := range puzzle_arry{

		splt_str := strings.Split(value, " ")
		
		move_dir := splt_str[0]
		move_count, _ := strconv.Atoi(splt_str[1])

		switch move_dir{
			case "R":
				for ; move_count >=1 ; move_count --{
					head.idx += 1
					tail = adjustTail(tail, head)
					tailLocations[tail]=true
				}
			case "L":
				for ; move_count >=1 ; move_count --{
					head.idx -= 1
					tail = adjustTail(tail, head)
					tailLocations[tail]=true
				}
			case "U":

				for ; move_count >=1 ; move_count --{
					head.row += 1
					tail = adjustTail(tail, head)
					tailLocations[tail]=true
				}
			case "D":

				for ; move_count >=1 ; move_count --{
					head.row -= 1
					tail = adjustTail(tail, head)
					tailLocations[tail]=true
				}
		}

	}

	return len(tailLocations)
}

// func partTwo() int {

// 	sum := 0

// 	// golang apparently doesn't like looping in idx order
// 	for i := 0; i < len(tree_map); i ++{
// 		temp := scenicScore(tree_map[i], i)
		
// 		if temp > sum{
// 			sum = temp
// 		}
// 	}
	
// 	return sum
// }


func main() {

	puzzle_arry := ReadFile("puzzle_9.txt")

	
	pt1_answer :=  partOne(puzzle_arry)


	fmt.Printf("Part 1 answer is: %d\n", pt1_answer)
		

}



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

var tree_map = make(map[int][]int)

func reverseArryOrder( arry []int) []int{

	var return_arry []int

	for i := len(arry)-1; i >= 0; i-- {
		return_arry = append(return_arry, arry[i])
	}

	return return_arry

}

func highestTreeInRow(arry []int, tree int) bool{

	result := true

	for _, value := range arry {
		if value >= tree {
			return false
		}
	}

	return result
}

func treesThatCanBeSeen(arry []int, tree int) int{

	result := 0

	// golang apparently doesn't like looping in idx order
	for i := 0; i < len(arry); i ++{
		if arry[i] >= tree {
			// Need to include the tree that is blocking our view
			result += 1
			break;
		} else {
			result += 1
		}
	}

	return result
}

func visibeTrees(row []int, idx int) int{

	number_of_visible_trees := 0
	
	// golang apparently doesn't like looping in idx order
	for i := 0; i < len(row); i ++{
		j := idx +1
		k := idx -1
		left_row := row[:i]
		right_row := row[i+1:]

		var col_below [] int

		var col_above [] int

		for ; j < len(tree_map); j ++{
			col_below = append(col_below, tree_map[j][i])
		}

		for ; k >= 0; k-- {
			col_above = append(col_above, tree_map[k][i])
		}

		treeVisible := false

		if len(left_row) > 0{
			treeVisible = highestTreeInRow(left_row, row[i])

			// Down the rabbit hole we `go` !
			if treeVisible{
				number_of_visible_trees += 1
			} else {
				if len(right_row) > 0{
					treeVisible = highestTreeInRow(right_row, row[i])
		
					if treeVisible{
						number_of_visible_trees += 1

					} else {
						if len(col_below) > 0{
							treeVisible = highestTreeInRow(col_below, row[i])

							if treeVisible{
								number_of_visible_trees += 1

							} else {

								if len(col_above) > 0{
									treeVisible = highestTreeInRow(col_above, row[i])

									if treeVisible{
										number_of_visible_trees += 1
									}

								} else {
									number_of_visible_trees += 1
								}
							}

						} else {
							number_of_visible_trees += 1
						}
					}
				}else {
					number_of_visible_trees += 1
				}
			}
		} else {
			number_of_visible_trees += 1
		}
	}

	return number_of_visible_trees
}


func scenicScore(row []int, idx int) int{

	highest_row_score := 0

	// golang apparently doesn't like looping in idx order
	for i := 0; i < len(row); i ++{
		
		var score_array []int

		j := idx +1
		k := idx -1
		temp_left_row := row[:i]

		// we see our trees to the left in reverse order
		left_row := reverseArryOrder(temp_left_row)

		right_row := row[i+1:]

		var col_below [] int

		var col_above [] int

		for ; j < len(tree_map); j ++{
			col_below = append(col_below, tree_map[j][i])
		}

		for ; k >= 0; k-- {
			col_above = append(col_above, tree_map[k][i])
		}

		// This could likely be a switch case or loop through nested array of arrays

		if len(left_row) > 0{
			//fmt.Println("Left")
			score_array = append(score_array, treesThatCanBeSeen(left_row, row[i]))
		} else {
			score_array = append(score_array, 0)
		}

		if len(right_row) > 0{
			//fmt.Println("Right")
			score_array = append(score_array, treesThatCanBeSeen(right_row, row[i]))
		
		} else {
			score_array = append(score_array, 0)
		}

		if len(col_below) > 0{
			//fmt.Println("Below")
			score_array = append(score_array, treesThatCanBeSeen(col_below, row[i]))
		
		} else {
			score_array = append(score_array, 0)
		}

		if len(col_above) > 0{
			//fmt.Println("Above")
			score_array = append(score_array, treesThatCanBeSeen(col_above, row[i]))
		
		} else {
			score_array = append(score_array, 0)
		}

		number_of_visible_trees := score_array[0]

		for l := 1; l < len(score_array); l ++{
			number_of_visible_trees *= score_array[l]
		}

		if number_of_visible_trees > highest_row_score{
			highest_row_score = number_of_visible_trees
		}

	}

	return highest_row_score
}

// ints were too large to convert without breaking into arry first
func ReadLargeInts(r string){
	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idx := 0
	
	for scanner.Scan() {
		
		for _, value := range strings.Split(scanner.Text(), ""){

			x, _ := strconv.Atoi(string(value))
			if err != nil {
				fmt.Println(err)
			}

			tree_map[idx] = append(tree_map[idx], x)
		}

		idx += 1	
	}
	 
}

func partOne() int {

	sum := 0

	// golang apparently doesn't like looping in idx order
	for i := 0; i < len(tree_map); i ++{
		sum += visibeTrees(tree_map[i], i)
	}
	
	return sum
}

func partTwo() int {

	sum := 0

	// golang apparently doesn't like looping in idx order
	for i := 0; i < len(tree_map); i ++{
		temp := scenicScore(tree_map[i], i)
		
		if temp > sum{
			sum = temp
		}
	}
	
	return sum
}


func main() {

	ReadLargeInts("puzzle_8.txt")

	fmt.Printf("Part 1 answer is: %d\n", partOne())

	fmt.Printf("Part 2 answer is: %d\n", partTwo())
		

}



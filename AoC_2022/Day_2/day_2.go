package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
)

var score_map = make(map[string]int)


func ReadFile(r string) (int, int) {
	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//var result []int
	var pt1_score, pt2_score, column_2_to_lose, column_2_to_win int
	for scanner.Scan() {
		str_arry := strings.Split(scanner.Text(), " ")
		opponent_move := score_map[str_arry[0]] 
		column_2 := score_map[str_arry[1]] // pt1 our_move & pt2 outcome

		// If equal we know it's draw
		if opponent_move == column_2 {
			pt1_score += 3
			
		}else {
			// We only need to worry about winning scenarios
			switch column_2 {
			case 1: // rock
				if opponent_move == 3 { // scissors
					pt1_score += 6
				}
			case 2: // paper
				if opponent_move == 1 { // rock
					pt1_score += 6
				}
			case 3: // scissors
				if opponent_move == 2 { // paper
					pt1_score += 6
				}
			}
		}
		// constant
		pt1_score += column_2

		switch column_2 {

			// lose
			case 1:
				switch opponent_move{
					case 1: // rock
						column_2_to_lose = 3 // scissors
					case 2: // paper
						column_2_to_lose = 1 // rock
					case 3: // scissors
						column_2_to_lose = 2 // paper
				}
				pt2_score += column_2_to_lose

			// draw
			case 2:
				pt2_score += 3
				pt2_score += opponent_move

			// win
			case 3:
				switch opponent_move{
					case 1:
						column_2_to_win = 2
					case 2:
						column_2_to_win = 3
					case 3:
						column_2_to_win = 1
				}
				pt2_score += 6
				pt2_score += column_2_to_win
		}

	}

	return pt1_score, pt2_score
}

// I miss arry = arry.remove()
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}


func main() {

	score_map["A"] = 1 // rock & lose
	score_map["B"] = 2 // paper & draw 
	score_map["C"] = 3 // scissors & win
	score_map["X"] = 1 
	score_map["Y"] = 2 
	score_map["Z"] = 3 

	pt1_score, pt2_score := ReadFile("puzzle_2.txt")

	fmt.Printf( "Puzzle 1 answer is: %d\n", pt1_score);

	fmt.Printf( "Puzzle 2 answer is: %d\n", pt2_score)
	

}
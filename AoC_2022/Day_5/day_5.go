package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"regexp"
	//"sort"
	//"encoding/csv"
)

var stack_map = make(map[int][]string)

// I miss arry = arry.remove()
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}


func partOne(stack_arry []string, move_arry []string) string {

	// flip it and reverse it ya'll
	for i := len(stack_arry)-1; i >= 0; i-- {

		count := 0

		// There might be a neater way to do this - But I clearly did it this way
		for j := 0; j <= len(stack_arry[i]); j += 4{

			count += 1

			end := j+4

			if end > len(stack_arry[i]){
				end = len(stack_arry[i])
			}
			
			temp_str := stack_arry[i][j:end]

			re := regexp.MustCompile(`[a-zA-Z0-9]`)

			submatchall := re.FindAllString(temp_str, -1)

			if len(submatchall) > 0{
				stack_map[count] = append(stack_map[count],submatchall[0])
			}			
		}
		
	}

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	for _, value := range move_arry {
		var moves []string
		submatchall := re.FindAllString(value, -1)
		for _, element := range submatchall {
			moves = append(moves, element)
		}

		crates_to_move, _ := strconv.Atoi(moves[0])

		stack_to_move_from, _ := strconv.Atoi(moves[1])

		stack_to_move_to, _ := strconv.Atoi(moves[2])		
		
		for i := crates_to_move; i > 0; i-- {

			idx_value := len(stack_map[stack_to_move_from]) -1

			crate_to_move := stack_map[stack_to_move_from][idx_value]

			stack_map[stack_to_move_to] = append(stack_map[stack_to_move_to], crate_to_move)

			stack_map[stack_to_move_from] = RemoveIndex(stack_map[stack_to_move_from], idx_value)

		}

	}
	var result []string
	for i := 1; i <= len(stack_map); i ++{
		result = append(result, stack_map[i][len(stack_map[i])-1])
	}

	return strings.Join(result, "")
}

func partTwo(stack_arry []string, move_arry []string) string {

	// flip it and reverse it ya'll
	for i := len(stack_arry)-1; i >= 0; i-- {

		count := 0

		// There might be a neater way to do this - But I clearly did it this way
		for j := 0; j <= len(stack_arry[i]); j += 4{

			count += 1

			end := j+4

			if end > len(stack_arry[i]){
				end = len(stack_arry[i])
			}
			
			temp_str := stack_arry[i][j:end]

			re := regexp.MustCompile(`[a-zA-Z0-9]`)

			submatchall := re.FindAllString(temp_str, -1)

			if len(submatchall) > 0{
				stack_map[count] = append(stack_map[count],submatchall[0])
			}			
		}
		
	}

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	for _, value := range move_arry {
		var moves []string
		submatchall := re.FindAllString(value, -1)
		for _, element := range submatchall {
			moves = append(moves, element)
		}

		crates_to_move, _ := strconv.Atoi(moves[0])

		stack_to_move_from, _ := strconv.Atoi(moves[1])

		stack_to_move_to, _ := strconv.Atoi(moves[2])	
		
		// We probably don't need a whole "new" func for this - tidy on revist
		var crate_holder_9001 []string
		
		for i := crates_to_move; i > 0; i-- {

			idx_value := len(stack_map[stack_to_move_from]) -1

			crate_to_move := stack_map[stack_to_move_from][idx_value]

			crate_holder_9001 = append(crate_holder_9001, crate_to_move)

			stack_map[stack_to_move_from] = RemoveIndex(stack_map[stack_to_move_from], idx_value)

		}

		for i := len(crate_holder_9001)-1; i >= 0; i-- {
			stack_map[stack_to_move_to] = append(stack_map[stack_to_move_to], crate_holder_9001[i])
		}

	}
	var result []string
	for i := 1; i <= len(stack_map); i ++{
		result = append(result, stack_map[i][len(stack_map[i])-1])
	}

	return strings.Join(result, "")

}


func ReadFile(r string) ([]string, []string){

	var stack_arry []string
	var move_arry []string

	file, err := os.Open(r)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "move") {
			move_arry = append(move_arry, scanner.Text())
		} else {
			if scanner.Text() != "" {
				stack_arry = append(stack_arry, scanner.Text())
			}	
		}
		
	}

	return stack_arry, move_arry
}


func main() {

	stack_arry, move_arry := ReadFile("puzzle_5.txt")
	
	pt1_answer :=  partOne(stack_arry,move_arry)

	pt2_answer := partTwo(stack_arry,move_arry)

	fmt.Printf( "Puzzle 1 answer is: %s\n", pt1_answer);

	fmt.Printf( "Puzzle 2 answer is: %s\n", pt2_answer);
	

}
package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	//"strings"
	//"strconv"
	//"regexp"
	//"sort"
	//"encoding/csv"
)

// I miss arry = arry.remove()
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func areValuesUnique(check []string, check_len int) bool{
	// Apparently the set datatype is not a thing in golang
	marker_set := make(map[string]bool)

	var verdict bool

	for _, value := range check {
        marker_set[value] = true
    }

	if len(marker_set) == check_len{
		verdict = true
	} else {
		verdict = false
	}

	return verdict

}

func partOneAndTwo(puzzle_arry []string) {

	marker_verdict := false

	message_verdict := false

	var marker_arry []string

	var message_arry []string

	for idx, value := range puzzle_arry[0]{

		char := string(value)

		if len(marker_arry) == 4{
			marker_arry = RemoveIndex(marker_arry, 0)
		}

		if len(message_arry) == 14{
			message_arry = RemoveIndex(message_arry, 0)
		}

		marker_arry = append(marker_arry, char)

		message_arry = append(message_arry, char)

		
		if ! marker_verdict {
			marker_verdict = areValuesUnique(marker_arry, 4)

			if marker_verdict {
				fmt.Printf("Puzzle 1 answer is: %d\n", idx + 1);
			}
		}

		if ! message_verdict {
			message_verdict = areValuesUnique(message_arry, 14)

			if message_verdict {
				fmt.Printf("Puzzle 2 answer is: %d\n", idx + 1);
			}
		}

	}
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

	puzzle_arry := ReadFile("puzzle_6.txt")
	
	partOneAndTwo(puzzle_arry)
	

}
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

// In Linux everythig is a file...
type fileStruct struct {
	name string
	size int
	isFile bool
	subDirs map[string]*fileStruct
	rootDir *fileStruct
}

// Let's get recursive
func dirSize(root fileStruct)int{

	size:=0

	if root.isFile {
		return root.size
	} else {
		for _, d := range root.subDirs{
			size += dirSize(*d)
		}
	}

	return size
	
}

func partOneandTwo(puzzle_arry []string) {

	var current_dir *fileStruct

	dirs := []*fileStruct{}

	for _, value := range puzzle_arry {

		cmds := strings.Split(value, " ")
		
		// Need the spaces because otherwise we get thrown by filenames with valuse in them (like 'cd')
		if strings.Contains(value, "cd "){
			// 0 == $, 1 == cd
			switch cmds[2]{
				case "..":
					current_dir = current_dir.rootDir
				case "/":
					current_dir = &fileStruct{"/", 0, false, make(map[string]*fileStruct), nil}
					dirs = append(dirs, current_dir)
				default:
					current_dir = current_dir.subDirs[cmds[2]]
			}		


		}else if strings.Contains(value, "dir "){
			
			current_dir.subDirs[cmds[1]] = &fileStruct{cmds[1], 0, false, make(map[string]*fileStruct), current_dir}
			dirs = append(dirs, current_dir.subDirs[cmds[1]])

		}else if ! (strings.Contains(value, "$ ")){
			size, _ := strconv.Atoi(cmds[0])
			current_dir.subDirs[cmds[1]] = &fileStruct{cmds[1], size, true, nil, current_dir}
		}
	}

	sum := 0

	// 30000000 - (used space)
	space_req := 30000000 - (70000000 - dirSize(*dirs[0]))

	var candidate_dir_sizes []int

	for _, dir := range dirs{
		
		size := dirSize(*dir)

		if size <= 100000 {
			sum += size
		}

		if size > space_req {
			candidate_dir_sizes = append(candidate_dir_sizes, size)
		}
	}

	fmt.Printf("Part 1 answer is: %d\n", sum) 

	min := candidate_dir_sizes[0]
	for _, value := range candidate_dir_sizes {
		if value < min {
			min = value
		}
	}

	fmt.Printf("Part 2 answer is: %d\n", min) 

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

	puzzle_arry := ReadFile("puzzle_7.txt")
	
	partOneandTwo(puzzle_arry)
	

}
package lem_in

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	ants      int
	links     [][]string
	startNode string
	endNode   string
	allRoms   []string
)

func ReadFile() (int, [][]string, string, string) {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <file.txt>")
		os.Exit(1)
	}

	nodesFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer nodesFile.Close()
	start := false
	end := false

	scanner := bufio.NewScanner(nodesFile)

	for scanner.Scan() {
		switch checkline(scanner.Text())[0] {
		case "start":
			start = true
		case "end":
			end = true
		case "node":
			id := checkline(scanner.Text())[1]
			if (slices.Index(allRoms, id)) != -1 {
				fmt.Println("ERROR: invalid data format, some rooms are redeclared")
				os.Exit(1)
			}
			allRoms = append(allRoms, id)
			if start && startNode == "" {
				startNode = id
				start = false
			}
			if start && startNode != "" {
				fmt.Println("ERROR: invalid data format, START room declared more than once")
				os.Exit(1)
			}
			if end && endNode == "" {
				endNode = id
				end = false
			}
			if end && endNode != "" {
				fmt.Println("ERROR: invalid data format, END room declared more than once")
				os.Exit(1)
			}
		case "edge":
			
			from := checkline(scanner.Text())[1]
			to := checkline(scanner.Text())[2]
			if (slices.Index(allRoms, from)) == -1 || (slices.Index(allRoms, to)) == -1{
				fmt.Println("ERROR: invalid data format, some rooms not defined")
				os.Exit(1)
			}
			e := []string{from, to}
			ereversed := []string{to, from}
			if NotIn(links, e) && NotIn(links, ereversed) {
				links = append(links, e)
			} else {
				fmt.Println("ERROR: invalid data format, some edges declared more than once")
				os.Exit(1)
			}

		}
	}
	return ants, links, startNode, endNode
}

func checkline(s string) []string {
	spl := strings.Split(s, " ")
	if len(spl) == 3 {
		x, err2 := strconv.Atoi(spl[1])
		y, err3 := strconv.Atoi(spl[2])
		if err2 != nil || err3 != nil || x < 0 || y < 0 {
			fmt.Println("ERROR: invalid data format, some rooms have invalid coordinates")
			os.Exit(1)
		}
		if validRoom(spl[0]) {
			return []string{"node", spl[0]}
		}

	}
	if len(spl) == 1 && spl[0] == "##start" {
		return []string{"start"}
	}
	if len(spl) == 1 && spl[0] == "##end" {
		return []string{"end"}
	}

	if len(spl) == 1 {
		antsNum, err := strconv.Atoi(spl[0])
		if err != nil || antsNum <= 0 {
			antsNum = 0
		} else {
			ants = antsNum
		}
	}
	spl = nil
	spl = strings.Split(s, "-")
	if len(spl) == 2 {
		if validRoom(spl[0]) && validRoom(spl[1]) {
			return []string{"edge", spl[0], spl[1]}
		}
	}
	spl = nil

	return []string{"nil"}
}

func validRoom(room string) bool {
	if string(room[0]) == "L" || string(room[0]) == "#" {
		fmt.Println("ERROR: invalid data format, some rooms starts with 'L' or '#' ")
		os.Exit(1)
	}
	return true
}

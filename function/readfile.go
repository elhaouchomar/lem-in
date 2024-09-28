package lem_in

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ants      int
	links     [][]string
	startNode string
	endNode   string
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
			if start && startNode == "" {
				startNode = id
				start = false
			}
			if end && endNode == "" {
				endNode = id
				end = false
			}
		case "edge":
			from := checkline(scanner.Text())[1]
			to := checkline(scanner.Text())[2]
			e := []string{from, to}
			if NotIn(links, e) {
				links = append(links, e)
			}
			
		}
	}
	return ants, links, startNode, endNode
}

func checkline(s string) []string {
	spl := strings.Split(s, "-")
	if len(spl) == 2 {
		return []string{"edge", spl[0], spl[1]}
	}
	spl = nil
	spl = strings.Split(s, " ")
	if len(spl) == 3 {
		return []string{"node", spl[0]}
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

	return []string{"nil"}
}

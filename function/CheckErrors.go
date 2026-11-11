package lem_in

import "fmt"

func CheckErrors(ants int, start, end string, l int) string {
	if ants == 0 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		return "error"
	}

	if start == "" {
		fmt.Println("ERROR: invalid data format, no START room found")
		return "error"
	}

	if end == "" {
		fmt.Println("ERROR: invalid data format, no END room found")
		return "error"
	}

	if l == 0 {
		fmt.Println("ERROR: invalid data format, no valid path between START and END room")
		return "error"
	}
	return "nil"
}

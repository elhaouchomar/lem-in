package lem_in

import (
	"fmt"
	"os"
)

func CheckErrors(ants int, start, end string, l int) {
	if ants == 0 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(1)
	}

	if start == "" {
		fmt.Println("ERROR: invalid data format, no START room found")
		os.Exit(1)
	}

	if end == "" {
		fmt.Println("ERROR: invalid data format, no END room found")
		os.Exit(1)
	}

	if l == 0 {
		fmt.Println("ERROR: invalid data format, no valid path between START and END room")
		os.Exit(1)
	}
}

package lem_in

import (
	"fmt"
	"sort"
	"strings"
)

var antMovements = [][]string{} // Stores the sequences of ant movements

func PrintAnts(totalAnts int, paths [][]string) {
	rankedPaths := RankPaths(paths) // Ranked paths after filtering
	uniquePaths := FilterUniquePaths(rankedPaths) // Unique paths for ant movement
	sort.Slice(uniquePaths, func(i, j int) bool {
		return len(uniquePaths[i]) < len(uniquePaths[j]) // Sort paths by length
	})
	assignAntsToPaths(totalAnts, uniquePaths) // Assign ants to paths
	currentStep := 0 // Current output index
	currentAnt := 1 // Current ant number
	allAntsMoved := false // Flag to stop movement
	for !allAntsMoved {
		for pathIndex := range antPathAssignments {
			if antPathAssignments[pathIndex] != 0 {
				moveAnts(currentAnt, uniquePaths[pathIndex], currentStep) // Move ants along paths
				antPathAssignments[pathIndex]--
				currentAnt++
				if currentAnt > totalAnts {
					allAntsMoved = true // Stop if all ants are moved
					break
				}
			}
		}
		currentStep++ // Increment output index
	}

	printAntMovements(antMovements) // Print final output
	fmt.Println(len(antMovements)) // Print the length of output
}

func printAntMovements(movements [][]string) {
	for i := 0; i < len(movements); i++ {
		fmt.Println(strings.Join(movements[i], " ")) // Print each movement sequence
	}
}

func moveAnts(antNumber int, path []string, stepIndex int) {
	for i := 0; i < len(path); i++ {
		movement := fmt.Sprintf("L%v-%v", antNumber, path[i]) // Format ant movement
		if stepIndex < len(antMovements) {
			antMovements[stepIndex] = append(antMovements[stepIndex], movement) // Append movement to existing output
			stepIndex++
		} else {
			antMovements = append(antMovements, []string{movement}) // Create new output entry
			stepIndex++
		}
	}
}

var antPathAssignments = make(map[int]int) // Map to track ants assigned to paths

func assignAntsToPaths(totalAnts int, lastPath [][]string) {
	tempPaths := [][]string{} // Temporary slice for path processing
	antCount := 1 // Counter for ants
	lastPathProcessed := false // Flag for processing last path
	for i := 1; i < len(lastPath); i++ {
		lengthDifference := len(lastPath[i]) - len(lastPath[i-1]) // Length difference between paths
		tempPaths = append(tempPaths, lastPath[i-1]) // Append previous path
		for antCount <= lengthDifference {
			for k := 0; k < len(tempPaths); k++ {
				antPathAssignments[k]++ // Increment ant assignment
				antCount++
			}
		}
	}
	for antCount <= totalAnts {
		if !lastPathProcessed {
			tempPaths = append(tempPaths, lastPath[len(lastPath)-1]) // Process last path
			lastPathProcessed = true
		}
		for k := 0; k < len(tempPaths); k++ {
			antPathAssignments[k]++ // Increment ant assignment
			antCount++
			if antCount == totalAnts+1 {
				break // Stop if all ants are assigned
			}
		}
	}
}

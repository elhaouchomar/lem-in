package lem_in

import (
	"reflect"
	"slices"
)

// Paths stores all discovered paths.
var Paths [][]string

// DepthFirstSearch performs DFS to find all paths to the end node.
func DepthFirstSearch(endNode string, graph map[string][]string, currentPath []string) {
	currentNode := currentPath[len(currentPath)-1]
	if currentNode == endNode {
		currentPath = currentPath[1:]
		Paths = append(Paths, currentPath)
		return
	}
	for _, neighbor := range graph[currentNode] {
		if slices.Index(currentPath, neighbor) == -1 {
			newPath := make([]string, len(currentPath))
			copy(newPath, currentPath)
			newPath = append(newPath, neighbor)
			DepthFirstSearch(endNode, graph, newPath)
		}
		if endNode == neighbor {
			return
		}
	}
}

// IsPathUnique checks if a path is not already in the list of paths.
func IsPathUnique(allPaths [][]string, newPath []string) bool {
	for _, existingPath := range allPaths {
		if reflect.DeepEqual(existingPath, newPath) {
			return false
		}
	}
	return true
}

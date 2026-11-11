package lem_in

// FilterUniquePaths filters out duplicate paths.
func FilterUniquePaths(paths [][]string) [][]string {
	uniquePaths := [][]string{paths[0]}
	for i := 1; i < len(paths); i++ {
		isUnique := true
		for j := 0; j < len(uniquePaths); j++ {
			if !arePathsEqual(paths[i], uniquePaths[j]) {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniquePaths = append(uniquePaths, paths[i])
		}
	}
	return uniquePaths
}

// arePathsEqual checks if two paths are equivalent.
func arePathsEqual(pathA, pathB []string) bool {
	if len(pathA) > len(pathB) {
		pathB, pathA = pathA, pathB
	}
	for i := 0; i < len(pathA)-1; i++ {
		if pathA[i] == pathB[i] {
			return false
		}
	}
	j := len(pathB) - 2
	for i := len(pathA) - 2; i >= 0; i-- {
		if pathA[i] == pathB[j] {
			return false
		}
		j--
	}
	return true
}

// RankPaths ranks the paths based on the frequency of nodes.
func RankPaths(paths [][]string) [][]string {
	ranks := calculatePathRanks(paths)
	return sortPathsByRank(paths, ranks)
}

// calculatePathRanks calculates the rank for each path based on node frequency.
func calculatePathRanks(paths [][]string) []int {
	rankings := make([]int, len(paths))
	nodeCount := make(map[string]int)
	for _, path := range paths {
		for _, node := range path {
			nodeCount[node]++
		}
	}
	for i, path := range paths {
		total := 0
		for _, node := range path {
			total += nodeCount[node]
		}
		rankings[i] = total
	}
	return rankings
}

// sortPathsByRank sorts paths based on their calculated ranks.
func sortPathsByRank(paths [][]string, ranks []int) [][]string {
	for i := 0; i < len(ranks); i++ {
		for j := i + 1; j < len(ranks); j++ {
			if ranks[i] > ranks[j] {
				ranks[i], ranks[j] = ranks[j], ranks[i]
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	return paths
}

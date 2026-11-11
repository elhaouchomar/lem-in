package main

import (
	lem_in "lem_in/function"
)

func main() {
	ants, links, start, end := lem_in.ReadFile()
	graph := lem_in.BuildGraph(links)
	lem_in.DepthFirstSearch(end, graph, []string{start})

	errorMsg := lem_in.CheckErrors(ants, start, end, len(lem_in.Paths))

	if errorMsg == "error" {
		return
	}

	lem_in.PrintAnts(ants, lem_in.Paths)
}

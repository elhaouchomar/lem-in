package lem_in

// BuildGraph converts the edge list into an adjacency list representation.
func BuildGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

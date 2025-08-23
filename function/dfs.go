package lem_in

import (
	"reflect"
	"slices"
)

var Paths [][]string

func Defs(end string, graph map[string][]string, path []string) {
	n := path[len(path)-1]
	if n == end {
		path = path[1:]
		if NotIn(Paths, path) {
			Paths = append(Paths, path)
		}
		return
	}
	for i := 0; i < len(graph[n]); i++ {
		if slices.Index(path, graph[n][i]) == -1 {
			new := make([]string, len(path))
			copy(new, path)
			new = append(new, graph[n][i])
			Defs(end, graph, new)
		}
		if end == graph[n][i] {
			return
		}
	}
}

func NotIn(paths [][]string, path []string) bool {
	for i := 0; i < len(paths); i++ {
		if reflect.DeepEqual(paths[i], path) {
			return false
		}
	}
	return true
}

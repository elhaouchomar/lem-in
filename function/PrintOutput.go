package lem_in

import (
	"fmt"
	"strings"
)

var output = [][]string{}

func PrintAnts(ants int, paths [][]string) {
	arr := Choose(paths)
	paths = ChosePath(arr)
	affectPaths(ants, paths)
	j := 0
	an := 1
	stop := false
	for !stop {
		for p := range antsPaths {
			if antsPaths[p] != 0 {
				moveAnts(an, paths[p], j)
				antsPaths[p]--
				an++
				if an > ants {
					stop = true
					break
				}
			}
		}
		j++

	}

	printoutput(output)
}

func printoutput(output [][]string) {
	for i := 0; i < len(output); i++ {
		fmt.Println(strings.Join(output[i], " "))
	}
}

func moveAnts(ant int, path []string, j int) {
	for i := 0; i < len(path); i++ {
		toPrint := fmt.Sprintf("L%v-%v", ant, path[i])
		if j < len(output) {
			output[j] = append(output[j], toPrint)
			j++
		} else {
			output = append(output, []string{toPrint})
			j++
		}
	}
}

var antsPaths = make(map[int]int)

func affectPaths(ants int, lastPath [][]string) {
	ar := [][]string{}
	a := 1
	for i := 1; i < len(lastPath); i++ {
		diff := len(lastPath[i]) - len(lastPath[i-1])
		ar = append(ar, lastPath[i-1])
		for a <= diff {
			for k := 0; k < len(ar); k++ {
				antsPaths[k]++
				a++
			}
		}
	}

	ar = append(ar, lastPath[len(lastPath)-1])

	for a <= ants {
		for k := 0; k < len(ar); k++ {
			antsPaths[k]++
			a++
			if a == ants+1 {
				break
			}
		}
	}
}

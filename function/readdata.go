package lem_in

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadData(s string) (int, [][]string, string, string) {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <file.txt>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	datas := string(data)
	datas = strings.ReplaceAll(datas, "\n", " ")
	sl_datas := strings.Split(datas, " ")
	nmbAnts, err := strconv.Atoi(sl_datas[0])
	if err != nil {
		nmbAnts = 0
	}
	var AllPath []string
	var Path [][]string
	var start, end string
	for i, datas := range sl_datas {
		if sl_datas[i] == "##start" {
			start = sl_datas[i+1]
		} else if sl_datas[i] == "##end" {
			end = sl_datas[i+1]
		}
		for j, c := range datas {
			if c == '-' {
				strat := datas[:j]
				AllPath = append(AllPath, string(strat))

				if j+1 < len(datas) {
					target := datas[j+1:]
					AllPath = append(AllPath, string(target))
				}

				Path = append(Path, AllPath)
				AllPath = nil
				break
			}
		}
	}
	return nmbAnts, Path, start, end
}

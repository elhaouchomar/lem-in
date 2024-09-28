package lem_in

func ConvertToGraph(ar [][]string) map[string][]string {
	m := make(map[string][]string)
	for i := 0; i < len(ar); i++ {
		m[ar[i][0]] = append(m[ar[i][0]], ar[i][1])
		m[ar[i][1]] = append(m[ar[i][1]], ar[i][0])
	}
	return m
}

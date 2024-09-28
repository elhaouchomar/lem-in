package lem_in

func ChosePath(ar [][]string) [][]string {
	arr := [][]string{ar[0]}
	for i := 1; i < len(ar); i++ {
		isvalid := true
		for j := 0; j < len(arr); j++ {
			if !compare(ar[i], arr[j]) {
				isvalid = false
				break
			}
		}

		if isvalid {
			arr = append(arr, ar[i])
		}
	}
	return arr
}

func compare(a, b []string) bool {
	if len(a) > len(b) {
		b, a = a, b
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] == b[i] {
			return false
		}
	}
	j := len(b) - 2
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] == b[j] {
			return false
		}
		j--
	}
	// fmt.Println(a,b)
	return true
}

func Choose(arr [][]string) [][]string {
	r := rating(arr)
	arr = sortt(arr, r)
	return arr
}

func rating(arr [][]string) []int {
	r := make([]int, len(arr))
	m := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			m[arr[i][j]]++
		}
	}
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr[i]); j++ {
			sum += m[arr[i][j]]
		}
		r[i] = sum
	}
	return r
}

func sortt(arr [][]string, r []int) [][]string {
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

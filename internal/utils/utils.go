package utils

import "fmt"

func SplitSliceIntoChunks(slice []int, chunkSize int) [][]int {
	if slice == nil || chunkSize == 0 || chunkSize < 0 {
		return nil
	}

	l := len(slice)

	if chunkSize > l {
		return [][]int{slice[:]}
	}

	capacity := l / chunkSize
	if l%chunkSize > 0 {
		capacity += 1
	}
	res := make([][]int, 0, capacity)

	i := 0
	for i < l {
		end := i + chunkSize
		if end > l {
			end = l
		}
		res = append(res, slice[i:end])
		i = end
	}

	return res
}

func InvertedMap(m map[int]int) map[int]int {
	res := make(map[int]int, len(m))
	for k, v := range m {
		if _, found := res[v]; found {
			panic(fmt.Sprintf("duplicate value %d", v))
		}
		res[v] = k
	}
	return res
}

func Filter(s []int) []int {
	var list = [...]int{1, 2, 3, 4}
	contains := func(v int) bool {
		for _, v2 := range list {
			if v2 == v {
				return true
			}
		}
		return false
	}

	r := []int{}
	for _, v := range s {
		if !contains(v) {
			r = append(r, v)
		}
	}

	return r
}

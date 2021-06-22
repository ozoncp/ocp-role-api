package utils

import (
	"fmt"

	"github.com/ozoncp/ocp-role-api/internal/model"
)

func SplitToBulks(roles []*model.Role, batchSize uint) [][]*model.Role {
	if roles == nil || batchSize == 0 {
		return nil
	}

	l := uint(len(roles))

	if batchSize > l {
		return [][]*model.Role{roles[:]}
	}

	capacity := l / batchSize
	if l%batchSize > 0 {
		capacity += 1
	}
	res := make([][]*model.Role, 0, capacity)

	i := uint(0)
	for i < l {
		end := i + batchSize
		if end > l {
			end = l
		}
		res = append(res, roles[i:end])
		i = end
	}

	return res
}

func GetServiceRolesMap(roles []model.Role) map[string][]model.Role {
	res := make(map[string][]model.Role, len(roles))
	for _, r := range roles {
		res[r.Service] = append(res[r.Service], r)
	}
	return res
}

func GetRolesMap(roles []model.Role) (map[string]model.Role, error) {
	res := make(map[string]model.Role, len(roles))
	for _, r := range roles {
		if _, found := res[r.Service]; found {
			return nil, fmt.Errorf("not unique key \"%s\"", r.Service)
		}
		res[r.Service] = r
	}
	return res, nil
}

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

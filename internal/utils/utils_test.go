package utils

import (
	"reflect"
	"testing"
)

func TestSplitSliceIntoChunks(t *testing.T) {
	cases := []struct {
		s    []int
		n    int
		want [][]int
	}{
		{
			nil,
			2,
			nil,
		},
		{
			[]int{1, 2},
			0,
			nil,
		},
		{
			[]int{1, 2},
			-1,
			nil,
		},
		{
			[]int{1, 2},
			1,
			[][]int{{1}, {2}},
		},
		{
			[]int{1, 2, 3, 4},
			2,
			[][]int{{1, 2}, {3, 4}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			[]int{1, 2, 3, 4},
			3,
			[][]int{{1, 2, 3}, {4}},
		},
		{
			[]int{1, 2, 3, 4},
			4,
			[][]int{{1, 2, 3, 4}},
		},
		{
			[]int{1, 2, 3, 4},
			5,
			[][]int{{1, 2, 3, 4}},
		},
	}

	for _, c := range cases {
		got := SplitSliceIntoChunks(c.s, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SplitSliceIntoChunks(%v, %v) == %v, want %v", c.s, c.n, got, c.want)
		}
	}
}

func TestInvertedMap(t *testing.T) {
	cases := []struct {
		in, want map[int]int
	}{
		{
			nil,
			nil,
		},
		{
			map[int]int{},
			map[int]int{},
		},
		{
			map[int]int{1: 2},
			map[int]int{2: 1},
		},
		{
			map[int]int{1: 2, 3: 4},
			map[int]int{2: 1, 4: 3},
		},
		{
			map[int]int{1: 2, 3: 4},
			map[int]int{2: 1, 4: 3},
		},
	}

	for _, c := range cases {
		for k1, v1 := range c.in {
			if _, found := c.want[v1]; !found {
				t.Errorf("key %d not found in %v", v1, c.want)
			}

			found := false
			for _, v2 := range c.want {
				if v2 == k1 {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("value %d not found in %v", k1, c.want)

			}
		}
	}

	in := map[int]int{1: 2, 3: 2}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("InvertedMap(%v) did not panic", in)
		}
	}()

	InvertedMap(in)
}

func TestFilter(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for _, c := range cases {
		got := Filter(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Filter(%v) != %v", got, c.want)
		}
	}
}

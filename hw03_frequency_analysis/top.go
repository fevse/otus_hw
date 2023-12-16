package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type keyValue struct {
	Key   string
	Value int
}

func Top10(s string) []string {
	ms := make(map[string]int)
	sl := strings.Fields(s)
	for _, v := range sl {
		ms[v]++
	}
	res := mapToSortSlice(ms)
	return res
}

func mapToSortSlice(m map[string]int) []string {
	skv := make([]keyValue, 0, len(m))
	for k, v := range m {
		skv = append(skv, keyValue{k, v})
	}
	sort.Slice(skv, func(i, j int) bool {
		if skv[i].Value == skv[j].Value {
			return skv[i].Key < skv[j].Key
		}
		return skv[i].Value > skv[j].Value
	})
	res := make([]string, 0, len(skv))
	for _, v := range skv {
		res = append(res, v.Key)
		if len(res) == 10 {
			break
		}
	}
	return res
}

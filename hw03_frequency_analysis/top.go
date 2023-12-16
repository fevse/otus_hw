package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	ms := make(map[string]int)
	sl := strings.Fields(s)
	for _, v := range sl {
		ms[v]++
	}
	res := mapToSortSlice(ms)
	if len(res) > 10 {
		return res[:10]
	}
	return res
}

func mapToSortSlice(m map[string]int) []string {
	type keyValue struct {
		Key   string
		Value int
	}
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
	}
	return res
}

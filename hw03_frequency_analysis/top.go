package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type pair struct {
	key   string
	value int
}

func Top10(t string) []string {
	words := regexp.MustCompile(`\p{L}+(-\p{L}+)*|-{2,}`).FindAllString(strings.ToLower(t), -1)
	counter := make(map[string]int)

	for i := range words {
		counter[words[i]]++
	}
	z := make([]pair, 0, len(counter))

	for k, v := range counter {
		z = append(z, pair{k, v})
	}

	sort.Slice(z, func(i, j int) bool {
		if z[i].value == z[j].value {
			return len(z[i].key) > len(z[j].key)
		}
		return z[i].value > z[j].value
	})

	result := make([]string, 0, len(z))
	for _, pair := range z[:min(10, len(z))] {
		result = append(result, pair.key)
	}

	return result
}

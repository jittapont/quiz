package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupWords(words []string) [][]string {
	wordGroupMapper := make(map[string][]string)
	groups := make([][]string, 0)
	for _, word := range words {
		sortedWord := sortString(word)
		group, ok := wordGroupMapper[sortedWord]
		if !ok {
			wordGroupMapper[sortedWord] = append(make([]string, 0), word)
		} else {
			wordGroupMapper[sortedWord] = append(group, word)
		}
	}
	for _, group := range wordGroupMapper {
		groups = append(groups, group)
	}
	return groups
}

func groupWords2(words []string) [][]string {
	wordGroupMapper := make(map[[26]int][]string)
	groups := make([][]string, 0)
	for _, word := range words {
		arr := [26]int{}
		for _, char := range word {
			arr[int(char) - int('A')] += 1
		}
		group, ok := wordGroupMapper[arr]
		if !ok {
			wordGroupMapper[arr] = append(make([]string, 0), word)
		} else {
			wordGroupMapper[arr] = append(group, word)
		}
	}
	for _, group := range wordGroupMapper {
		groups = append(groups, group)
	}
	return groups
}

func main() {
	words := []string{"VMRCO", "VORCM", "MCRTOX", "ZXTAC", "XZATC", "XMTCOR", "OXVS", "AZTXC", "VXOS", "RZAT", "MRCOTX", "SVXO", "TRAZ", "ZTAR", "MVOCR"}
	groups := groupWords2(words)
	fmt.Println(groups)
}

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

func main() {
	words := []string{"VMRCO", "VORCM", "MCRTOX", "ZXTAC", "XZATC", "XMTCOR", "OXVS", "AZTXC", "VXOS", "RZAT", "MRCOTX", "SVXO", "TRAZ", "ZTAR", "MVOCR"}
	groups := groupWords(words)
	fmt.Println(groups)
}

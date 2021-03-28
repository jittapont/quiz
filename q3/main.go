package main

import (
	"fmt"
	"sort"
)

func main() {
	sampleStrings := []string{"ABC", "112", "AB"}
	for _, sampleString := range sampleStrings {
		fmt.Printf("String: %s\n", sampleString)
		fmt.Printf("Result: %s\n", getPermutations(sampleString))
	}
}

func getPermutations(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}
	current := str[0:1]
	remainingString := str[1:]
	perms := getPermutations(remainingString)
	permutations := make([]string, 0)
	permutationsMapper := make(map[string]bool)
	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			newPermutation := insertAt(i, current, perm)
			if !permutationsMapper[newPermutation] {
				permutationsMapper[newPermutation] = true
				permutations = append(permutations, newPermutation)
			}
		}
	}
	sort.Strings(permutations)
	return permutations
}

func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:]
	return start + char + end
}

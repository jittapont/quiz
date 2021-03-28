package main

import (
	"reflect"
	"testing"
)

func Test_getPermutations(t *testing.T) {
	tests := []struct {
		name string
		word string
		want []string
	}{
		{
			name: "When we provide a word, it should return all permutations of string",
			word: "ABC",
			want: []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"},
		},
		{
			name: "When we provide a word, it should return all permutations of string",
			word: "112",
			want: []string{"112", "121", "211"},
		},
		{
			name: "When we provide a word, it should return all permutations of string",
			word: "AB",
			want: []string{"AB", "BA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPermutations(tt.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPermutations() result mismatch got : %v, want : %v", got, tt.want)
			}
		})
	}
}

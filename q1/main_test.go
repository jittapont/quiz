package main

import (
	"reflect"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "When we provide a palindrome word, it should return true",
			word: "Deleveled",
			want: true,
		},
		{
			name: "When we provide a normal word, it should return false",
			word: "notPalindrome",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPalindrome() result mismatch got : %v, want : %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_sortString(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			name: "When we provide a word, it sort this word",
			word: "VORCM",
			want: "CMORV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortString(tt.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortString() result mismatch got : %v, want : %v", got, tt.want)
			}
		})
	}
}

func Test_groupWords(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  [][]string
	}{
		{
			name:  "When we provide a words, it should return correct group",
			words: []string{"VMRCO", "VORCM", "MCRTOX", "ZXTAC", "XZATC", "XMTCOR", "OXVS", "AZTXC", "VXOS", "RZAT", "MRCOTX", "SVXO", "TRAZ", "ZTAR", "MVOCR"},
			want:  [][]string{{"MCRTOX", "XMTCOR", "MRCOTX"}, {"OXVS", "VXOS", "SVXO"}, {"RZAT", "TRAZ", "ZTAR"}, {"VMRCO", "VORCM", "MVOCR"}, {"ZXTAC", "XZATC", "AZTXC"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupWords(tt.words)
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) == 0 && len(got[j]) == 0 {
					return false
				}
				if len(got[i]) == 0 || len(got[j]) == 0 {
					return len(got[i]) == 0
				}
				return got[i][0] < got[j][0]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupWords() result mismatch got : %v, want : %v", got, tt.want)
			}
		})
	}

}

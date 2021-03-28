package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_calculateSumUntil(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr error
	}{
		{
			name:    "When we provide a n=2, it should return a correct sum",
			n:       2,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "When we provide a n=3, it should return a correct sum",
			n:       3,
			want:    32,
			wantErr: nil,
		},
		{
			name:    "When we provide a n=4, it should return a correct sum",
			n:       4,
			want:    288,
			wantErr: nil,
		},
		{
			name:    "When we provide a n=5, it should return a correct sum",
			n:       5,
			want:    3413,
			wantErr: nil,
		},
		{
			name:    "When we provide a negative n, it should return error",
			n:       -1,
			want:    0,
			wantErr: errNIsNegative,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateSumUntil(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSumUntil() result mismatch got : %v, want : %v", got, tt.want)
			}
			if err != nil && !errors.Is(err, errNIsNegative) {
				t.Errorf("calculateSumUntil() error mismatch got : %v, want : %v", err, tt.wantErr)
			}
		})
	}
}

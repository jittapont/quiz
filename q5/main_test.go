package main

import (
	"reflect"
	"testing"
)

func Test_calculateParkingFee(t *testing.T) {
	tests := []struct {
		name   string
		hour   int
		minute int
		want   int
	}{
		{
			name:   "When we park for 4 hours and 15 minutes, it should return a correct parking fee",
			hour:   4,
			minute: 15,
			want:   50,
		},
		{
			name:   "When we park for 2 hours and 10 minutes, it should return a correct parking fee",
			hour:   2,
			minute: 10,
			want:   30,
		},
		{
			name:   "When we park for 0 hours and 10 minutes, it should return a correct parking fee",
			hour:   0,
			minute: 10,
			want:   0,
		},
		{
			name:   "When we park for 0 hours and 31 minutes, it should return a correct parking fee",
			hour:   0,
			minute: 31,
			want:   10,
		},
		{
			name:   "When we park for 1 hours and 40 minutes, it should return a correct parking fee",
			hour:   1,
			minute: 40,
			want:   30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateParkingFee(tt.hour, tt.minute)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateParkingFee() result mismatch got : %v, want : %v", got, tt.want)
			}
		})
	}
}

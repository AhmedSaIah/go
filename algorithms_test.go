package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := map[string]struct {
		args args
		want []int
	}{
		"Case array was already sorted": {
			args: args{[]int{1, 2, 4, 5}},
			want: []int{1, 2, 4, 5},
		},
		"Case array was not sorted": {
			args: args{[]int{1, 9, 4, -3}},
			want: []int{-3, 1, 4, 9},
		},
		"Case array was of length 1": {
			args: args{[]int{100}},
			want: []int{100},
		},
		"Case array was empty": {
			args: args{[]int{}},
			want: []int{},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			if got := MergeSort(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wanted %v, but got %v", got, tt.want)
			}
		})
	}
}

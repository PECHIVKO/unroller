package main

import (
	"reflect"
	"testing"
)

var square [][]int = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}
var squareRotated [][]int = [][]int{
	{3, 6, 9},
	{2, 5, 8},
	{1, 4, 7},
}
var squareUnrolled []int = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

var rect1 [][]int = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{10, 11, 12},
}
var rect1Rotated [][]int = [][]int{
	{3, 6, 9, 12},
	{2, 5, 8, 11},
	{1, 4, 7, 10},
}
var rect1Unrolled []int = []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8}

var rect2 [][]int = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
}
var rect2Rotated [][]int = [][]int{
	{4, 8, 12},
	{3, 7, 11},
	{2, 6, 10},
	{1, 5, 9},
}
var rect2Unrolled []int = []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

var emptyMatrix [][]int

var emptySlice []int

func Test_unroll(t *testing.T) {
	type args struct {
		rolled   [][]int
		unrolled []int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]int
		want1 []int
	}{
		{"Test_square", args{square, emptySlice}, emptyMatrix, squareUnrolled},
		{"Test_rect1", args{rect1, emptySlice}, emptyMatrix, rect1Unrolled},
		{"Test_rect2", args{rect2, emptySlice}, emptyMatrix, rect2Unrolled},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := unroll(tt.args.rolled, tt.args.unrolled)
			for i := range got {
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("unroll() got[%v][%v] = %v, want[%v][%v] = %v", i, j, got, i, j, tt.want)
					}
				}
			}
			for i := range got1 {
				if got1[i] != tt.want1[i] {
					t.Errorf("unroll() got1[%v] = %v, want1[%v] = %v", i, got, i, tt.want)
				}
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		unrotated [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Test_square", args{square}, squareRotated},
		{"Test_rect1", args{rect1}, rect1Rotated},
		{"Test_rect2", args{rect2}, rect2Rotated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.unrotated); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

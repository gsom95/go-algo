package search

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		arr  []int
		find int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.arr, tt.args.find); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchRecursive(t *testing.T) {
	type args struct {
		arr  []int
		find int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRecursive(tt.args.arr, tt.args.find); got != tt.want {
				t.Errorf("SearchRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRecursive(t *testing.T) {
	type args struct {
		arr  []int
		find int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRecursive(tt.args.arr, tt.args.find, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("searchRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

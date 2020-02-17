package sudoku

import (
	"reflect"
	"testing"
)

func TestField_GetCol(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		f    *Field
		args args
		want []int
	}{
		{
			"col0",
			f1,
			args{0},
			[]int{1, 4, 7, 2, 5, 8, 3, 6, 9},
		},
		{
			"col3",
			f1,
			args{3},
			[]int{4, 7, 1, 5, 8, 2, 6, 9, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.GetCol(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.GetCol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_GetQuad(t *testing.T) {
	type args struct {
		r int
		c int
	}
	tests := []struct {
		name string
		f    *Field
		args args
		want []int
	}{
		{
			"q1:0,0",
			f1,
			args{0, 0},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"q1:2,1",
			f1,
			args{2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"q5:3,4",
			f1,
			args{3, 4},
			[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
		},
		{
			"q7:8,2",
			f1,
			args{8, 2},
			[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.GetQuad(tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.GetQuad() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMissingValues(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"123",
			args{[]int{4, 5, 6, 7, 8, 9}},
			[]int{1, 2, 3},
		},
		{
			"1234",
			args{[]int{9, 6, 5, 5, 9, 0, 0, 5, 6, 7, 8, 9}},
			[]int{1, 2, 3, 4},
		},
		{
			"",
			args{[]int{4, 5, 6, 7, 8, 9, 1, 2, 3}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingValues(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MissingValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvExample(t *testing.T) {
	tests := []struct {
		name   string
		input  *Field
		expect *Field
	}{
		{"f2", f2, f2Solved},
		{"f3", f3, f3Solved},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.input
			f.Solve()
			fSolved := tt.expect
			if !reflect.DeepEqual(f.Values, fSolved.Values) {
				t.Errorf("Got:\n%s\nExp:\n%s\n", f, fSolved)
			}
		})
	}
}

var f1 = &Field{
	Values: [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
		[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
		[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
		[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
		[]int{6, 7, 8, 9, 1, 2, 3, 4, 4},
		[]int{9, 1, 2, 3, 4, 5, 6, 7, 6},
	},
}

var f2 = &Field{
	Values: [][]int{
		[]int{0, 0, 0, 0, 0, 0, 6, 2, 1},
		[]int{0, 9, 1, 0, 6, 0, 0, 8, 4},
		[]int{0, 6, 8, 0, 2, 7, 0, 0, 5},
		[]int{0, 0, 0, 8, 0, 0, 9, 0, 0},
		[]int{0, 1, 7, 0, 5, 0, 2, 4, 0},
		[]int{0, 0, 3, 0, 0, 4, 0, 0, 0},
		[]int{1, 0, 0, 7, 8, 0, 4, 5, 0},
		[]int{8, 7, 0, 0, 3, 0, 1, 6, 0},
		[]int{3, 5, 4, 0, 0, 0, 0, 0, 0},
	},
}

var f3 = &Field{
	Values: [][]int{
		[]int{0, 0, 2, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 6, 0, 3, 8, 0, 0, 0},
		[]int{7, 5, 0, 0, 4, 6, 0, 0, 0},
		[]int{0, 0, 0, 4, 0, 0, 9, 2, 0},
		[]int{0, 6, 5, 0, 9, 0, 4, 1, 0},
		[]int{0, 4, 9, 0, 0, 3, 0, 0, 0},
		[]int{0, 0, 0, 6, 1, 0, 0, 8, 9},
		[]int{0, 0, 0, 3, 7, 0, 5, 6, 1},
		[]int{0, 0, 0, 0, 0, 0, 7, 0, 0},
	},
}
var f2Solved = &Field{
	Values: [][]int{
		[]int{7, 3, 5, 9, 4, 8, 6, 2, 1},
		[]int{2, 9, 1, 5, 6, 3, 7, 8, 4},
		[]int{4, 6, 8, 1, 2, 7, 3, 9, 5},
		[]int{5, 4, 2, 8, 7, 1, 9, 3, 6},
		[]int{9, 1, 7, 3, 5, 6, 2, 4, 8},
		[]int{6, 8, 3, 2, 9, 4, 5, 1, 7},
		[]int{1, 2, 6, 7, 8, 9, 4, 5, 3},
		[]int{8, 7, 9, 4, 3, 5, 1, 6, 2},
		[]int{3, 5, 4, 6, 1, 2, 8, 7, 9},
	},
}

var f3Solved = &Field{
	Values: [][]int{
		[]int{4, 8, 2, 1, 5, 9, 6, 3, 7},
		[]int{9, 1, 6, 7, 3, 8, 2, 5, 4},
		[]int{7, 5, 3, 2, 4, 6, 1, 9, 8},
		[]int{3, 7, 8, 4, 6, 1, 9, 2, 5},
		[]int{2, 6, 5, 8, 9, 7, 4, 1, 3},
		[]int{1, 4, 9, 5, 2, 3, 8, 7, 6},
		[]int{5, 2, 7, 6, 1, 4, 3, 8, 9},
		[]int{8, 9, 4, 3, 7, 2, 5, 6, 1},
		[]int{6, 3, 1, 9, 8, 5, 7, 4, 2},
	},
}

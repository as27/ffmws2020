package euler01

import "testing"

func TestVielfache(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"10",
			args{10},
			23,
		},
		{
			"11",
			args{11},
			23,
		},
		{
			"1",
			args{1},
			0,
		},
		{
			"5",
			args{5},
			8,
		},
		{
			"15",
			args{15},
			50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vielfache(tt.args.n); got != tt.want {
				t.Errorf("Vielfache() = %v, want %v", got, tt.want)
			}
		})
	}
}

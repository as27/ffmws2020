package modulo

import "testing"

func TestModulo(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"modulo 10 5",
			args{10, 5},
			0,
		},
		{
			"modulo 10 4",
			args{10, 4},
			2,
		},
		{
			"modulo 7 3",
			args{7, 3},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Modulo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Modulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

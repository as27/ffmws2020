package primzahlen

import "testing"

func TestBerechnePrimzahl(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"1",
			args{1},
			int64(2),
		},
		{
			"6",
			args{6},
			int64(13),
		},
		{
			"25",
			args{25},
			int64(97),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BerechnePrimzahl(tt.args.n); got != tt.want {
				t.Errorf("BerechnePrimzahl() = %v, want %v", got, tt.want)
			}
		})
	}
}

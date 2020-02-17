package split

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		r    io.Reader
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    Chunks
		wantErr bool
	}{
		{
			"split letzter chunk nicht voll",
			args{
				bytes.NewBuffer([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
				5,
			},
			Chunks{
				Chunk{0, 1, 2, 3, 4},
				Chunk{5, 6, 7, 8, 9},
				Chunk{10, 11, 12, 13, 14},
				Chunk{15},
			},
			false,
		},
		{
			"split letzter chunk voll",
			args{bytes.NewBuffer([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
				5,
			},
			Chunks{
				Chunk{0, 1, 2, 3, 4},
				Chunk{5, 6, 7, 8, 9},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Split(tt.args.r, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMerge() {
	c := Chunks{
		Chunk{1, 2, 3},
		Chunk{4, 5},
	}
	b := &bytes.Buffer{}
	Merge(c, b)
	fmt.Printf("%v", b.Bytes())
	// Output: [1 2 3 4 5]
}

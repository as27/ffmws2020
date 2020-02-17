package getval

import "testing"

func TestGetIndex(t *testing.T) {
	type args struct {
		search string
		sl     []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			"index 1",
			args{
				"eins",
				[]string{"null", "eins", "zwei", "drei"},
			},
			1,
			true,
		},
		{
			"index 3",
			args{
				"drei",
				[]string{"null", "eins", "zwei", "drei"},
			},
			3,
			true,
		},
		{
			"nicht vorhanden",
			args{
				"acht",
				[]string{"null", "eins", "zwei", "drei"},
			},
			-1,
			false,
		},
		{
			"leeres slice",
			args{
				"acht",
				[]string{},
			},
			-1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetIndex(tt.args.search, tt.args.sl)
			if got != tt.want {
				t.Errorf("GetVal() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetVal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

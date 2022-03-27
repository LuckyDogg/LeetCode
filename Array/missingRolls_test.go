package Array

import (
	"reflect"
	"testing"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				rolls: []int{3, 2, 4, 3},
				mean:  4,
				n:     2,
			},
			want: []int{6, 6},
		},
		{
			name: "test2",
			args: args{
				rolls: []int{1, 5, 6},
				mean:  3,
				n:     4,
			},
			want: []int{3, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}

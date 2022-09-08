package pgslice

import (
	"reflect"
	"testing"
)

func TestDiffSliceItem(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s1: []string{"asd", "dd", "my love", "ai"},
				s2: []string{"12", "asd", "dd", "aas", "au", "ai"},
			},
			want: []string{"my love", "12", "aas", "au"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffSliceItem(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffSliceItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

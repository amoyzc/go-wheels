package string_set

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSetToSlice(t *testing.T) {
	type args struct {
		set map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Happy Path",
			args: args{
				set: map[string]struct{}{"a": {}, "b": {}, "c": {}},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetToSlice(tt.args.set); len(got) != len(tt.want) {
				t.Errorf("SetToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToSet(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "Happy Path",
			args: args{
				slice: []string{"a", "b", "c"},
			},
			want: map[string]struct{}{"a": {}, "b": {}, "c": {}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToSet(tt.args.slice); !cmp.Equal(got, tt.want) {
				t.Errorf("SliceToSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

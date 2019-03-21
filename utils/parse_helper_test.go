package utils

import (
	"reflect"
	"testing"
)

func TestSplitAndParseUint64(t *testing.T) {
	tests := []struct {
		name string
		s    string
		sep  string
		want uint64
	}{
		{
			name: "OK",
			s:    "1000 KB",
			sep:  " ",
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitAndParseUint64(tt.s, tt.sep)
			if got != tt.want {
				t.Errorf("SplitAndParseUint64 error got = %#v, want = %#v", got, tt.want)
			}
		})
	}
}

func TestSplitAndRangeParseInt64(t *testing.T) {
	tests := []struct {
		name string
		s    string
		sep  string
		want []int64
	}{
		{
			name: "OK",
			s:    "100 200 300 400 500",
			sep:  " ",
			want: []int64{100, 200, 300, 400, 500},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitAndRangeParseInt64(tt.s, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitAndRangeParseInt64 error got = %#v, want = %#v", got, tt.want)
			}
		})
	}
}

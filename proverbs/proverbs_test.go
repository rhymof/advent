package proverbs_test

import (
	"testing"

	"github.com/rhymof/advent/proverbs"
)

func TestProverbsFromDate(t *testing.T) {
	tests := [...]struct {
		date int
		want string
	}{
		{date: 1, want: "Concurrency is not parallelism."},
		{date: 2, want: "Channels orchestrate; mutexes serialize."},
		{date: 19, want: "Don't communicate by sharing memory, share memory by communicating."},
	}
	for _, tt := range tests {
		got := proverbs.FromDate(tt.date)
		if got != tt.want {
			t.Errorf("want %q, but got %q", tt.want, got)
		}
	}
}

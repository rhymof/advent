package main

import (
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	tt := []struct {
		inputString string
		want        int
	}{
		{"2021-11-30T23:59:59+09:00", -1},
		{"2021-12-01T00:00:00+09:00", 24},
		{"2021-12-24T23:59:59+09:00", 1},
		{"2021-12-25T23:59:59+09:00", 0},
		{"2021-12-26T00:00:00+09:00", -1},
	}
	for i, testcase := range tt {
		input, err := time.Parse(time.RFC3339, testcase.inputString)
		if err != nil {
			t.Errorf("test is wrong. case = %d", i)
		}
		got := cowntdown(input)
		if got != testcase.want {
			t.Errorf("case: %s, got: %d, want:%d", input.Format(time.RFC3339), got, testcase.want)
		}
	}
}

func Test_isChristmas(t *testing.T) {
	tests := []struct {
		inputString string
		want        bool
	}{
		{"2021-11-30T23:59:59+09:00", false},
		{"2021-12-01T00:00:00+09:00", false},
		{"2021-12-24T23:59:59+09:00", false},
		{"2021-12-25T23:59:59+09:00", true},
		{"2021-12-26T00:00:00+09:00", false},
	}
	for _, tt := range tests {
		t.Run(tt.inputString, func(t *testing.T) {
			input, err := time.Parse(time.RFC3339, tt.inputString)
			if err != nil {
				t.Fatalf("test is wrong. err: %v", err)
			}
			if got := isChristmas(input); got != tt.want {
				t.Errorf("isChristmas() = %v, want %v", got, tt.want)
			}
		})
	}
}

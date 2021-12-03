package main

import (
	"io"
	"net/http/httptest"
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

func Test_handler(t *testing.T) {
	backupDecorates(t)

	tests := []struct {
		name        string
		path        string
		now         time.Time
		wantStatus  int
		wantMessage string
	}{
		{
			name:       "not christmas",
			path:       "/",
			now:        time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC),
			wantStatus: 200,
		},
		{
			name:        "christmas",
			path:        "/",
			now:         time.Date(2021, 12, 25, 0, 0, 0, 0, time.UTC),
			wantStatus:  200,
			wantMessage: "Hi, Merry Christmas !🎅🎄✨",
		},
		{
			name:        "christmas with sachiko",
			path:        "/sachiko",
			now:         time.Date(2021, 12, 25, 0, 0, 0, 0, time.UTC),
			wantStatus:  200,
			wantMessage: "Hi, Merry Christmas sachiko!🎅🎄✨",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restore := setCurrentTime(tt.now)
			defer restore()
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			handler(rec, req)
			if tt.wantStatus != rec.Code {
				t.Fatalf("want status %d, but got %d", tt.wantStatus, rec.Code)
			}
			gotBody := rec.Body.String()
			if tt.wantMessage != gotBody {
				t.Fatalf("- want body %s\n- got body %s", tt.wantMessage, gotBody)
			}
		})
	}
}

func setCurrentTime(t time.Time) (restore func()) {
	backup := nowFunc
	nowFunc = func() time.Time {
		return t
	}
	return func() {
		nowFunc = backup
	}
}

func backupDecorates(t *testing.T) {
	backup := decorateFunc
	backupWriter := decorateWriterFunc
	decorateFunc = func(msg string) string { return msg }
	decorateWriterFunc = func(w io.Writer) io.Writer { return w }
	t.Cleanup(func() {
		decorateFunc = backup
		decorateWriterFunc = backupWriter
	})
}

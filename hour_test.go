package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_Hour(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale  string
		options Options
		want    string
	}{
		{"lv", Options{Hour: Hour2Digit}, "03"},
		{"fa", Options{Hour: HourNumeric}, "۳"},
		{"fa", Options{Hour: Hour2Digit}, "۰۳"},
	}

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.locale+"_"+tt.options.Hour.String(), func(t *testing.T) {
			t.Parallel()

			got := NewDateTimeFormat(language.Make(tt.locale), tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

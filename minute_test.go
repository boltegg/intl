package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_Minute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale  string
		options Options
		want    string
	}{
		{"lv", Options{Minute: MinuteNumeric}, "04"},
		{"fa", Options{Minute: MinuteNumeric}, "۰۴"},
		{"fa", Options{Minute: Minute2Digit}, "۰۴"},
	}

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.locale+"_"+tt.options.Minute.String(), func(t *testing.T) {
			t.Parallel()

			got := NewDateTimeFormat(language.Make(tt.locale), tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

func TestDateTimeFormat_HourMinute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale  string
		options Options
		want    string
	}{
		{"lv", Options{Hour: HourNumeric, Minute: MinuteNumeric}, "03:04"},
		{"es", Options{Hour: HourNumeric, Minute: MinuteNumeric}, "3:04"},
		{"fa", Options{Hour: HourNumeric, Minute: MinuteNumeric}, "۰۳:۰۴"},
	}

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.locale, func(t *testing.T) {
			t.Parallel()

			got := NewDateTimeFormat(language.Make(tt.locale), tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

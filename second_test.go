package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_Second(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale  string
		options Options
		want    string
	}{
		{"lv", Options{Second: SecondNumeric}, "05"},
		{"fa", Options{Second: SecondNumeric}, "۰۵"},
		{"fa", Options{Second: Second2Digit}, "۰۵"},
	}

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.locale+"_"+tt.options.Second.String(), func(t *testing.T) {
			t.Parallel()

			got := NewDateTimeFormat(language.Make(tt.locale), tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

func TestDateTimeFormat_MinuteSecond(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale  string
		options Options
		want    string
	}{
		{"lv", Options{Minute: MinuteNumeric, Second: SecondNumeric}, "04:05"},
		{"es", Options{Minute: MinuteNumeric, Second: SecondNumeric}, "04:05"},
		{"fa", Options{Minute: MinuteNumeric, Second: SecondNumeric}, "۰۴:۰۵"},
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

func TestDateTimeFormat_HourMinuteSecond(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale  string
		options Options
		want    string
	}{
		{"lv", Options{Hour: HourNumeric, Minute: MinuteNumeric, Second: SecondNumeric}, "03:04:05"},
		{"es", Options{Hour: HourNumeric, Minute: MinuteNumeric, Second: SecondNumeric}, "3:04:05"},
		{"fa", Options{Hour: HourNumeric, Minute: MinuteNumeric, Second: SecondNumeric}, "۰۳:۰۴:۰۵"},
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

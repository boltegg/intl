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
		{"lv", Options{Second: SecondNumeric}, "5"},
		{"fa", Options{Second: SecondNumeric}, "۵"},
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
		{"lv", Options{Minute: MinuteNumeric, Second: SecondNumeric}, "4:5"},
		{"es", Options{Minute: MinuteNumeric, Second: SecondNumeric}, "4:5"},
		{"fa", Options{Minute: MinuteNumeric, Second: SecondNumeric}, "۴:۵"},
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
		{"lv", Options{Hour: HourNumeric, Minute: MinuteNumeric, Second: SecondNumeric}, "03:4:5"},
		{"es", Options{Hour: HourNumeric, Minute: MinuteNumeric, Second: SecondNumeric}, "3:4:5"},
		{"fa", Options{Hour: HourNumeric, Minute: MinuteNumeric, Second: SecondNumeric}, "۰۳:۴:۵"},
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

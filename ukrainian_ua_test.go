package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_UkrainianUkraine(t *testing.T) {
	t.Parallel()

	locale := language.MustParse("uk-UA")

	tests := []struct {
		name   string
		layout string
		date   time.Time
		want   string
	}{
		{"E", "E", time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC), "ср"},
		{"MEd", "MEd", time.Date(2025, 3, 26, 0, 0, 0, 0, time.UTC), "ср, 26.03"},
		{"yMMMEd", "yMMMEd", time.Date(2025, 11, 28, 0, 0, 0, 0, time.UTC), "пт, 28 лист. 2025 р."},
		{"yMMMMEEEEd", "yMMMMEEEEd", time.Date(2025, 6, 3, 0, 0, 0, 0, time.UTC), "вівторок, 3 червня 2025 р."},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewDateTimeFormatLayout(locale, tt.layout).Format(tt.date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

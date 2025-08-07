package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_Layout_Hms(t *testing.T) {
	t.Parallel()

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	got := NewDateTimeFormatLayout(language.Persian, "Hms").Format(date)

	if got != "۳:۴:۵" {
		t.Fatalf("want %q got %q", "۳:۴:۵", got)
	}
}

func TestDateTimeFormat_Layout_Patterns(t *testing.T) {
	t.Parallel()

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	tests := []struct {
		layout string
		want   string
	}{
		{"y", "2024"},
		{"yM", "1/2024"},
		{"yMEd", "Tue, 1/2/2024"},
		{"yMMM", "Jan/2024"},
		{"yMMMEd", "Tue, Jan 2, 2024"},
		{"yMMMM", "January/2024"},
		{"yMMMMEEEEd", "Tuesday, January 2, 2024"},
		{"yMMMMd", "January 2, 2024"},
		{"yMMMd", "Jan 2, 2024"},
		{"yMd", "1/2/2024"},
		{"yQQQ", "Q1 2024"},
		{"yQQQQ", "1st quarter 2024"},
	}

	for _, tt := range tests {
		t.Run(tt.layout, func(t *testing.T) {
			got := NewDateTimeFormatLayout(language.English, tt.layout).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

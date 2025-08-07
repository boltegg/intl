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

	if got != "۰۳:۰۴:۰۵" {
		t.Fatalf("want %q got %q", "۰۳:۰۴:۰۵", got)
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
		{"QQQ", "Q1"},
		{"QQQQ", "1st quarter"},
		{"d", "2"},
		{"E", "Tue"},
		{"EEEE", "Tuesday"},
		{"H", "03"},
		{"Hm", "03:04"},
		{"Hms", "03:04:05"},
		{"LLL", "Jan"},
		{"LLLL", "January"},
		{"j", "03"},
		{"jm", "03:04"},
		{"jms", "03:04:05"},
		{"m", "04"},
		{"ms", "04:05"},
		{"s", "05"},
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

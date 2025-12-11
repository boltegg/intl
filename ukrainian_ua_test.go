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
		{"M", "M", time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), "03"},
		{"E", "E", time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC), "ср"},
		{"MEd", "MEd", time.Date(2025, 3, 26, 0, 0, 0, 0, time.UTC), "ср, 26.03"},
		{"yMMMEd", "yMMMEd", time.Date(2025, 11, 28, 0, 0, 0, 0, time.UTC), "пт, 28 лист. 2025 р."},
		{"yMMMMEEEEd", "yMMMMEEEEd", time.Date(2025, 6, 3, 0, 0, 0, 0, time.UTC), "вівторок, 3 червня 2025 р."},
		{"yMMMMdjm", "yMMMMdjmm", time.Date(2023, 8, 15, 14, 5, 0, 0, time.UTC), "15 серпня 2023 р., 14:05"},
		{"yMMMdjhm", "yMMMdjm", time.Date(2022, 12, 1, 9, 30, 0, 0, time.UTC), "1 груд. 2022 р., 09:30"},
		{"yMMMEEEEdjms", "yMMMEEEEdjms", time.Date(2021, 1, 4, 16, 45, 9, 0, time.UTC), "понеділок, 4 січ. 2021 р., 16:45:09"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fmt, err := NewDateTimeFormatLayout(locale, tt.layout)
			if err != nil {
				t.Fatal(err)
			}
			got := fmt.Format(tt.date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

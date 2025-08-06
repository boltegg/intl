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

func TestDateTimeFormat_Layout_yMMMMEEEEd(t *testing.T) {
	t.Parallel()

	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	got := NewDateTimeFormatLayout(language.English, "yMMMMEEEEd").Format(date)

	want := "Tuesday, Jan/2/2024"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

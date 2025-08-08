package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_EnglishColombiaYMd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 11, 17, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-CO")

	got := NewDateTimeFormatLayout(locale, "yMd").Format(date)
	want := "17/11/2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishColombiaMEd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 6, 14, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-CO")

	got := NewDateTimeFormatLayout(locale, "MEd").Format(date)
	want := "Sat, 14/06"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

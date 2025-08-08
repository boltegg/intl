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

func TestDateTimeFormat_EnglishColombiaYMdSeptember(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 9, 5, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-CO")

	got := NewDateTimeFormatLayout(locale, "yMd").Format(date)
	want := "5/09/2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishColombiaMEd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 2, 4, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-CO")

	got := NewDateTimeFormatLayout(locale, "MEd").Format(date)
	want := "Tue, 4/02"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishColombiaMMMMd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-CO")

	got := NewDateTimeFormatLayout(locale, "MMMMd").Format(date)
	want := "1 November"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishColombiaMMMMEEEEd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 4, 3, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-CO")

	got := NewDateTimeFormatLayout(locale, "MMMMEEEEd").Format(date)
	want := "Thursday, 3 April"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

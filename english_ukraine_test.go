package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_EnglishUkraineYMd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-UA")

	got := NewDateTimeFormatLayout(locale, "yMd").Format(date)
	want := "31/12/2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishUkraineYMed(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-UA")

	got := NewDateTimeFormatLayout(locale, "yMEd").Format(date)
	want := "Tue 1/4/2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishUkraineYMd_NoPadding(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.October, 6, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-UA")

	got := NewDateTimeFormatLayout(locale, "yMd").Format(date)
	want := "6/10/2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishUkraineMMMMEEEEd(t *testing.T) {
	t.Parallel()

	date := time.Date(2014, time.November, 30, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-UA")

	got := NewDateTimeFormatLayout(locale, "MMMMEEEEd").Format(date)
	want := "Sunday, 30 November"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishUkraineMEd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-UA")

	got := NewDateTimeFormatLayout(locale, "MEd").Format(date)
	want := "Wed 1/1"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_EnglishUkraineYM(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("en-UA")

	got := NewDateTimeFormatLayout(locale, "yM").Format(date)
	want := "01/2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

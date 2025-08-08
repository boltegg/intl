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

func TestDateTimeFormat_EnglishUkraineMMMd(t *testing.T) {
        t.Parallel()

        date := time.Date(2025, time.July, 21, 0, 0, 0, 0, time.UTC)
        locale := language.MustParse("en-UA")

        got := NewDateTimeFormatLayout(locale, "MMMd").Format(date)
        want := "21 Jul"
        if got != want {
                t.Fatalf("want %q got %q", want, got)
        }
}

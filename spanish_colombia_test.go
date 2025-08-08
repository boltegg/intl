package intl

import (
	"golang.org/x/text/language"
	"testing"
	"time"
)

func TestDateTimeFormat_SpanishColombia_MMMd(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("es-CO")
	got := NewDateTimeFormat(locale, Options{Month: MonthShort, Day: DayNumeric}).Format(date)
	want := "2 de sep."
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_SpanishColombia_jm(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 1, 1, 4, 0, 0, 0, time.UTC)
	locale := language.MustParse("es-CO")

	fmt, err := NewDateTimeFormatLayout(locale, "jm")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Format(date)
	want := "4:00 a.m."
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_SpanishColombia_MMM(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("es-CO")

	fmt, err := NewDateTimeFormatLayout(locale, "MMM")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Format(date)
	want := "sept."
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_SpanishColombia_yMMM(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("es-CO")

	fmt, err := NewDateTimeFormatLayout(locale, "yMMM")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Format(date)
	want := "abr. de 2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

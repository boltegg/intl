package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_RussianUkraine(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("ru-UA")

	got := NewDateTimeFormat(locale, Options{Year: YearNumeric, Quarter: QuarterShort}).Format(date)
	want := "2-й кв. 2025 г."
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_RussianUkraineWeekday(t *testing.T) {
	t.Parallel()

	date := time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("ru-UA")

	got := NewDateTimeFormatLayout(locale, "E").Format(date)
	want := "Ср"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_RussianUkraineYMed(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 5, 5, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("ru-UA")

	got := NewDateTimeFormatLayout(locale, "yMEd").Format(date)
	want := "Пн, 5.05.2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestDateTimeFormat_RussianUkraineYMd(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	locale := language.MustParse("ru-UA")

	got := NewDateTimeFormatLayout(locale, "yMd").Format(date)
	want := "01.01.2025"
	if got != want {
		t.Fatalf("want %q got %q", want, got)
	}
}

package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_WeekdayTime(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	got := NewDateTimeFormat(language.English, Options{Weekday: WeekdayShort, Hour: Hour2Digit, Minute: Minute2Digit}).Format(date)
	if got != "Tue, 03:04" {
		t.Fatalf("want %q got %q", "Tue, 03:04", got)
	}
}

func TestDateTimeFormat_MonthWeekdayTime(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	opts := Options{Month: MonthShort, Weekday: WeekdayShort, Hour: Hour2Digit, Minute: Minute2Digit}
	got := NewDateTimeFormat(language.English, opts).Format(date)
	if got != "Jan, Tue, 03:04" {
		t.Fatalf("want %q got %q", "Jan, Tue, 03:04", got)
	}
}

func TestDateTimeFormat_YearMonthWeekdayTime(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	opts := Options{Year: YearNumeric, Month: MonthShort, Weekday: WeekdayShort, Hour: Hour2Digit, Minute: Minute2Digit}
	got := NewDateTimeFormat(language.English, opts).Format(date)
	if got != "2024, Jan, Tue, 03:04" {
		t.Fatalf("want %q got %q", "2024, Jan, Tue, 03:04", got)
	}
}

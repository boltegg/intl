package intl

import (
	"golang.org/x/text/language"
	"testing"
	"time"
)

func TestDateTimeFormat_Month(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	tests := []struct {
		options Options
		want    string
	}{
		{Options{Month: MonthNumeric}, "1"},
		{Options{Month: MonthShort}, "Jan"},
		{Options{Month: MonthLong}, "January"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.options.Month.String(), func(t *testing.T) {
			t.Parallel()
			got := NewDateTimeFormat(language.English, tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

func TestDateTimeFormat_MonthDay(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	tests := []struct {
		options Options
		want    string
	}{
		{Options{Month: MonthNumeric, Day: DayNumeric}, "1/2"},
		{Options{Month: MonthShort, Day: DayNumeric}, "Jan 2"},
		{Options{Month: MonthLong, Day: DayNumeric}, "January 2"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.options.Month.String(), func(t *testing.T) {
			t.Parallel()
			got := NewDateTimeFormat(language.English, tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

func TestDateTimeFormat_WeekdayMonthDay(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	tests := []struct {
		options Options
		want    string
	}{
		{Options{Weekday: WeekdayShort, Month: MonthNumeric, Day: DayNumeric}, "Tue, 1/2"},
		{Options{Weekday: WeekdayShort, Month: MonthShort, Day: DayNumeric}, "Tue, Jan 2"},
		{Options{Weekday: WeekdayLong, Month: MonthLong, Day: DayNumeric}, "Tuesday, January 2"},
	}
	for _, tt := range tests {
		tt := tt
		name := tt.options.Weekday.String() + "_" + tt.options.Month.String()
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := NewDateTimeFormat(language.English, tt.options).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

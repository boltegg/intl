package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_Ukrainian(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	locale := language.MustParse("uk")

	tests := []struct {
		name string
		opts Options
		want string
	}{
		{"Month", Options{Month: MonthLong}, "січень"},
		{"MonthDay", Options{Month: MonthLong, Day: DayNumeric}, "2 січня"},
		{"WeekdayMonthDay", Options{Weekday: WeekdayShort, Month: MonthLong, Day: DayNumeric}, "вт, 2 січня"},
		{"WeekdayLongMonthDay", Options{Weekday: WeekdayLong, Month: MonthLong, Day: DayNumeric}, "вівторок, 2 січня"},
		{"YearMonthDay", Options{Year: YearNumeric, Month: MonthLong, Day: DayNumeric}, "2 січня 2024 р."},
		{"YearMonth", Options{Year: YearNumeric, Month: MonthLong}, "січень 2024 р."},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewDateTimeFormat(locale, tt.opts).Format(date)
			if got != tt.want {
				t.Fatalf("want %q got %q", tt.want, got)
			}
		})
	}
}

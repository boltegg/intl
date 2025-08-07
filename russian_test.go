package intl

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestDateTimeFormat_Russian(t *testing.T) {
	t.Parallel()
	date := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	locale := language.MustParse("ru")

	tests := []struct {
		name string
		opts Options
		want string
	}{
		{"Month", Options{Month: MonthLong}, "январь"},
		{"MonthDay", Options{Month: MonthLong, Day: DayNumeric}, "2 января"},
		{"WeekdayMonthDay", Options{Weekday: WeekdayShort, Month: MonthLong, Day: DayNumeric}, "Вт, 2 января"},
		{"WeekdayLongMonthDay", Options{Weekday: WeekdayLong, Month: MonthLong, Day: DayNumeric}, "вторник, 2 января"},
		{"YearMonthDay", Options{Year: YearNumeric, Month: MonthLong, Day: DayNumeric}, "2 января 2024 г."},
		{"YearMonth", Options{Year: YearNumeric, Month: MonthLong}, "январь 2024 г."},
		{"QuarterShort", Options{Quarter: QuarterShort}, "1-й кв."},
		{"YearQuarterShort", Options{Year: YearNumeric, Quarter: QuarterShort}, "1-й кв. 2024"},
		{"QuarterLong", Options{Quarter: QuarterLong}, "1-й квартал"},
		{"YearQuarterLong", Options{Year: YearNumeric, Quarter: QuarterLong}, "1-й квартал 2024 г."},
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

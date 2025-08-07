package cldr

import (
	"strings"
	"time"
)

type TimeReader interface {
	Year() int
	Month() time.Month
	Day() int
	Weekday() time.Weekday
	Hour() int
	Minute() int
	Second() int
}

type Fmt []FmtFunc

func (f Fmt) Format(t TimeReader) string {
	var b strings.Builder

	for _, fn := range f {
		fn.Format(&b, t)
	}

	return b.String()
}

type FmtFunc interface {
	Format(*strings.Builder, TimeReader)
}

type Text string

func (t Text) Format(b *strings.Builder, _ TimeReader) {
	b.WriteString(string(t))
}

type YearNumeric Digits

func (y YearNumeric) Format(b *strings.Builder, t TimeReader) {
	Digits(y).appendNumeric(b, t.Year())
}

type YearTwoDigit Digits

func (y YearTwoDigit) Format(b *strings.Builder, t TimeReader) {
	Digits(y).appendTwoDigit(b, t.Year())
}

type MonthNumeric Digits

func (m MonthNumeric) Format(b *strings.Builder, t TimeReader) {
	Digits(m).appendNumeric(b, int(t.Month()))
}

type MonthTwoDigit Digits

func (m MonthTwoDigit) Format(b *strings.Builder, t TimeReader) {
	Digits(m).appendTwoDigit(b, int(t.Month()))
}

type Month CalendarMonths

func (m Month) Format(b *strings.Builder, t TimeReader) {
	b.WriteString(m[t.Month()-1])
}

type Weekday CalendarWeekdays

func (w Weekday) Format(b *strings.Builder, t TimeReader) {
	b.WriteString(w[t.Weekday()])
}

type DayNumeric Digits

func (d DayNumeric) Format(b *strings.Builder, t TimeReader) {
	Digits(d).appendNumeric(b, t.Day())
}

type DayTwoDigit Digits

func (d DayTwoDigit) Format(b *strings.Builder, t TimeReader) {
	Digits(d).appendTwoDigit(b, t.Day())
}

type HourNumeric Digits

func (h HourNumeric) Format(b *strings.Builder, t TimeReader) {
	Digits(h).appendNumeric(b, t.Hour())
}

type HourTwoDigit Digits

func (h HourTwoDigit) Format(b *strings.Builder, t TimeReader) {
	Digits(h).appendTwoDigit(b, t.Hour())
}

type MinuteNumeric Digits

func (m MinuteNumeric) Format(b *strings.Builder, t TimeReader) {
	Digits(m).appendNumeric(b, t.Minute())
}

type MinuteTwoDigit Digits

func (m MinuteTwoDigit) Format(b *strings.Builder, t TimeReader) {
	Digits(m).appendTwoDigit(b, t.Minute())
}

type SecondNumeric Digits

func (s SecondNumeric) Format(b *strings.Builder, t TimeReader) {
	Digits(s).appendNumeric(b, t.Second())
}

type SecondTwoDigit Digits

func (s SecondTwoDigit) Format(b *strings.Builder, t TimeReader) {
	Digits(s).appendTwoDigit(b, t.Second())
}

// QuarterShort formats the quarter in a short form like Q1.
type QuarterShort Digits

func (q QuarterShort) Format(b *strings.Builder, t TimeReader) {
	quarter := (int(t.Month())-1)/3 + 1
	b.WriteByte('Q')
	Digits(q).appendNumeric(b, quarter)
}

// QuarterLong formats the quarter in a long form like 1st quarter.
type QuarterLong Digits

func (q QuarterLong) Format(b *strings.Builder, t TimeReader) {
	quarter := (int(t.Month())-1)/3 + 1
	Digits(q).appendNumeric(b, quarter)
	var suffix string
	switch quarter {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	default:
		suffix = "th"
	}
	b.WriteString(suffix)
	b.WriteString(" quarter")
}

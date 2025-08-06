package cldr

import (
	"strings"
	"time"
)

type TimeReader interface {
	Year() int
	Month() time.Month
	Day() int
	Hour() int
	Minute() int
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

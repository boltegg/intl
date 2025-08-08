package cldr

import (
	"strings"
	"time"

	"golang.org/x/text/language"
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

type Hour12Numeric Digits

func (h Hour12Numeric) Format(b *strings.Builder, t TimeReader) {
	hour := t.Hour() % 12
	if hour == 0 {
		hour = 12
	}
	Digits(h).appendNumeric(b, hour)
}

type Hour12TwoDigit Digits

func (h Hour12TwoDigit) Format(b *strings.Builder, t TimeReader) {
	hour := t.Hour() % 12
	if hour == 0 {
		hour = 12
	}
	Digits(h).appendTwoDigit(b, hour)
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

type dayPeriod struct {
	names [2]string
}

// DayPeriod returns a formatter for day period such as "AM" or "PM".
func DayPeriod(names [2]string) FmtFunc {
	return dayPeriod{names: names}
}

func (d dayPeriod) Format(b *strings.Builder, t TimeReader) {
	if t.Hour() < 12 {
		b.WriteString(d.names[0])
	} else {
		b.WriteString(d.names[1])
	}
}

// quarterShort formats the quarter in a short form.
type quarterShort struct {
	digits Digits
	locale language.Tag
}

// QuarterShort returns a formatter for a short quarter string like "Q1" or "1-й кв.".
func QuarterShort(locale language.Tag, digits Digits) FmtFunc {
	return quarterShort{digits: digits, locale: locale}
}

func (q quarterShort) Format(b *strings.Builder, t TimeReader) {
	quarter := (int(t.Month())-1)/3 + 1

	if base, _ := q.locale.Base(); base.String() == "uk" || base.String() == "ru" {
		Digits(q.digits).appendNumeric(b, quarter)
		b.WriteString("-й кв.")
		return
	}

	b.WriteByte('Q')
	Digits(q.digits).appendNumeric(b, quarter)
}

// quarterLong formats the quarter in a long form.
type quarterLong struct {
	digits Digits
	locale language.Tag
}

// QuarterLong returns a formatter for a long quarter string like "1st quarter" or "1-й квартал".
func QuarterLong(locale language.Tag, digits Digits) FmtFunc {
	return quarterLong{digits: digits, locale: locale}
}

func (q quarterLong) Format(b *strings.Builder, t TimeReader) {
	quarter := (int(t.Month())-1)/3 + 1

	if base, _ := q.locale.Base(); base.String() == "uk" || base.String() == "ru" {
		Digits(q.digits).appendNumeric(b, quarter)
		b.WriteString("-й квартал")
		return
	}

	Digits(q.digits).appendNumeric(b, quarter)
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

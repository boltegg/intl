package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

// seqWeekdayYearMonthDay formats a full date with a leading weekday.
func seqWeekdayYearMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace)
	seq.AddSeq(seqYearMonthDay(locale, opts))
	return seq
}

// seqWeekdayYearMonthDayPersian formats a full date with weekday for the
// Persian calendar.
func seqWeekdayYearMonthDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace)
	seq.AddSeq(seqYearMonthDayPersian(locale, opts))
	return seq
}

// seqWeekdayYearMonthDayBuddhist formats a full date with weekday for the
// Buddhist calendar.
func seqWeekdayYearMonthDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace)
	seq.AddSeq(seqYearMonthDayBuddhist(locale, opts))
	return seq
}

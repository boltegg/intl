package intl

import (
	"github.com/boltegg/intl/internal/cldr"
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

// seqWeekdayYearMonthDay formats a full date with a leading weekday.
func seqWeekdayYearMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, _, region := locale.Raw()

	if lang == cldr.RU && region == cldr.RegionUA && opts.Month.numeric() && opts.Day.numeric() {
		return symbols.NewSeq(locale).
			Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace,
				symbols.Symbol_d, '.', symbols.Symbol_MM, '.', opts.Year.symbol())
	}

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

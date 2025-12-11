package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqWeekdayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace, opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol())
	if !opts.Second.und() {
		seq.Add(symbols.TxtColon, opts.Second.symbol())
	}
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqMonthWeekdayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Month.symbolFormat(), symbols.TxtComma, symbols.TxtSpace, opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace, opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol())
	if !opts.Second.und() {
		seq.Add(symbols.TxtColon, opts.Second.symbol())
	}
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqYearMonthWeekdayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.AddSeq(seqWeekdayYearMonthDay(locale, opts))
	seq.Add(symbols.TxtComma, symbols.TxtSpace)
	if !opts.Second.und() {
		seq.AddSeq(seqHourMinuteSecond(locale, opts))
		return seq
	}

	seq.AddSeq(seqHourMinute(locale, opts))
	return seq
}

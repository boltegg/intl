package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqWeekdayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace, opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol())
	if !opts.Second.und() {
		seq.Add(symbols.TxtColon, opts.Second.symbol())
	}
	return seq
}

func seqMonthWeekdayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Month.symbolFormat(), symbols.TxtComma, symbols.TxtSpace, opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace, opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol())
	if !opts.Second.und() {
		seq.Add(symbols.TxtColon, opts.Second.symbol())
	}
	return seq
}

func seqYearMonthWeekdayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Year.symbol(), symbols.TxtComma, symbols.TxtSpace, opts.Month.symbolFormat(), symbols.TxtComma, symbols.TxtSpace, opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace, opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol())
	if !opts.Second.und() {
		seq.Add(symbols.TxtColon, opts.Second.symbol())
	}
	return seq
}

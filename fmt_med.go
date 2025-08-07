package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqWeekdayMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace)
	seq.AddSeq(seqMonthDay(locale, opts))
	return seq
}

func seqWeekdayMonthDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace)
	seq.AddSeq(seqMonthDayPersian(locale, opts))
	return seq
}

func seqWeekdayMonthDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	seq.Add(opts.Weekday.symbol(), symbols.TxtComma, symbols.TxtSpace)
	seq.AddSeq(seqMonthDayBuddhist(locale, opts))
	return seq
}

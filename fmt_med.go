package intl

import (
	"github.com/boltegg/intl/internal/cldr"
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqWeekdayMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, _, region := locale.Raw()

	if lang == cldr.EN && region == cldr.RegionUA {
		return symbols.NewSeq(locale).
			Add(opts.Weekday.symbol(), symbols.TxtSpace).
			AddSeq(seqMonthDay(locale, opts))
	}

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

package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqHourMinute(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol())
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqHourMinutePersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol())
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqHourMinuteBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol())
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

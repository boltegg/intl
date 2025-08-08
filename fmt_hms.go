package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqHourMinuteSecond(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqHourMinuteSecondPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqHourMinuteSecondBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale), symbols.TxtColon, opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

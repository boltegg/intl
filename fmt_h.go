package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqHour(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale))
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqHourPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale))
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

func seqHourBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.hourSymbol(locale))
	if opts.use12Hour(locale) {
		seq.Add(symbols.TxtSpace, symbols.Symbol_a)
	}
	return seq
}

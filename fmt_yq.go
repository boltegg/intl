package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

// seqYearQuarter formats a year and quarter, e.g. "Q1 2024" or "1st quarter 2024".
func seqYearQuarter(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale).Add(opts.Quarter.symbol(), symbols.TxtSpace, opts.Year.symbol())
	if base, _ := locale.Base(); base.String() == "uk" && !opts.Quarter.short() {
		seq.Add(symbols.TxtNNBSP, symbols.Txtр, '.')
	} else if base.String() == "ru" && !opts.Quarter.short() {
		seq.Add(symbols.TxtNNBSP, symbols.Txt00)
	}

	return seq
}

func seqYearQuarterPersian(locale language.Tag, opts Options) *symbols.Seq {
	return seqYearQuarter(locale, opts)
}

func seqYearQuarterBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return seqYearQuarter(locale, opts)
}

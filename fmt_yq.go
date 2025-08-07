package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

// seqYearQuarter formats a year and quarter, e.g. "Q1 2024" or "1st quarter 2024".
func seqYearQuarter(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Quarter.symbol(), symbols.TxtSpace, opts.Year.symbol())
}

func seqYearQuarterPersian(locale language.Tag, opts Options) *symbols.Seq {
	return seqYearQuarter(locale, opts)
}

func seqYearQuarterBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return seqYearQuarter(locale, opts)
}

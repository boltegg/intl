package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqQuarter(locale language.Tag, opt Quarter) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqQuarterPersian(locale language.Tag, opt Quarter) *symbols.Seq {
	return seqQuarter(locale, opt)
}

func seqQuarterBuddhist(locale language.Tag, opt Quarter) *symbols.Seq {
	return seqQuarter(locale, opt)
}

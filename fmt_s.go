package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqSecond(locale language.Tag, opt Second) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqSecondPersian(locale language.Tag, opt Second) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqSecondBuddhist(locale language.Tag, opt Second) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

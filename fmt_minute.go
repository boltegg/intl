package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqMinute(locale language.Tag, opt Minute) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqMinutePersian(locale language.Tag, opt Minute) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqMinuteBuddhist(locale language.Tag, opt Minute) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqMinuteSecond(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
}

func seqMinuteSecondPersian(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
}

func seqMinuteSecondBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
}

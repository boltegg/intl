package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqHourMinuteSecond(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
}

func seqHourMinuteSecondPersian(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
}

func seqHourMinuteSecondBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol(), symbols.TxtColon, opts.Second.symbol())
}

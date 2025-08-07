package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqHourMinute(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol())
}

func seqHourMinutePersian(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol())
}

func seqHourMinuteBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Hour.symbol(), symbols.TxtColon, opts.Minute.symbol())
}

package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqHour(locale language.Tag, opt Hour) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqHourPersian(locale language.Tag, opt Hour) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqHourBuddhist(locale language.Tag, opt Hour) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

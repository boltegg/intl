package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqWeekdayDay(locale language.Tag, opts Options) *symbols.Seq {
	if opts.Weekday != 0 && opts.Day != 0 {
		return symbols.NewSeq(locale).Add(opts.Weekday.symbol(), symbols.TxtSpace, opts.Day.symbol())
	}
	return symbols.NewSeq(locale).Add(symbols.TxtSpace)
}

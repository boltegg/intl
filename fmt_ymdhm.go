package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqYearMonthDayTime(locale language.Tag, opts Options) *symbols.Seq {
	seq := seqYearMonthDay(locale, opts)
	seq.Add(symbols.TxtComma, symbols.TxtSpace)

	if !opts.Second.und() {
		seq.AddSeq(seqHourMinuteSecond(locale, opts))
		return seq
	}
	seq.AddSeq(seqHourMinute(locale, opts))
	return seq
}

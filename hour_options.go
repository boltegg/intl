package intl

import (
	"github.com/boltegg/intl/internal/symbols"
	"golang.org/x/text/language"
)

func (o Options) use12Hour(locale language.Tag) bool {
	if o.hourSet {
		return o.hour12
	}
	if o.hourFromDefault {
		return is12HourLocale(locale)
	}
	return false
}

func (o Options) hourSymbol(locale language.Tag) symbols.Symbol {
	if o.use12Hour(locale) {
		if o.Hour.twoDigit() {
			return symbols.Symbol_hh
		}
		return symbols.Symbol_h
	}
	return o.Hour.symbol()
}

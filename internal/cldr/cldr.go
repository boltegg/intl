package cldr

import "golang.org/x/text/language"

func MonthNames(locale string, context, width string) CalendarMonths {
	indexes := MonthLookup[locale]

	// Russian locales mix abbreviated and wide month names when formatting
	// dates with the abbreviated width. The base data uses the abbreviated
	// set, but certain months (March, May, June, July) should use the wide
	// genitive forms. Adjust the result accordingly.
	if width == "abbreviated" && context == "format" {
		if base, _ := language.Make(locale).Base(); base.String() == "ru" {
			var out CalendarMonths
			if v := int(indexes[0]); v > 0 && v < len(CalendarMonthNames) {
				out = CalendarMonthNames[v]
			}
			if v := int(indexes[2]); v > 0 && v < len(CalendarMonthNames) {
				wide := CalendarMonthNames[v]
				out[2] = wide[2] // March
				out[4] = wide[4] // May
				out[5] = wide[5] // June
				out[6] = wide[6] // July
			}
			return out
		}
	}

	var i int

	// "abbreviated" width index is 0
	switch width {
	case "wide":
		i += 2 // 1*2
	case "narrow":
		i += 4 // 2*2
	}

	// "format" context index is 0
	if context == "stand-alone" {
		i++
	}

	if i >= 0 && i < len(indexes) { // isInBounds()
		if v := int(indexes[i]); v > 0 && v < len(CalendarMonthNames) { // isInBounds()
			return CalendarMonthNames[v]
		}
	}

	return CalendarMonths{}
}

func EraName(locale language.Tag) Era {
	era, ok := EraLookup[locale.String()]
	if ok {
		return era
	}

	lang, _ := locale.Base()

	if script, confidence := locale.Script(); confidence == language.Exact {
		era, ok := EraLookup[lang.String()+"-"+script.String()]
		if ok {
			return era
		}
	}

	if era, ok := EraLookup[lang.String()]; ok {
		return era
	}

	return EraLookup["aa"]
}

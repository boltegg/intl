package cldr

import "golang.org/x/text/language"

// DayPeriods returns localized day period names (AM, PM).
// Currently it provides a limited set with a default to English.
func DayPeriods(locale language.Tag) [2]string {
	base, _ := locale.Base()
	switch base.String() {
	case "en":
		return [2]string{"AM", "PM"}
	default:
		return [2]string{"AM", "PM"}
	}
}

package cldr

import (
	"strings"

	"golang.org/x/text/language"
)

// CalendarWeekdays stores localized weekday names.
type CalendarWeekdays [7]string

var weekdayData = map[string]map[string]CalendarWeekdays{
	"en": {
		"abbreviated": {"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		"wide":        {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		"narrow":      {"S", "M", "T", "W", "T", "F", "S"},
	},
	"es": {
		"abbreviated": {"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
		"wide":        {"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		"narrow":      {"D", "L", "M", "X", "J", "V", "S"},
	},
	"uk": {
		"abbreviated": {"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		"wide":        {"неділя", "понеділок", "вівторок", "середа", "четвер", "пʼятниця", "субота"},
		"narrow":      {"Н", "П", "В", "С", "Ч", "П", "С"},
	},
	"ru": {
		"abbreviated": {"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
		"wide":        {"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"},
		"narrow":      {"В", "П", "В", "С", "Ч", "П", "С"},
	},
}

// WeekdayNames returns localized weekday names for the given locale and width.
// If the exact locale is unknown, it falls back to the base language before
// defaulting to English.
func WeekdayNames(locale, width string) CalendarWeekdays {
	// allow pattern widths like "E", "EEEE", "EEEEE" in addition to the
	// CLDR keywords "abbreviated", "wide", and "narrow"
	if strings.Trim(width, "E") == "" { // width consists solely of 'E'
		switch n := len(width); {
		case n <= 3:
			width = "abbreviated"
		case n == 4:
			width = "wide"
		default:
			width = "narrow"
		}
	}

	if m, ok := weekdayData[locale]; ok {
		if w, ok := m[width]; ok {
			return w
		}
	}

	lang, _ := language.Make(locale).Base()
	if m, ok := weekdayData[lang.String()]; ok {
		if w, ok := m[width]; ok {
			return w
		}
	}

	if w, ok := weekdayData["en"][width]; ok {
		return w
	}

	return weekdayData["en"]["abbreviated"]
}

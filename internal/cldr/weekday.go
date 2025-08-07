package cldr

// CalendarWeekdays stores localized weekday names.
type CalendarWeekdays [7]string

var weekdayData = map[string]map[string]CalendarWeekdays{
	"en": {
		"abbreviated": {"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		"wide":        {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		"narrow":      {"S", "M", "T", "W", "T", "F", "S"},
	},
	"uk": {
		"abbreviated": {"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		"wide":        {"неділя", "понеділок", "вівторок", "середа", "четвер", "пʼятниця", "субота"},
		"narrow":      {"Н", "П", "В", "С", "Ч", "П", "С"},
	},
	"ru": {
		// CLDR uses lowercase weekday abbreviations for Russian.
		// See https://unicode-org.github.io/cldr-staging/charts/43/summary/ru.html
		// The previous implementation returned capitalized forms ("Вт"),
		// which caused tests like WeekdayMonthDay to fail because the
		// expected output is "вт, 2 января".  The data below now matches
		// the CLDR-provided lowercase abbreviations.
		"abbreviated": {"вс", "пн", "вт", "ср", "чт", "пт", "сб"},
		"wide":        {"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"},
		"narrow":      {"В", "П", "В", "С", "Ч", "П", "С"},
	},
}

// WeekdayNames returns localized weekday names for the given locale and width.
// If locale is unknown, English names are returned.
func WeekdayNames(locale, width string) CalendarWeekdays {
	if m, ok := weekdayData[locale]; ok {
		if w, ok := m[width]; ok {
			return w
		}
	}

	return weekdayData["en"]["abbreviated"]
}

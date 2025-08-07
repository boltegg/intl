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
		"abbreviated": {"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
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

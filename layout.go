package intl

import (
	"fmt"

	"golang.org/x/text/language"
)

// ParseLayout converts a CLDR skeleton layout string into [Options].
//
// The layout uses the skeleton symbols defined by CLDR such as "y", "M",
// "d", "H", etc. Only a subset of symbols understood by the package is
// supported. Unsupported symbols result in an error.
func ParseLayout(layout string) (Options, error) {
	var opts Options

	for i := 0; i < len(layout); {
		ch := layout[i]
		j := i + 1
		for j < len(layout) && layout[j] == ch {
			j++
		}
		count := j - i

		switch ch {
		case 'G':
			switch {
			case count >= 5:
				opts.Era = EraNarrow
			case count >= 4:
				opts.Era = EraLong
			default:
				opts.Era = EraShort
			}
		case 'y':
			if count == 2 {
				opts.Year = Year2Digit
			} else {
				opts.Year = YearNumeric
			}
		case 'M', 'L':
			switch count {
			case 1:
				opts.Month = MonthNumeric
			case 2:
				opts.Month = Month2Digit
			case 3:
				opts.Month = MonthShort
			case 4:
				opts.Month = MonthLong
			default:
				opts.Month = MonthNarrow
			}
		case 'Q':
			if count >= 4 {
				opts.Quarter = QuarterLong
			} else {
				opts.Quarter = QuarterShort
			}
		case 'd':
			if count == 2 {
				opts.Day = Day2Digit
			} else {
				opts.Day = DayNumeric
			}
		case 'E':
			switch {
			case count >= 5:
				opts.Weekday = WeekdayNarrow
			case count == 4:
				opts.Weekday = WeekdayLong
			default:
				opts.Weekday = WeekdayShort
			}
		case 'H', 'h':
			if count == 2 {
				opts.Hour = Hour2Digit
			} else {
				opts.Hour = HourNumeric
			}
		case 'm':
			if count == 2 {
				opts.Minute = Minute2Digit
			} else {
				opts.Minute = MinuteNumeric
			}
		case 's':
			if count == 2 {
				opts.Second = Second2Digit
			} else {
				opts.Second = SecondNumeric
			}
		default:
			return Options{}, fmt.Errorf("unsupported skeleton symbol %q", string(ch))
		}

		i = j
	}

	return opts, nil
}

// MustParseLayout is like [ParseLayout] but panics on error.
func MustParseLayout(layout string) Options {
	opts, err := ParseLayout(layout)
	if err != nil {
		panic(err)
	}

	return opts
}

// NewDateTimeFormatLayout creates [DateTimeFormat] using a CLDR skeleton
// layout string. It panics if the layout contains unsupported symbols.
func NewDateTimeFormatLayout(locale language.Tag, layout string) DateTimeFormat {
	return NewDateTimeFormat(locale, MustParseLayout(layout))
}

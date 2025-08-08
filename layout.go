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

	// single minute or second layouts should use numeric form without
	// zero-padding. When other fields are present, minutes and seconds
	// should be rendered in two-digit form for consistency.
	onlyMinute := layout == "m"
	onlySecond := layout == "s"

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
			if ch == 'L' {
				opts.MonthStandalone = true
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
		case 'H':
			if count == 2 {
				opts.Hour = Hour2Digit
			} else {
				opts.Hour = HourNumeric
			}
			opts.hour12 = false
			opts.hourSet = true
			opts.hourFromDefault = false
		case 'h':
			if count == 2 {
				opts.Hour = Hour2Digit
			} else {
				opts.Hour = HourNumeric
			}
			opts.hour12 = true
			opts.hourSet = true
			opts.hourFromDefault = false
		case 'j':
			if count == 2 {
				opts.Hour = Hour2Digit
			} else {
				opts.Hour = HourNumeric
			}
			opts.hourFromDefault = true
		case 'm':
			if count == 2 || !onlyMinute {
				opts.Minute = Minute2Digit
			} else {
				opts.Minute = MinuteNumeric
			}
		case 's':
			if count == 2 || !onlySecond {
				opts.Second = Second2Digit
			} else {
				opts.Second = SecondNumeric
			}
		default:
			return Options{}, fmt.Errorf("unsupported skeleton symbol %q", string(ch))
		}

		i = j
	}

	if !opts.Month.und() && opts.Day.und() && opts.Year.und() && opts.Weekday.und() && opts.Hour.und() && opts.Minute.und() && opts.Second.und() && opts.Era.und() && opts.Quarter.und() {
		opts.MonthStandalone = true
	}

	return opts, nil
}

// NewDateTimeFormatLayout creates [DateTimeFormat] using a CLDR skeleton
// layout string. It returns an error if the layout contains unsupported symbols.
func NewDateTimeFormatLayout(locale language.Tag, layout string) (DateTimeFormat, error) {
	opts, err := ParseLayout(layout)
	if err != nil {
		return DateTimeFormat{}, err
	}

	return NewDateTimeFormat(locale, opts), nil
}

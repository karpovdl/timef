package timef

import (
	"strings"
	"time"
)

// ToFormat convert date to desired format
func ToFormat(value, layout, format string) (string, error) {
	var (
		res string
		trd time.Time
		err error
	)

	if layout == "" {
		return value, nil
	}

	if trd, err = time.Parse(layout, value); err != nil {
		return value, err
	}

	if format != "" {
		res = trd.Format(format)
	} else {
		res = trd.String()
	}

	return res, err
}

// ToYYYYMMDD convert to YYYY-MM-DD or YYYY/MM/DD or YYYY.MM.DD or YYYYMMDD
func ToYYYYMMDD(day, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtBegin1))
	case 47: // "/""
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtBegin2))
	case 46: // "."
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtBegin3))
	case 32: // ""
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtBegin4))
	default:
		return day, nil
	}
}

// ToYYYYMMDDHHMMSS convert to YYYY-MM-DD HH24:MI:SS or YYYY/MM/DD HH24:MI:SS or YYYY.MM.DD HI:MM:SS or YYYYMMDD HH24:MI:SS
func ToYYYYMMDDHHMMSS(date, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(date, layout, GetFormat(FormatDateLongYearAtBegin21))
	case 47: // "/""
		return ToFormat(date, layout, GetFormat(FormatDateLongYearAtBegin22))
	case 46: // "."
		return ToFormat(date, layout, GetFormat(FormatDateLongYearAtBegin23))
	case 32: // ""
		return ToFormat(date, layout, GetFormat(FormatDateLongYearAtBegin24))
	default:
		return date, nil
	}
}

// ToDDMMYYYY convert to DD-MM-YYYY or DD/MM/YYYY or DD.MM.YYYY or DDMMYYYY
func ToDDMMYYYY(day, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtEnd1))
	case 47: // "/""
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtEnd2))
	case 46: // "."
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtEnd3))
	case 32: // ""
		return ToFormat(day, layout, GetFormat(FormatDayLongYearAtEnd4))
	default:
		return day, nil
	}
}

// GetFormat get a predefined format
func GetFormat(format string) string {
	switch format {
	case FormatDateLongYearAtBegin11:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":")
	case FormatDateLongYearAtBegin21:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":")
	case FormatDateLongYearAtBegin31:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli
	case FormatDateLongYearAtBegin41:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro
	case FormatDateLongYearAtBegin51:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano

	case FormatDateYearAtBegin11:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":")
	case FormatDateYearAtBegin21:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":")
	case FormatDateYearAtBegin31:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli
	case FormatDateYearAtBegin41:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro
	case FormatDateYearAtBegin51:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano

	case FormatDayLongYearAtBegin1:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-")
	case FormatDayLongYearAtBegin2:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/")
	case FormatDayLongYearAtBegin3:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".")
	case FormatDayLongYearAtBegin4:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "")

	case FormatDayYearAtBegin1:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-")
	case FormatDayYearAtBegin2:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "/")
	case FormatDayYearAtBegin3:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, ".")
	case FormatDayYearAtBegin4:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "")

	case FormatDayLongYearAtEnd1:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-")
	case FormatDayLongYearAtEnd2:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/")
	case FormatDayLongYearAtEnd3:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".")
	case FormatDayLongYearAtEnd4:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "")

	case FormatDayYearAtEnd1:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, "-")
	case FormatDayYearAtEnd2:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, "/")
	case FormatDayYearAtEnd3:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, ".")
	case FormatDayYearAtEnd4:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, "")

	default:
		return ""
	}
}

// TryConvertMonthRuToEn converts the name of the Russian month into the English name of the month
func TryConvertMonthRuToEn(m string) (string, bool) {
	if m == "" {
		return "", false
	}

	m = strings.ToLower(m)
	m = strings.ReplaceAll(m, "ё", "е")

	for ind, month := range MonthsRu {
		month = strings.ToLower(month)

		if month == m {
			return MonthsEn[ind], true
		}
	}

	for ind, month := range MonthsRuGenitive {
		month = strings.ToLower(month)

		if month == m {
			return MonthsEn[ind], true
		}
	}

	for ind, month := range MonthsRuDative {
		month = strings.ToLower(month)

		if month == m {
			return MonthsEn[ind], true
		}
	}

	for ind, month := range MonthsRuAccusative {
		month = strings.ToLower(month)

		if month == m {
			return MonthsEn[ind], true
		}
	}

	for ind, month := range MonthsRuAblative {
		month = strings.ToLower(month)

		if month == m {
			return MonthsEn[ind], true
		}
	}

	for ind, month := range MonthsRuPrepositional {
		month = strings.ToLower(month)

		if month == m {
			return MonthsEn[ind], true
		}
	}

	return "", false
}

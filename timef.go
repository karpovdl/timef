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
		return ToFormat(day, layout, GetFormat(FormatDayLong1))
	case 47: // "/""
		return ToFormat(day, layout, GetFormat(FormatDayLong2))
	case 46: // "."
		return ToFormat(day, layout, GetFormat(FormatDayLong3))
	case 32: // ""
		return ToFormat(day, layout, GetFormat(FormatDayLong4))
	default:
		return day, nil
	}
}

// ToDDMMYYYY convert to DD-MM-YYYY or DD/MM/YYYY or DD.MM.YYYY or DDMMYYYY
func ToDDMMYYYY(day, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(day, layout, GetFormat(FormatDayLongBackwards1))
	case 47: // "/""
		return ToFormat(day, layout, GetFormat(FormatDayLongBackwards2))
	case 46: // "."
		return ToFormat(day, layout, GetFormat(FormatDayLongBackwards3))
	case 32: // ""
		return ToFormat(day, layout, GetFormat(FormatDayLongBackwards4))
	default:
		return day, nil
	}
}

// GetFormat get a predefined format
func GetFormat(format string) string {
	switch format {
	case FormatDateLong11:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":")
	case FormatDateLong21:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":")
	case FormatDateLong31:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli
	case FormatDateLong41:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro
	case FormatDateLong51:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano

	case FormatDate11:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":")
	case FormatDate21:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":")
	case FormatDate31:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli
	case FormatDate41:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro
	case FormatDate51:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano

	case FormatDayLong1:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-")
	case FormatDayLong2:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/")
	case FormatDayLong3:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".")
	case FormatDayLong4:
		return strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "")

	case FormatDay1:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-")
	case FormatDay2:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "/")
	case FormatDay3:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, ".")
	case FormatDay4:
		return strings.Join([]string{Year, ZeroMonth, ZeroDay}, "")

	case FormatDayLongBackwards1:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-")
	case FormatDayLongBackwards2:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/")
	case FormatDayLongBackwards3:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".")
	case FormatDayLongBackwards4:
		return strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "")

	case FormatDayBackwards1:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, "-")
	case FormatDayBackwards2:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, "/")
	case FormatDayBackwards3:
		return strings.Join([]string{ZeroDay, ZeroMonth, Year}, ".")
	case FormatDayBackwards4:
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

	for indx, month := range monthsRu {
		month = strings.ToLower(month)

		if month == m {
			return monthsEn[indx], true
		}
	}

	for indx, month := range monthsRuGenitive {
		month = strings.ToLower(month)

		if month == m {
			return monthsEn[indx], true
		}
	}

	for indx, month := range monthsRuDative {
		month = strings.ToLower(month)

		if month == m {
			return monthsEn[indx], true
		}
	}

	for indx, month := range monthsRuAccusative {
		month = strings.ToLower(month)

		if month == m {
			return monthsEn[indx], true
		}
	}

	for indx, month := range monthsRuAblative {
		month = strings.ToLower(month)

		if month == m {
			return monthsEn[indx], true
		}
	}

	for indx, month := range monthsRuPrepositional {
		month = strings.ToLower(month)

		if month == m {
			return monthsEn[indx], true
		}
	}

	return "", false
}

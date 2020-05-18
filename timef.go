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
		return ToFormat(day, layout, Format[FormatDayLongYearAtBegin1])
	case 47: // "/""
		return ToFormat(day, layout, Format[FormatDayLongYearAtBegin2])
	case 46: // "."
		return ToFormat(day, layout, Format[FormatDayLongYearAtBegin3])
	case 32: // ""
		return ToFormat(day, layout, Format[FormatDayLongYearAtBegin4])
	default:
		return day, nil
	}
}

// ToYYYYMMDDHHMMSS convert to YYYY-MM-DD HH24:MI:SS or YYYY/MM/DD HH24:MI:SS or YYYY.MM.DD HI:MM:SS or YYYYMMDD HH24:MI:SS
func ToYYYYMMDDHHMMSS(date, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(date, layout, Format[FormatDateLongYearAtBegin21])
	case 47: // "/""
		return ToFormat(date, layout, Format[FormatDateLongYearAtBegin22])
	case 46: // "."
		return ToFormat(date, layout, Format[FormatDateLongYearAtBegin23])
	case 32: // ""
		return ToFormat(date, layout, Format[FormatDateLongYearAtBegin24])
	default:
		return date, nil
	}
}

// ToDDMMYYYY convert to DD-MM-YYYY or DD/MM/YYYY or DD.MM.YYYY or DDMMYYYY
func ToDDMMYYYY(day, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(day, layout, Format[FormatDayLongYearAtEnd1])
	case 47: // "/""
		return ToFormat(day, layout, Format[FormatDayLongYearAtEnd2])
	case 46: // "."
		return ToFormat(day, layout, Format[FormatDayLongYearAtEnd3])
	case 32: // ""
		return ToFormat(day, layout, Format[FormatDayLongYearAtEnd4])
	default:
		return day, nil
	}
}

// ToDDMMYYYYHHMMSS convert to DD-MM-YYYY HH24:MI:SS or DD/MM/YYYY HH24:MI:SS or DD.MM.YYYY HI:MM:SS or DDMMYYYY HH24:MI:SS
func ToDDMMYYYYHHMMSS(date, layout string, separator rune) (string, error) {
	switch separator {
	case 45: // "-"
		return ToFormat(date, layout, Format[FormatDateLongYearAtEnd21])
	case 47: // "/""
		return ToFormat(date, layout, Format[FormatDateLongYearAtEnd22])
	case 46: // "."
		return ToFormat(date, layout, Format[FormatDateLongYearAtEnd23])
	case 32: // ""
		return ToFormat(date, layout, Format[FormatDateLongYearAtEnd24])
	default:
		return date, nil
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

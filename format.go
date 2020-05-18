package timef

import "strings"

type format map[string]string

// Format list of predefined formats
var Format = format{
	// Format for date year at begin plus timestamp
	FormatDateLongYearAtBegin11: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateLongYearAtBegin12: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateLongYearAtBegin13: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateLongYearAtBegin14: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),

	FormatDateLongYearAtBegin21: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateLongYearAtBegin22: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateLongYearAtBegin23: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateLongYearAtBegin24: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),

	FormatDateLongYearAtBegin31: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateLongYearAtBegin32: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateLongYearAtBegin33: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateLongYearAtBegin34: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,

	FormatDateLongYearAtBegin41: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateLongYearAtBegin42: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateLongYearAtBegin43: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateLongYearAtBegin44: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,

	FormatDateLongYearAtBegin51: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,
	FormatDateLongYearAtBegin52: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,
	FormatDateLongYearAtBegin53: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,
	FormatDateLongYearAtBegin54: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,

	FormatDateYearAtBegin11: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateYearAtBegin12: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateYearAtBegin13: strings.Join([]string{Year, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateYearAtBegin14: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),

	FormatDateYearAtBegin21: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateYearAtBegin22: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateYearAtBegin23: strings.Join([]string{Year, ZeroMonth, ZeroDay}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateYearAtBegin24: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),

	FormatDateYearAtBegin31: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateYearAtBegin41: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateYearAtBegin51: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,

	// Format for date year at end plus timestamp
	FormatDateLongYearAtEnd11: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateLongYearAtEnd12: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateLongYearAtEnd13: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateLongYearAtEnd14: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),

	FormatDateLongYearAtEnd21: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateLongYearAtEnd22: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateLongYearAtEnd23: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateLongYearAtEnd24: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),

	FormatDateLongYearAtEnd31: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateLongYearAtEnd32: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateLongYearAtEnd33: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateLongYearAtEnd34: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,

	FormatDateLongYearAtEnd41: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateLongYearAtEnd42: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateLongYearAtEnd43: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateLongYearAtEnd44: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,

	FormatDateLongYearAtEnd51: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,
	FormatDateLongYearAtEnd52: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,
	FormatDateLongYearAtEnd53: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,
	FormatDateLongYearAtEnd54: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,

	FormatDateYearAtEnd11: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "-") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateYearAtEnd12: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "/") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateYearAtEnd13: strings.Join([]string{ZeroDay, Year, ZeroMonth}, ".") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
	FormatDateYearAtEnd14: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),

	FormatDateYearAtEnd21: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateYearAtEnd22: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "/") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateYearAtEnd23: strings.Join([]string{ZeroDay, Year, ZeroMonth}, ".") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),
	FormatDateYearAtEnd24: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":"),

	FormatDateYearAtEnd31: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Milli,
	FormatDateYearAtEnd41: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Micro,
	FormatDateYearAtEnd51: strings.Join([]string{ZeroDay, Year, ZeroMonth}, "-") + " " + strings.Join([]string{Hour, ZeroMinute, ZeroSecond}, ":") + "." + Nano,

	// Format for day year at begin
	FormatDayLongYearAtBegin1: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "-"),
	FormatDayLongYearAtBegin2: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/"),
	FormatDayLongYearAtBegin3: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "."),
	FormatDayLongYearAtBegin4: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, ""),

	FormatDayYearAtBegin1: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "-"),
	FormatDayYearAtBegin2: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "/"),
	FormatDayYearAtBegin3: strings.Join([]string{Year, ZeroMonth, ZeroDay}, "."),
	FormatDayYearAtBegin4: strings.Join([]string{Year, ZeroMonth, ZeroDay}, ""),

	// Format for day year at end
	FormatDayLongYearAtEnd1: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "-"),
	FormatDayLongYearAtEnd2: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "/"),
	FormatDayLongYearAtEnd3: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, "."),
	FormatDayLongYearAtEnd4: strings.Join([]string{ZeroDay, ZeroMonth, LongYear}, ""),

	FormatDayYearAtEnd1: strings.Join([]string{ZeroDay, ZeroMonth, Year}, "-"),
	FormatDayYearAtEnd2: strings.Join([]string{ZeroDay, ZeroMonth, Year}, "/"),
	FormatDayYearAtEnd3: strings.Join([]string{ZeroDay, ZeroMonth, Year}, "."),
	FormatDayYearAtEnd4: strings.Join([]string{ZeroDay, ZeroMonth, Year}, ""),
}

// Contains returns true if value exists
func (f *format) Contains(value string) bool {
	if _, ok := (*f)[value]; ok {
		return true
	}

	return false
}

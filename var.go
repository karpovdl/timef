package timef

// List of dates in number or string format
const (
	StampDashDateLongYearAtBegin = "2006-01-02 15:04:05" // 2006-01-02 15:04:05
	StampDashDateYearAtBegin     = "06-01-02 15:04:05"   // 06-01-02 15:04:05
	StampDashDayLongYearAtBegin  = "2006-01-02"          // 2006-01-02

	StampSlashDateLongYearAtBegin = "2006/01/02 15:04:05" // 2006/01/02 15:04:05
	StampSlashDateYearAtBegin     = "06/01/02 15:04:05"   // 06/01/02 15:04:05
	StampSlashDayLongYearAtBegin  = "2006/01/02"          // 2006/01/02

	StampDotDateLongYearAtBegin = "2006.01.02 15:04:05" // 2006.01.02 15:04:05
	StampDotDateYearAtBegin     = "06.01.02 15:04:05"   // 06.01.02 15:04:05
	StampDotDayLongYearAtBegin  = "2006.01.02"          // 2006.01.02

	StampDateLongYearAtBegin = "20060102 15:04:05" // 20060102 15:04:05
	StampDateYearAtBegin     = "060102 15:04:05"   // 060102 15:04:05
	StampDayLongYearAtBegin  = "20060102"          // 20060102

	StampDashDayYearAtBegin  = "06-01-02" // 06-01-02
	StampSlashDayYearAtBegin = "06/01/02" // 06/01/02
	StampDotDayYearAtBegin   = "06.01.02" // 06.01.02
	StampDayYearAtBegin      = "060102"   // 060102

	StampDashDateLongYearAtEnd = "02-01-2006 15:04:05" // 02-01-2006 15:04:05
	StampDashDateYearAtEnd     = "02-01-06 15:04:05"   // 02-01-06 15:04:05
	StampDashDayLongYearAtEnd  = "02-01-2006"          // 02-01-2006

	StampSlashDateLongYearAtEnd = "02/01/2006 15:04:05" // 02/01/2006 15:04:05
	StampSlashDateYearAtEnd     = "02/01/06 15:04:05"   // 02/01/06 15:04:05
	StampSlashDayLongYearAtEnd  = "02/01/2006"          // 02/01/2006

	StampDotDateLongYearAtEnd = "02.01.2006 15:04:05" // 02.01.2006 15:04:05
	StampDotDateYearAtEnd     = "02.01.06 15:04:05"   // 02.01.06 15:04:05
	StampDotDayLongYearAtEnd  = "02.01.2006"          // 02.01.2006

	StampDateLongYearAtEnd = "02012006 15:04:05" // 02012006 15:04:05
	StampDateYearAtEnd     = "020106 15:04:05"   // 020106 15:04:05
	StampDayLongYearAtEnd  = "02012006"          // 02012006

	StampDashDayYearAtEnd  = "02-01-06" // 02-01-06
	StampSlashDayYearAtEnd = "02/01/06" // 02/01/06
	StampDotDayYearAtEnd   = "02.01.06" // 02.01.06
	StampDayYearAtEnd      = "020106"   // 020106

	StampTime = "15:04:05" // 15:04:05

	StampWordDayLongMonthLongYear = "2 January 2006" // 2 January 2006
	StampWordDayLongMonthYear     = "2 January 06"   // 2 January 06
	StampWordDayLongMonth         = "2 January"      // 2 January
	StampWordDayMonthYear         = "2 Jan 06"       // 2 Jan 06
	StampWordDayMonth             = "2 Jan"          // 2 Jan

	LongYear  = "2006"    // 2006
	Year      = "06"      // 06
	LongMonth = "January" // January
	Month     = "Jan"     // Jan
	ZeroMonth = "01"      // 01
	NumMonth  = "1"       // 1
	Day       = "2"       // 2
	ZeroDay   = "02"      // 02

	Hour       = "15"        // 15
	Hour12     = "3"         // 3
	Minute     = "4"         // 4
	ZeroMinute = "04"        // 04
	Second     = "5"         // 5
	ZeroSecond = "05"        // 05
	Zero       = "0"         // 0
	Milli      = "000"       // 000
	Micro      = "000000"    // 000000
	Nano       = "000000000" // 000000000
)

// Full date format and year at the beginning and timestamp
const (
	// FormatDateLongYearAtBegin11 its YYYY-MM-DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtBegin11 = "YYYY-MM-DD HH24:MI"
	// FormatDateLongYearAtBegin12 its YYYY/MM/DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtBegin12 = "YYYY/MM/DD HH24:MI"
	// FormatDateLongYearAtBegin13 its YYYY.MM.DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtBegin13 = "YYYY.MM.DD HH24:MI"
	// FormatDateLongYearAtBegin14 its YYYYMMDD HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtBegin14 = "YYYYMMDD HH24:MI"

	// FormatDateLongYearAtBegin21 its YYYY-MM-DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtBegin21 = "YYYY-MM-DD HH24:MI:SS"
	// FormatDateLongYearAtBegin22 its YYYY/MM/DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtBegin22 = "YYYY/MM/DD HH24:MI:SS"
	// FormatDateLongYearAtBegin23 its YYYY.MM.DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtBegin23 = "YYYY.MM.DD HH24:MI:SS"
	// FormatDateLongYearAtBegin24 its YYYYMMDD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtBegin24 = "YYYYMMDD HH24:MI:SS"

	// FormatDateLongYearAtBegin31 its YYYY-MM-DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtBegin31 = "YYYY-MM-DD HH24:MI:SS.FFF"
	// FormatDateLongYearAtBegin32 its YYYY/MM/DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtBegin32 = "YYYY/MM/DD HH24:MI:SS.FFF"
	// FormatDateLongYearAtBegin33 its YYYY.MM.DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtBegin33 = "YYYY.MM.DD HH24:MI:SS.FFF"
	// FormatDateLongYearAtBegin34 its YYYYMMDD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtBegin34 = "YYYYMMDD HH24:MI:SS.FFF"

	// FormatDateLongYearAtBegin41 its YYYY-MM-DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtBegin41 = "YYYY-MM-DD HH24:MI:SS.FFFFFF"
	// FormatDateLongYearAtBegin42 its YYYY/MM/DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtBegin42 = "YYYY/MM/DD HH24:MI:SS.FFFFFF"
	// FormatDateLongYearAtBegin43 its YYYY.MM.DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtBegin43 = "YYYY.MM.DD HH24:MI:SS.FFFFFF"
	// FormatDateLongYearAtBegin44 its YYYYMMDD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtBegin44 = "YYYYMMDD HH24:MI:SS.FFFFFF"

	// FormatDateLongYearAtBegin51 its YYYY-MM-DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtBegin51 = "YYYY-MM-DD HH24:MI:SS.FFFFFFFFF"
	// FormatDateLongYearAtBegin52 its YYYY/MM/DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtBegin52 = "YYYY/MM/DD HH24:MI:SS.FFFFFFFFF"
	// FormatDateLongYearAtBegin53 its YYYY.MM.DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtBegin53 = "YYYY.MM.DD HH24:MI:SS.FFFFFFFFF"
	// FormatDateLongYearAtBegin54 its YYYYMMDD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtBegin54 = "YYYYMMDD HH24:MI:SS.FFFFFFFFF"
)

// Date format and year at the beginning and timestamp
const (
	// FormatDateYearAtBegin11 its YY-MM-DD HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtBegin11 = "YY-MM-DD HH24:MI"
	// FormatDateYearAtBegin12 its YY/MM/DD HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtBegin12 = "YY/MM/DD HH24:MI"
	// FormatDateYearAtBegin13 its YY.MM.DD HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtBegin13 = "YY.MM.DD HH24:MI"
	// FormatDateYearAtBegin14 its YYMMDD HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtBegin14 = "YYMMDD HH24:MI"

	// FormatDateYearAtBegin21 its YY-MM-DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtBegin21 = "YY-MM-DD HH24:MI:SS"
	// FormatDateYearAtBegin22 its YY/MM/DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtBegin22 = "YY/MM/DD HH24:MI:SS"
	// FormatDateYearAtBegin23 its YY.MM.DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtBegin23 = "YY.MM.DD HH24:MI:SS"
	// FormatDateYearAtBegin24 its YYMMDD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtBegin24 = "YYMMDD HH24:MI:SS"

	// FormatDateYearAtBegin31 its YY-MM-DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateYearAtBegin31 = "YY-MM-DD HH24:MI:SS.FFF"

	// FormatDateYearAtBegin41 its YY-MM-DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateYearAtBegin41 = "YY-MM-DD HH24:MI:SS.FFFFFF"

	// FormatDateYearAtBegin51 its YY-MM-DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateYearAtBegin51 = "YY-MM-DD HH24:MI:SS.FFFFFFFFF"
)

// Full day format and year at the beginning
const (
	// FormatDayLongYearAtBegin1 its YYYY-MM-DD. Complete day.
	FormatDayLongYearAtBegin1 = "YYYY-MM-DD"
	// FormatDayLongYearAtBegin2 its YYYY/MM/DD. Complete day.
	FormatDayLongYearAtBegin2 = "YYYY/MM/DD"
	// FormatDayLongYearAtBegin3 its YYYY.MM.DD. Complete day.
	FormatDayLongYearAtBegin3 = "YYYY.MM.DD"
	// FormatDayLongYearAtBegin4 its YYYYMMDD. Complete day.
	FormatDayLongYearAtBegin4 = "YYYYMMDD"
)

// Day format and year at the beginning
const (
	// FormatDayYearAtBegin1 its YY-MM-DD. Complete day.
	FormatDayYearAtBegin1 = "YY-MM-DD"
	// FormatDayYearAtBegin2 its YY/MM/DD. Complete day.
	FormatDayYearAtBegin2 = "YY/MM/DD"
	// FormatDayYearAtBegin3 its YY.MM.DD. Complete day.
	FormatDayYearAtBegin3 = "YY.MM.DD"
	// FormatDayYearAtBegin4 its YYMMDD. Complete day.
	FormatDayYearAtBegin4 = "YYMMDD"
)

// Full date format and year at the end and timestamp
const (
	// FormatDateLongYearAtEnd11 its YYYY-MM-DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtEnd11 = "DD-MM-YYYY HH24:MI"
	// FormatDateLongYearAtEnd12 its DD/MM/YYYY HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtEnd12 = "DD/MM/YYYY HH24:MI"
	// FormatDateLongYearAtEnd13 its DD.MM.YYYY HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtEnd13 = "DD.MM.YYYY HH24:MI"
	// FormatDateLongYearAtEnd14 its DDMMYYYY HH24:MI. Complete date plus hours and minutes.
	FormatDateLongYearAtEnd14 = "DDMMYYYY HH24:MI"

	// FormatDateLongYearAtEnd21 its DD-MM-YYYY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtEnd21 = "DD-MM-YYYY HH24:MI:SS"
	// FormatDateLongYearAtEnd22 its DD/MM/YYYY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtEnd22 = "DD/MM/YYYY HH24:MI:SS"
	// FormatDateLongYearAtEnd23 its DD.MM.YYYY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtEnd23 = "DD.MM.YYYY HH24:MI:SS"
	// FormatDateLongYearAtEnd24 its DDMMYYYY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLongYearAtEnd24 = "DDMMYYYY HH24:MI:SS"

	// FormatDateLongYearAtEnd31 its DD-MM-YYYY HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtEnd31 = "DD-MM-YYYY HH24:MI:SS.FFF"
	// FormatDateLongYearAtEnd32 its DD/MM/YYYY HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtEnd32 = "DD/MM/YYYY HH24:MI:SS.FFF"
	// FormatDateLongYearAtEnd33 its DD.MM.YYYY HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtEnd33 = "DD.MM.YYYY HH24:MI:SS.FFF"
	// FormatDateLongYearAtEnd34 its DDMMYYYY HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLongYearAtEnd34 = "DDMMYYYY HH24:MI:SS.FFF"

	// FormatDateLongYearAtEnd41 its DD-MM-YYYY HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtEnd41 = "DD-MM-YYYY HH24:MI:SS.FFFFFF"
	// FormatDateLongYearAtEnd42 its DD/MM/YYYY HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtEnd42 = "DD/MM/YYYY HH24:MI:SS.FFFFFF"
	// FormatDateLongYearAtEnd43 its DD.MM.YYYY HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtEnd43 = "DD.MM.YYYY HH24:MI:SS.FFFFFF"
	// FormatDateLongYearAtEnd44 its DDMMYYYY HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLongYearAtEnd44 = "DDMMYYYY HH24:MI:SS.FFFFFF"

	// FormatDateLongYearAtEnd51 its DD-MM-YYYY HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtEnd51 = "DD-MM-YYYY HH24:MI:SS.FFFFFFFFF"
	// FormatDateLongYearAtEnd52 its DD/MM/YYYY HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtEnd52 = "DD/MM/YYYY HH24:MI:SS.FFFFFFFFF"
	// FormatDateLongYearAtEnd53 its DD.MM.YYYY HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtEnd53 = "DD.MM.YYYY HH24:MI:SS.FFFFFFFFF"
	// FormatDateLongYearAtEnd54 its DDMMYYYY HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLongYearAtEnd54 = "DDMMYYYY HH24:MI:SS.FFFFFFFFF"
)

// Date format and year at the end and timestamp
const (
	// FormatDateYearAtEnd11 its DD-MM-YY HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtEnd11 = "DD-MM-YY HH24:MI"
	// FormatDateYearAtEnd12 its DD/MM/YY HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtEnd12 = "DD/MM/YY HH24:MI"
	// FormatDateYearAtEnd13 its DD.MM.YY HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtEnd13 = "DD.MM.YY HH24:MI"
	// FormatDateYearAtEnd14 its DDMMYY HH24:MI. Complete date plus hours and minutes.
	FormatDateYearAtEnd14 = "DDMMYY HH24:MI"

	// FormatDateYearAtEnd21 its DD-MM-YY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtEnd21 = "DD-MM-YY HH24:MI:SS"
	// FormatDateYearAtEnd22 its DD/MM/YY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtEnd22 = "DD/MM/YY HH24:MI:SS"
	// FormatDateYearAtEnd23 its DD.MM.YY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtEnd23 = "DD.MM.YY HH24:MI:SS"
	// FormatDateYearAtEnd24 its DDMMYY HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateYearAtEnd24 = "DDMMYY HH24:MI:SS"

	// FormatDateYearAtEnd31 its DD-MM-YY HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateYearAtEnd31 = "DD-MM-YY HH24:MI:SS.FFF"

	// FormatDateYearAtEnd41 its DD-MM-YY HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateYearAtEnd41 = "DD-MM-YY HH24:MI:SS.FFFFFF"

	// FormatDateYearAtEnd51 its DD-MM-YY HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateYearAtEnd51 = "DD-MM-YY HH24:MI:SS.FFFFFFFFF"
)

// Full day format and year at the end
const (
	// FormatDayLongYearAtEnd1 its DD-MM-YYYY. Complete day.
	FormatDayLongYearAtEnd1 = "DD-MM-YYYY"
	// FormatDayLongYearAtEnd2 its DD/MM/YYYY. Complete day.
	FormatDayLongYearAtEnd2 = "DD/MM/YYYY"
	// FormatDayLongYearAtEnd3 its DD.MM.YYYY. Complete day.
	FormatDayLongYearAtEnd3 = "DD.MM.YYYY"
	// FormatDayLongYearAtEnd4 its DDMMYYYY. Complete day.
	FormatDayLongYearAtEnd4 = "DDMMYYYY"
)

// Day format and year at the end
const (
	// FormatDayYearAtEnd1 its DD-MM-YY. Complete day.
	FormatDayYearAtEnd1 = "DD-MM-YY"
	// FormatDayYearAtEnd2 its DD/MM/YY. Complete day.
	FormatDayYearAtEnd2 = "DD/MM/YY"
	// FormatDayYearAtEnd3 its DD.MM.YY. Complete day.
	FormatDayYearAtEnd3 = "DD.MM.YY"
	// FormatDayYearAtEnd4 its DDMMYY. Complete day.
	FormatDayYearAtEnd4 = "DDMMYY"
)

// Collection of days of the week in English
var DaysEn = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// Collection of days of the week in abbreviated form in English
var DaysEnAbbreviated = [...]string{
	"Su",
	"Mo",
	"Tu",
	"We",
	"Th",
	"Fr",
	"Sa",
}

// Collection of days of the week in Russian
var DaysRu = [...]string{
	"Воскресенье",
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
}

// Collection of days of the week in abbreviated form in Russian
var DaysRuAbbreviated = [...]string{
	"Вс",
	"Пн",
	"Вт",
	"Ср",
	"Чт",
	"Пт",
	"Сб",
}

// Collection of months in English
var MonthsEn = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// Collection of months in simple form in Russian
var MonthsRu = [...]string{
	"Январь",
	"Февраль",
	"Март",
	"Апрель",
	"Май",
	"Июнь",
	"Июль",
	"Август",
	"Сентябрь",
	"Октябрь",
	"Ноябрь",
	"Декабрь",
}

// Collection of months in genitive form in Russian
var MonthsRuGenitive = [...]string{
	"Января",
	"Февраля",
	"Марта",
	"Апреля",
	"Мая",
	"Июня",
	"Июля",
	"Августа",
	"Сентября",
	"Октября",
	"Ноября",
	"Декабря",
}

// Collection of months in dative form in Russian
var MonthsRuDative = [...]string{
	"Январю",
	"Февралю",
	"Марту",
	"Апрелю",
	"Маю",
	"Июню",
	"Июлю",
	"Августу",
	"Сентябрю",
	"Октябрю",
	"Ноябрю",
	"Декабрю",
}

// Collection of months in accusative form in Russian
var MonthsRuAccusative = [...]string{
	"Январь",
	"Февраль",
	"Март",
	"Апрель",
	"Май",
	"Июнь",
	"Июль",
	"Август",
	"Сентябрь",
	"Октябрь",
	"Ноябрь",
	"Декабрь",
}

// Collection of months in ablative form in Russian
var MonthsRuAblative = [...]string{
	"Январем",
	"Февралем",
	"Мартом",
	"Апрелем",
	"Маем",
	"Июнем",
	"Июлем",
	"Августом",
	"Сентябрем",
	"Октябрем",
	"Ноябрем",
	"Декабрем",
}

// Collection of months in prepositional form in Russian
var MonthsRuPrepositional = [...]string{
	"Январе",
	"Феврале",
	"Марте",
	"Апреле",
	"Мае",
	"Июне",
	"Июле",
	"Августе",
	"Сентябре",
	"Октябре",
	"Ноябре",
	"Декабре",
}

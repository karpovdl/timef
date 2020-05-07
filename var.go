package timef

// List date time variables
const (
	StampDateLong = "2006-01-02 15:04:05" // 2006-01-02 15:04:05
	StampDate     = "06-01-02 15:04:05"   // 06-01-02 15:04:05
	StampDayLong  = "2006-01-02"          // 2006-01-02

	StampDotDateLong = "2006.01.02 15:04:05" // 2006.01.02 15:04:05
	StampDotDate     = "06.01.02 15:04:05"   // 06.01.02 15:04:05
	StampDotDayLong  = "2006.01.02"          // 2006.01.02

	StampDay  = "06-01-02" // 06-01-02
	StampTime = "15:04:05" // 15:04:05

	StampDotDay = "06.01.02" // 06.01.02

	StampDotDateLongBackwards = "02.01.2006 15:04:05" // 02.01.2006 15:04:05
	StampDotDateBackwards     = "02.01.06 15:04:05"   // 02.01.06 15:04:05
	StampDotDayLongBackwards  = "02.01.2006"          // 02.01.2006
	StampDotDayBackwards      = "02.01.06"            // 02.01.06

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

const (
	// FormatDateLong11 its YYYY-MM-DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLong11 = "YYYY-MM-DD HH24:MI"
	// FormatDateLong12 its YYYY/MM/DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLong12 = "YYYY/MM/DD HH24:MI"
	// FormatDateLong13 its YYYY.MM.DD HH24:MI. Complete date plus hours and minutes.
	FormatDateLong13 = "YYYY.MM.DD HH24:MI"
	// FormatDateLong14 its YYYYMMDD HH24:MI. Complete date plus hours and minutes.
	FormatDateLong14 = "YYYYMMDD HH24:MI"

	// FormatDateLong21 its YYYY-MM-DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLong21 = "YYYY-MM-DD HH24:MI:SS"
	// FormatDateLong22 its YYYY/MM/DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLong22 = "YYYY/MM/DD HH24:MI:SS"
	// FormatDateLong23 its YYYY.MM.DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLong23 = "YYYY.MM.DD HH24:MI:SS"
	// FormatDateLong24 its YYYYMMDD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDateLong24 = "YYYYMMDD HH24:MI:SS"

	// FormatDateLong31 its YYYY-MM-DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLong31 = "YYYY-MM-DD HH24:MI:SS.FFF"
	// FormatDateLong32 its YYYY/MM/DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLong32 = "YYYY/MM/DD HH24:MI:SS.FFF"
	// FormatDateLong33 its YYYY.MM.DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLong33 = "YYYY.MM.DD HH24:MI:SS.FFF"
	// FormatDateLong34 its YYYYMMDD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDateLong34 = "YYYYMMDD HH24:MI:SS.FFF"

	// FormatDateLong41 its YYYY-MM-DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLong41 = "YYYY-MM-DD HH24:MI:SS.FFFFFF"
	// FormatDateLong42 its YYYY/MM/DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLong42 = "YYYY/MM/DD HH24:MI:SS.FFFFFF"
	// FormatDateLong43 its YYYY.MM.DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLong43 = "YYYY.MM.DD HH24:MI:SS.FFFFFF"
	// FormatDateLong44 its YYYYMMDD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDateLong44 = "YYYYMMDD HH24:MI:SS.FFFFFF"

	// FormatDateLong51 its YYYY-MM-DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLong51 = "YYYY-MM-DD HH24:MI:SS.FFFFFFFFF"
	// FormatDateLong52 its YYYY/MM/DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLong52 = "YYYY/MM/DD HH24:MI:SS.FFFFFFFFF"
	// FormatDateLong53 its YYYY.MM.DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLong53 = "YYYY.MM.DD HH24:MI:SS.FFFFFFFFF"
	// FormatDateLong54 its YYYYMMDD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDateLong54 = "YYYYMMDD HH24:MI:SS.FFFFFFFFF"

	// FormatDate11 its YY-MM-DD HH24:MI. Complete date plus hours and minutes.
	FormatDate11 = "YY-MM-DD HH24:MI"
	// FormatDate21 its YY-MM-DD HH24:MI:SS. Complete date plus hours and minutes, seconds.
	FormatDate21 = "YY-MM-DD HH24:MI:SS"
	// FormatDate31 its YY-MM-DD HH24:MI:SS.FFF. Complete date plus hours, minutes, seconds, milliseconds.
	FormatDate31 = "YY-MM-DD HH24:MI:SS.FFF"
	// FormatDate41 its YY-MM-DD HH24:MI:SS.FFFFFF. Complete date plus hours, minutes, seconds, microseconds.
	FormatDate41 = "YY-MM-DD HH24:MI:SS.FFFFFF"
	// FormatDate51 its YY-MM-DD HH24:MI:SS.FFFFFFFFF. Complete date plus hours, minutes, seconds, nanoseconds.
	FormatDate51 = "YY-MM-DD HH24:MI:SS.FFFFFFFFF"

	// FormatDayLong1 its YYYY-MM-DD. Complete day.
	FormatDayLong1 = "YYYY-MM-DD"
	// FormatDayLong2 its YYYY/MM/DD. Complete day.
	FormatDayLong2 = "YYYY/MM/DD"
	// FormatDayLong3 its YYYY.MM.DD. Complete day.
	FormatDayLong3 = "YYYY.MM.DD"
	// FormatDayLong4 its YYYYMMDD. Complete day.
	FormatDayLong4 = "YYYYMMDD"

	// FormatDay1 its YY-MM-DD. Complete day.
	FormatDay1 = "YY-MM-DD"
	// FormatDay2 its YY/MM/DD. Complete day.
	FormatDay2 = "YY/MM/DD"
	// FormatDay3 its YY.MM.DD. Complete day.
	FormatDay3 = "YY.MM.DD"
	// FormatDay4 its YYMMDD. Complete day.
	FormatDay4 = "YYMMDD"

	// FormatDayLongBackwards1 its DD-MM-YYYY. Complete day.
	FormatDayLongBackwards1 = "DD-MM-YYYY"
	// FormatDayLongBackwards2 its DD/MM/YYYY. Complete day.
	FormatDayLongBackwards2 = "DD/MM/YYYY"
	// FormatDayLongBackwards3 its DD.MM.YYYY. Complete day.
	FormatDayLongBackwards3 = "DD.MM.YYYY"
	// FormatDayLongBackwards4 its DDMMYYYY. Complete day.
	FormatDayLongBackwards4 = "DDMMYYYY"

	// FormatDayBackwards1 its DD-MM-YY. Complete day.
	FormatDayBackwards1 = "DD-MM-YY"
	// FormatDayBackwards2 its DD/MM/YY. Complete day.
	FormatDayBackwards2 = "DD/MM/YY"
	// FormatDayBackwards3 its DD.MM.YY. Complete day.
	FormatDayBackwards3 = "DD.MM.YY"
	// FormatDayBackwards4 its DDMMYY. Complete day.
	FormatDayBackwards4 = "DDMMYY"
)

var days = [...]string{
	"Воскресенье",
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
}

var monthsEn = [...]string{
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

var monthsRu = [...]string{
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

// Genitive case of the month
var monthsRuGenitive = [...]string{
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

// Dative case of the month
var monthsRuDative = [...]string{
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

// Accusative case of the month
var monthsRuAccusative = [...]string{
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

// Ablative case of the month
var monthsRuAblative = [...]string{
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

// Prepositional case of the month
var monthsRuPrepositional = [...]string{
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

package timef

import (
	"testing"
	"time"
)

func TestToFormat(t *testing.T) {
	var tests = map[string]struct {
		prototype string
		date      string
		layout    string
		format    string
	}{
		"date_2020-05-20_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "2020-05-20 00:00:00",
			layout:    StampDashDateLongYearAtBegin,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_2020/05/20_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "2020/05/20 00:00:00",
			layout:    StampSlashDateLongYearAtBegin,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_2020.05.20_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "2020.05.20 00:00:00",
			layout:    StampDotDateLongYearAtBegin,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_20200520_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "20200520 00:00:00",
			layout:    StampDateLongYearAtBegin,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_2020-05-20": {
			prototype: "2020-05-20 00:00:00",
			date:      "2020-05-20",
			layout:    StampDashDayLongYearAtBegin,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_20.05.2020": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.2020",
			layout:    StampDotDayLongYearAtEnd,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_20.05.2020_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.2020 00:00:00",
			layout:    StampDotDateLongYearAtEnd,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_20.05.20_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.20 00:00:00",
			layout:    StampDotDateYearAtEnd,
			format:    Format[FormatDateLongYearAtBegin21],
		},
		"date_20.05.20": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.20",
			layout:    StampDotDayYearAtEnd,
			format:    Format[FormatDateLongYearAtBegin21],
		},
	}

	var endl = "\r\n"

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		prototype := test.prototype
		date := test.date
		layout := test.layout
		format := test.format

		if parseDate, err := ToFormat(date, layout, format); err != nil || parseDate == "" || parseDate != prototype {
			if err != nil {
				t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseDate, endl, "ERROR:", err)
			} else {
				t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseDate)
			}
		}
	}
}

func TestToFormatYYYYMMDD(t *testing.T) {
	var tests = map[string]struct {
		prototype string
		datas     []struct {
			date      string
			layout    string
			separator rune
		}
	}{
		"date_2020-05-20": {
			prototype: "2020-05-20",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "2020-05-20 00:00:00",
					layout:    StampDashDateLongYearAtBegin,
					separator: '-',
				},
				{
					date:      "2020-05-20",
					layout:    StampDashDayLongYearAtBegin,
					separator: '-',
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtBegin,
					separator: '-',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: '-',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: '-',
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongYearAtEnd,
					separator: '-',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtBegin,
					separator: '-',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: '-',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: '-',
				},
			},
		},

		"date_2020/05/20": {
			prototype: "2020/05/20",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "2020-05-20 00:00:00",
					layout:    StampDashDateLongYearAtBegin,
					separator: '/',
				},
				{
					date:      "2020-05-20",
					layout:    StampDashDayLongYearAtBegin,
					separator: '/',
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtBegin,
					separator: '/',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: '/',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: '/',
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongYearAtEnd,
					separator: '/',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtEnd,
					separator: '/',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: '/',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: '/',
				},
			},
		},

		"date_2020.05.20": {
			prototype: "2020.05.20",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "2020-05-20 00:00:00",
					layout:    StampDashDateLongYearAtBegin,
					separator: '.',
				},
				{
					date:      "2020-05-20",
					layout:    StampDashDayLongYearAtBegin,
					separator: '.',
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtBegin,
					separator: '.',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: '.',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: '.',
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongYearAtEnd,
					separator: '.',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtEnd,
					separator: '.',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: '.',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: '.',
				},
			},
		},

		"date_20200520": {
			prototype: "20200520",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "2020-05-20 00:00:00",
					layout:    StampDashDateLongYearAtBegin,
					separator: 32,
				},
				{
					date:      "2020-05-20",
					layout:    StampDashDayLongYearAtBegin,
					separator: 32,
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtBegin,
					separator: 32,
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: 32,
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: 32,
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongYearAtEnd,
					separator: 32,
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtEnd,
					separator: 32,
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: 32,
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: 32,
				},
			},
		},
	}
	var endl = "\r\n"
	//fmt.Println([]rune("/-. ")) // Output: [47 45 46 32]

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		prototype := test.prototype

		for _, data := range test.datas {
			date := data.date
			layout := data.layout
			separator := data.separator

			if parseDate, err := ToYYYYMMDD(date, layout, separator); err != nil || parseDate == "" || parseDate != prototype {
				if err != nil {
					t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseDate, endl, "ERROR:", err)
				} else {
					t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseDate)
				}
			}
		}
	}
}

func TestToFormatDDMMYYYY(t *testing.T) {
	var tests = map[string]struct {
		prototype string
		datas     []struct {
			date      string
			layout    string
			separator rune
		}
	}{
		"date_20-05-2020": {
			prototype: "20-05-2020",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "20-05-2020 00:00:00",
					layout:    StampDashDateLongYearAtEnd,
					separator: '-',
				},
				{
					date:      "20-05-2020",
					layout:    StampDashDayLongYearAtEnd,
					separator: '-',
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtEnd,
					separator: '-',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: '-',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: '-',
				},
				{
					date:      "2020.05.20 00:00:00",
					layout:    StampDotDateLongYearAtBegin,
					separator: '-',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtEnd,
					separator: '-',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: '-',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: '-',
				},
			},
		},

		"date_20/05/2020": {
			prototype: "20/05/2020",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "20-05-2020 00:00:00",
					layout:    StampDashDateLongYearAtEnd,
					separator: '/',
				},
				{
					date:      "20-05-2020",
					layout:    StampDashDayLongYearAtEnd,
					separator: '/',
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtEnd,
					separator: '/',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: '/',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: '/',
				},
				{
					date:      "2020.05.20 00:00:00",
					layout:    StampDotDateLongYearAtBegin,
					separator: '/',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtBegin,
					separator: '/',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: '/',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: '/',
				},
			},
		},

		"date_20.05.2020": {
			prototype: "20.05.2020",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "20-05-2020 00:00:00",
					layout:    StampDashDateLongYearAtEnd,
					separator: '.',
				},
				{
					date:      "20-05-2020",
					layout:    StampDashDayLongYearAtEnd,
					separator: '.',
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtEnd,
					separator: '.',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: '.',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: '.',
				},
				{
					date:      "2020.05.20 00:00:00",
					layout:    StampDotDateLongYearAtBegin,
					separator: '.',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtBegin,
					separator: '.',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: '.',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: '.',
				},
			},
		},

		"date_20052020": {
			prototype: "20052020",
			datas: []struct {
				date      string
				layout    string
				separator rune
			}{
				{
					date:      "20-05-2020 00:00:00",
					layout:    StampDashDateLongYearAtEnd,
					separator: 32,
				},
				{
					date:      "20-05-2020",
					layout:    StampDashDayLongYearAtEnd,
					separator: 32,
				},
				{
					date:      "20-05-20",
					layout:    StampDashDayYearAtEnd,
					separator: 32,
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongYearAtEnd,
					separator: 32,
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtEnd,
					separator: 32,
				},
				{
					date:      "2020.05.20 00:00:00",
					layout:    StampDotDateLongYearAtBegin,
					separator: 32,
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateYearAtBegin,
					separator: 32,
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLongYearAtBegin,
					separator: 32,
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayYearAtBegin,
					separator: 32,
				},
			},
		},
	}
	var endl = "\r\n"
	//fmt.Println([]rune("/-. ")) // Output: [47 45 46 32]

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		prototype := test.prototype

		for _, data := range test.datas {
			date := data.date
			layout := data.layout
			separator := data.separator

			if parseDate, err := ToDDMMYYYY(date, layout, separator); err != nil || parseDate == "" || parseDate != prototype {
				if err != nil {
					t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseDate, endl, "ERROR:", err)
				} else {
					t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseDate)
				}
			}
		}
	}
}

func TestTryConvertMonthRuToEn(t *testing.T) {
	tests := map[string]struct {
		prototype string
		months    []string
	}{
		"month_january": {
			prototype: time.January.String(),
			months: []string{
				"январь",
				"Январь",
				"ЯНВАРЬ",
				"января",
				"январю",
				"январем",
				"январЁм",
				"январе",
			},
		},
		"month_february": {
			prototype: time.February.String(),
			months: []string{
				"февраль",
				"Февраль",
				"ФЕВРАЛЬ",
				"февраля",
				"февралю",
				"февралем",
				"февралём",
				"февралЁм",
				"феврале",
			},
		},
	}

	var endl = "\r\n"

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		prototype := test.prototype
		for _, month := range test.months {
			if parseMonth, ok := TryConvertMonthRuToEn(month); !ok || parseMonth == "" || parseMonth != prototype {
				t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "MONTH:", parseMonth)
			}
		}
	}

	if parseMonth, ok := TryConvertMonthRuToEn(""); ok || parseMonth != "" {
		t.Error("ERROR PARSING", endl, "MONTH IS NOT NULL:", parseMonth)
	}

	if parseMonth, ok := TryConvertMonthRuToEn("NONE_MONTH"); ok || parseMonth != "" {
		t.Error("ERROR PARSING", endl, "MONTH IS NOT NULL:", parseMonth)
	}
}

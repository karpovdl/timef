package timef

import (
	"testing"
	"time"
)

func TestToFormat(t *testing.T) {
	tests := map[string]struct {
		prototype string
		date      string
		layout    string
		format    string
	}{
		"date_2020-05-20_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "2020-05-20 00:00:00",
			layout:    StampDateLong,
			format:    GetFormat(FormatDateLong21),
		},
		"date_2020-05-20": {
			prototype: "2020-05-20 00:00:00",
			date:      "2020-05-20",
			layout:    StampDayLong,
			format:    GetFormat(FormatDateLong21),
		},
		"date_20.05.2020": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.2020",
			layout:    StampDotDayLongBackwards,
			format:    GetFormat(FormatDateLong21),
		},
		"date_20.05.2020_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.2020 00:00:00",
			layout:    StampDotDateLongBackwards,
			format:    GetFormat(FormatDateLong21),
		},
		"date_20.05.20_00:00:00": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.20 00:00:00",
			layout:    StampDotDateBackwards,
			format:    GetFormat(FormatDateLong21),
		},
		"date_20.05.20": {
			prototype: "2020-05-20 00:00:00",
			date:      "20.05.20",
			layout:    StampDotDayBackwards,
			format:    GetFormat(FormatDateLong21),
		},
	}

	endl := "\r\n"

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
	tests := map[string]struct {
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
					layout:    StampDateLong,
					separator: '-',
				},
				{
					date:      "2020-05-20",
					layout:    StampDayLong,
					separator: '-',
				},
				{
					date:      "20-05-20",
					layout:    StampDay,
					separator: '-',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLong,
					separator: '-',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDay,
					separator: '-',
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongBackwards,
					separator: '-',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateBackwards,
					separator: '-',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongBackwards,
					separator: '-',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayBackwards,
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
					layout:    StampDateLong,
					separator: '/',
				},
				{
					date:      "2020-05-20",
					layout:    StampDayLong,
					separator: '/',
				},
				{
					date:      "20-05-20",
					layout:    StampDay,
					separator: '/',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLong,
					separator: '/',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDay,
					separator: '/',
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongBackwards,
					separator: '/',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateBackwards,
					separator: '/',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongBackwards,
					separator: '/',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayBackwards,
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
					layout:    StampDateLong,
					separator: '.',
				},
				{
					date:      "2020-05-20",
					layout:    StampDayLong,
					separator: '.',
				},
				{
					date:      "20-05-20",
					layout:    StampDay,
					separator: '.',
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLong,
					separator: '.',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDay,
					separator: '.',
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongBackwards,
					separator: '.',
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateBackwards,
					separator: '.',
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongBackwards,
					separator: '.',
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayBackwards,
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
					layout:    StampDateLong,
					separator: 32,
				},
				{
					date:      "2020-05-20",
					layout:    StampDayLong,
					separator: 32,
				},
				{
					date:      "20-05-20",
					layout:    StampDay,
					separator: 32,
				},
				{
					date:      "2020.05.20",
					layout:    StampDotDayLong,
					separator: 32,
				},
				{
					date:      "20.05.20",
					layout:    StampDotDay,
					separator: 32,
				},
				{
					date:      "20.05.2020 00:00:00",
					layout:    StampDotDateLongBackwards,
					separator: 32,
				},
				{
					date:      "20.05.20 00:00:00",
					layout:    StampDotDateBackwards,
					separator: 32,
				},
				{
					date:      "20.05.2020",
					layout:    StampDotDayLongBackwards,
					separator: 32,
				},
				{
					date:      "20.05.20",
					layout:    StampDotDayBackwards,
					separator: 32,
				},
			},
		},
	}

	endl := "\r\n"
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

	endl := "\r\n"

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		prototype := test.prototype
		for _, month := range test.months {
			if parseMonth, ok := TryConvertMonthRuToEn(month); !ok || parseMonth == "" || parseMonth != prototype {
				t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "DATE:", parseMonth)
			}
		}
	}
}

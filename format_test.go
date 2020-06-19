package timef

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		prototype string
		format    string
	}{
		"f_YYYY/MM/DD_HH24:MI": {
			prototype: strings.Join([]string{LongYear, ZeroMonth, ZeroDay}, "/") + " " + strings.Join([]string{Hour, ZeroMinute}, ":"),
			format:    FormatDateLongYearAtBegin12,
		},
	}

	var endl = "\r\n"

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		prototype := test.prototype
		format := test.format

		if parseFormat := Format[format]; parseFormat == "" || parseFormat != prototype {
			t.Error("ERROR PARSING", endl, "PROTOTYPE:", prototype, endl, "FORMAT:", parseFormat)
		}
	}
}

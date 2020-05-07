package timef

import (
	"testing"
)

func TestVar(t *testing.T) {
	if LongYear != "2006" {
		t.Error()
	}
}

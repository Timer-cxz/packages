package trace

import (
	"testing"
)

func TestRecordString(t *testing.T) {
	var r1 *Record
	strRNil := "[nil-record]"
	if r1.String() != strRNil {
		t.Errorf("string of nil Record is not \"%s\"", strRNil)
	}

	r2 := &Record{Name: "trace", File: "cxz", Line: 34}
	r2Str := r2.String()
	destStr := "cxz:34 trace"
	if r2Str != destStr {
		t.Errorf("Record string: \"%s\" != \"%s\"", r2Str, destStr)
	}
}

func TestCaller(t *testing.T) {
	r := Caller(0)
	t.Logf("%s", r.String())

	r = Caller(1)
	t.Logf("%s", r.String())

	r = Caller(2)
	t.Logf("%s", r.String())
}

func TestTrace(t *testing.T) {
	s := Trace()
	t.Logf("%s", s.String())
}

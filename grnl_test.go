package main

import (
	"testing"
	"time"
)

type parsePair struct {
	str     string
	entries []*Entry
}

func TestParse(t *testing.T) {
	tests := []parsePair{
		{
			"yesterday: Called in sick. Used the time to clean the house and spent 4h on writing my book.",
			[]*Entry{{time.Now(), "Called in sick", "Used the time to clean the house and spent 4h on writing my book."}},
		},
	}

	for _, tst := range tests {
		err, entries := parse(tst.str)
		if err != nil {
			t.Errorf("Got an error parsing '%s': %+v", tst.str, err)
		}

		if !EntryEqual(tst.entries, entries) {
			t.Errorf("Got %+v while parsing '%s', expected %+v.", entries, tst.str, tst.entries)
		}
	}
}

func EntryEqual(a, b []*Entry) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

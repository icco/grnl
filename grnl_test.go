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

		if len(tst.entries) != len(entries) {
			t.Errorf("Got %d entries, expected %d, while parsing '%s'.", len(entries), len(tst.entries), tst.str)
		} else {
			for i, v := range tst.entries {
				if v != entries[i] {
					t.Errorf("%+v != %+v.", entries[i], tst.entries[i])
				}
			}
		}
	}
}

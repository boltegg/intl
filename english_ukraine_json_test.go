package intl

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"golang.org/x/text/language"
)

type enUATestFile struct {
	Locale string `json:"locale"`
	Tests  []struct {
		Type      string `json:"type"`
		Scenarios []struct {
			Expected string `json:"expected"`
			Pattern  string `json:"pattern"`
			Value    string `json:"value"`
		} `json:"scenarios"`
	} `json:"tests"`
}

func TestEnglishUkraineFromTestData(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("internal/test/testdata/en_UA.json")
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}

	var file enUATestFile
	if err := json.Unmarshal(data, &file); err != nil {
		t.Fatalf("unmarshal testdata: %v", err)
	}

	locale := language.MustParse(file.Locale)

	for _, group := range file.Tests {
		if group.Type != "DateFormatter::Skeleton" {
			continue
		}
		for _, sc := range group.Scenarios {
			ms, err := strconv.ParseInt(sc.Value, 10, 64)
			if err != nil {
				t.Fatalf("parse value %q: %v", sc.Value, err)
			}
			date := time.Unix(0, ms*int64(time.Millisecond))

			opts, err := ParseLayout(sc.Pattern)
			if err != nil {
				t.Logf("skip pattern %q: %v", sc.Pattern, err)
				continue
			}

			got := NewDateTimeFormat(locale, opts).Format(date)
			if got != sc.Expected {
				t.Errorf("pattern %q value %v: want %q got %q", sc.Pattern, date, sc.Expected, got)
			}
		}
	}
}

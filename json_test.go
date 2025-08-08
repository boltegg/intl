package intl

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/language"
)

type jsonTestFile struct {
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

func TestFromTestData(t *testing.T) {
	t.Parallel()

	entries, err := os.ReadDir("internal/test/testdata")
	if err != nil {
		t.Fatalf("read dir: %v", err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		entryName := entry.Name()
		t.Run(entryName, func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(filepath.Join("internal/test/testdata", entryName))
			if err != nil {
				t.Fatalf("read testdata: %v", err)
			}

			var file jsonTestFile
			if err := json.Unmarshal(data, &file); err != nil {
				t.Fatalf("unmarshal testdata: %v", err)
			}

			locale := language.MustParse(file.Locale)
			if base, _ := locale.Base(); base.String() == "es" {
				t.Skip("locale not supported")
			}

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
					expected := strings.ReplaceAll(sc.Expected, "\u202f", " ")
					gotNormalized := strings.ReplaceAll(got, "\u202f", " ")
					if gotNormalized != expected {
						t.Errorf("pattern %q value %v: want %q got %q", sc.Pattern, date, sc.Expected, got)
					}
				}
			}
		})
	}
}

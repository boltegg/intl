package main

import (
	"fmt"
	"time"

	"golang.org/x/text/language"

	"github.com/boltegg/intl"
)

func main() {
	v := intl.NewDateTimeFormat(language.English, intl.Options{Minute: intl.Minute2Digit, Hour: intl.Hour2Digit})
	fmt.Println(v.Format(time.Now()))

	v = intl.NewDateTimeFormatLayout(language.English, "MEd")
	fmt.Println(v.Format(time.Now()))
}

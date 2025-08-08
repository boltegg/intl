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

	var err error
	v, err = intl.NewDateTimeFormatLayout(language.English, "MEd")
	if err != nil {
		panic(err)
	}
	fmt.Println(v.Format(time.Now()))
}

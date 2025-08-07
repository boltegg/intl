package main

import (
	"fmt"
	"time"

	"github.com/boltegg/intl"
	"golang.org/x/text/language"
)

func main() {
	v := intl.NewDateTimeFormat(language.Latvian, intl.Options{Year: intl.YearNumeric})

	fmt.Println(v.Format(time.Now()))
}

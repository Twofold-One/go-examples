package basic

import (
	"fmt"
	"log"
	"time"
)

func DatesAndTimeExample() {
	// to get current time
	now := time.Now()
	fmt.Printf("%s\n", now)

	// The “Unix Epoch” is the number of seconds elapsed since 1 January 1970 00:00:00 UTC.
	log.Println(now.Unix())

	// format the time
	// The format is "Mon Jan 2 ". This is because the format in Go is defined based on a reference date which is :
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// Monday
	// January 2 from the year 2006 at 15:04:05
	// MST means Mountain Standard Time. This is the timezone (UTC minus 7 hours)
	fmt.Println(now.Format("Mon Jan 2"))
	// specific formats
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC1123))

	// date in location
	// the function time.LoadLocation takes a string as parameter, which represents a timezone name from the IANA1 Time Zone Database2
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	nowNYC := now.In(loc)
	fmt.Printf("%s\n", nowNYC)

	// parse date/time
	timeToParse := "2019-02-15T07:33-05:00"
	t, err := time.Parse("2006-01-02T15:04-07:00", timeToParse)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	
}
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

	// get the time elapsed between two points using method Sub
	start := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed)

	// check if time is between two points
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	firstJanuary1980 := time.Date(1980, 1, 1, 0, 0, 0, 0, location)

	timeToParse = "2019-02-15T07:33-02:00"
	t1, err := time.Parse("2006-01-02T15:04-07:00", timeToParse)
	if err != nil {
		panic(err)
	}

	now = time.Now()
	if t1.After(firstJanuary1980) && t1.Before(now) {
		fmt.Println(t, " is between", firstJanuary1980, " and ", now)
	} else {
		fmt.Println("not between")
	}

	// adding days, hours, minutes and seconds to a time
	now = time.Now()
	fmt.Println(now)

	// +1 year, 6 months and 10 days
	now = now.AddDate(1, 6, 10)
	fmt.Println(now)

	now.Add(time.Second * 5)
	now.Add(time.Minute * 5)
	now.Add(time.Hour * 5)

	// iterate over time
	st, err := time.Parse("2006-01-02", "2019-02-19")
	if err != nil {
		panic(err)
	}
	en, err := time.Parse("2006-01-02", "2020-07-17")
	if err != nil {
		panic(err)
	}
	for i := st; i.Unix() < en.Unix(); i = i.AddDate(0, 0, 1) {
		fmt.Println(i.Format(time.RFC3339))
	}
}
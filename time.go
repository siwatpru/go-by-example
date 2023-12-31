package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

  // NOW!
	now := time.Now()
	p(now)

  // Parse time
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

  // Time components
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

  // Time comparison
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

  // Time difference
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

  // Add time
	p(then.Add(diff))
	p(then.Add(-diff))
}

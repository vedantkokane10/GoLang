package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now() // to get current time
	fmt.Println("Present Time: ", presentTime)
	// Present Time:  2026-03-11 21:03:48.540799 +0530 IST m=+0.000219834
	//				   Date		  Time (24 hrs)        Zone

	// to get time in custom format (Inside Format() function these are default values so we can't change them, we can only keep necessary stuff that we want)
	fmt.Println("Present Time: ", presentTime.Format("01-02-2006"))
	// Present Time:  03-11-2026

	fmt.Println("Present Time: ", presentTime.Format("01-02-2006 Monday"))
	// Present Time:  03-11-2026 Wednesday

	fmt.Println("Present Time: ", presentTime.Format("01-02-2006 Monday"))
	// Present Time:  03-11-2026 Wednesday


	fmt.Println("Present Time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))
	// Present Time:  03-11-2026 21:08:17 Wednesday

	fmt.Println(presentTime.Format("Monday"))
	// Wednesday

	// custom date
	customDate := time.Date(2024, time.April, 3, 8, 58, 0, 0, time.UTC);

	fmt.Println(customDate)


	// Many things to explore inside time package -> https://pkg.go.dev/time
}

//
// 改编自https://play.golang.org/p/sWZhYeKsGQ
//

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	defaultFormat := "2006-01-02 15:04:05 PM -07:00 Jan Mon MST"

	formats := []map[string]string{
		{"format": "2006", "description": "Year"},
		{"format": "06", "description": "Year"},

		{"format": "01", "description": "Month"},
		{"format": "1", "description": "Month"},
		{"format": "Jan", "description": "Month"},
		{"format": "January", "description": "Month"},

		{"format": "02", "description": "Day"},
		{"format": "2", "description": "Day"},

		{"format": "Mon", "description": "Week day"},
		{"format": "Monday", "description": "Week day"},

		{"format": "03", "description": "Hours"},
		{"format": "3", "description": "Hours"},
		{"format": "15", "description": "Hours"},

		{"format": "04", "description": "Minutes"},
		{"format": "4", "description": "Minutes"},

		{"format": "05", "description": "Seconds"},
		{"format": "5", "description": "Seconds"},

		{"format": "PM", "description": "AM or PM"},

		{"format": ".000", "description": "Miliseconds"},
		{"format": ".000000", "description": "Microseconds"},
		{"format": ".000000000", "description": "Nanoseconds"},

		{"format": "-0700", "description": "Timezone offset"},
		{"format": "-07:00", "description": "Timezone offset"},
		{"format": "Z0700", "description": "Timezone offset"},
		{"format": "Z07:00", "description": "Timezone offset"},

		{"format": "MST", "description": "Timezone"}}

	fmt.Printf("\n\n%s \n\n", now.Format(defaultFormat))
	for _, f := range formats {
		fmt.Printf("%-15s | %-12s | %-12s \n", f["description"], f["format"], now.Format(f["format"]))
	}
	fmt.Printf("%-15s + %-12s + %12s \n", strings.Repeat("-", 15), strings.Repeat("-", 12), strings.Repeat("-", 12))
}

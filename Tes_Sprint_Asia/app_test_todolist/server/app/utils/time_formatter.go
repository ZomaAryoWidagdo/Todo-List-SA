package utils

import (
	"fmt"
	"time"
)

func StringToTime(date *string) *time.Time {
	if date == nil || *date == "" {
		return nil
	}

	layout := "2006-01-02 15:04:05 -0700"

	parsedTime, err := time.Parse(layout, *date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}

	return &parsedTime
}

package controller

import (
	"fmt"
	"time"
)

func Get_date() string {
	// Set the time zone to "Asia/Jakarta"
	TimeZone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		// Handle error if timezone loading fails
		return fmt.Sprintf("Error loading timezone: %s", err)
	}

	// Get the current time in the Indonesian timezone
	currentTime := time.Now().In(TimeZone)

	// Format the time as per your requirement (e.g., RFC1123)
	formattedTime := currentTime.Format(time.RFC1123)

	return formattedTime
}

package controller

import (
	"fmt"
	"os"
	"time"
)

func Get_Date() (time.Time, error) {
	// Set the time zone to based on setting in env
	location := os.Getenv("Location")
	timeZone, err := time.LoadLocation(location)

	if err != nil {
		// Handle error if timezone loading fails
		fmt.Printf("Error loading timezone: %s\n", err)
		// Return the current time in UTC as a fallback
		return time.Now().UTC(), err
	}

	// Get the current time in the Indonesian timezone
	currentTime := time.Now().In(timeZone)

	return currentTime, nil
}

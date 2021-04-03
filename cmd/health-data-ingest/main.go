package main

import (
	"os"
)

func main() {
	switch os.Getenv("DEVICE_TYPE") {
	case "FITBIT":
		triggerFBEvents()
	case "STRAVA":
		triggerStravaEvents()
	case "RUNKEEPER":
		triggerRKEvents()
	}
}

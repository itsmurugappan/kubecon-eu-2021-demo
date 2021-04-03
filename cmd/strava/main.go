package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/itsmurugappan/kubecon/pkg/ce"
	"github.com/itsmurugappan/kubecon/pkg/types"
)

const (
	source = "https://www.strava.com"
)

func main() {
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	h, err := cloudevents.NewHTTPReceiveHandler(ctx, p, processCE)
	if err != nil {
		log.Fatalf("failed to create handler: %s", err.Error())
	}

	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}))

	http.ListenAndServe(":8080", nil)
}

func processCE(ctx context.Context, event cloudevents.Event) {
	log.Printf("cloud event received %v\n", event)
	var data []types.Strava
	if err := event.DataAs(&data); err != nil {
		log.Printf("Error marshaling data %s\n", err)
		return
	}
	log.Printf("data %v\n", data)

	//make ce client
	c := ce.CEClient()

	for _, act := range data {
		t, _ := time.Parse(time.RFC3339, act.StartDate)
		metaDetails := types.Meta{
			Name: fmt.Sprintf("StravaAthlete-%d", act.Athlete.ID),
			Date: t.Format("2006-01-02"),
		}
		log.Printf("meta %v\n", metaDetails)

		//record id
		idEvent := ce.GetCE(types.ID_Type, source)
		idEvent.SetData("application/json", metaDetails)
		c.SendMessage(idEvent)

		var actType string
		switch act.ActivityName {
		case "Elliptical":
			actType = types.Elliptical_Type
		case "Walk":
			actType = types.Walk_Type
		case "Run":
			actType = types.Run_Type
		case "Ride":
			actType = types.Ride_Type
		default:
			continue
		}
		actData := types.Activity{
			ActivityDuration: act.Duration,
			ID:               metaDetails,
		}
		actEvent := ce.GetCE(actType, source)
		actEvent.SetData("application/json", actData)
		c.SendMessage(actEvent)
	}
}

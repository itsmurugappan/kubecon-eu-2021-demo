package main

import (
	"context"
	"encoding/json"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/itsmurugappan/kubecon/pkg/ce"
	"github.com/itsmurugappan/kubecon/pkg/types"
)

const (
	source = "https://api.fitbit.com"
)

func main() {
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClientObserved(p, cloudevents.WithTracePropagation)
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("will listen on :8080\n")
	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, processCE))
}

func processCE(ctx context.Context, event cloudevents.Event) {
	log.Printf("cloud event received %v\n", event)
	var data types.Fitbit
	if err := event.DataAs(&data); err != nil {
		log.Printf("Error marshaling data %s\n", err)
		return
	}

	var metaText string
	var metaDetails types.Meta

	if err := event.ExtensionAs(types.Meta_Extension, &metaText); err != nil {
		log.Printf("Error getting extension %s\n", err)
		return
	}

	if err := json.Unmarshal([]byte(metaText), &metaDetails); err != nil {
		log.Printf("Error getting extension %s\n", err)
		return
	}

	//make ce client
	c := ce.CEClient()

	//record id
	idEvent := ce.GetCE(types.ID_Type, source)
	idEvent.SetData("application/json", metaDetails)
	c.SendMessage(idEvent)

	//record steps
	if data.Summary.Steps > 0 {
		stepsData := types.Steps{
			Steps: data.Summary.Steps,
			ID:    metaDetails,
		}

		stepsEvent := ce.GetCE(types.Step_Type, source)
		stepsEvent.SetData("application/json", stepsData)
		c.SendMessage(stepsEvent)
	}

	for _, act := range data.Activities {
		var actType string
		switch act.ActivityName {
		case "Elliptical":
			actType = types.Elliptical_Type
		case "Walk":
			actType = types.Walk_Type
		case "Run":
			actType = types.Run_Type
		default:
			continue
		}
		actData := types.Activity{
			ActivityDuration: act.Duration / 1000,
			ID:               metaDetails,
		}
		actEvent := ce.GetCE(actType, source)
		actEvent.SetData("application/json", actData)
		c.SendMessage(actEvent)
	}
}

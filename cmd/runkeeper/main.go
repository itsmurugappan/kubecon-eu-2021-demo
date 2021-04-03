package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/itsmurugappan/kubecon/pkg/ce"
	"github.com/itsmurugappan/kubecon/pkg/types"
)

const (
	source = "https://www.runkeeper.com"
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
	// event.SetExtension(types.Meta_Extension, "def")
	log.Printf("cloud event received %v\n", event)
	var data types.Runkeeper
	if err := event.DataAs(&data); err != nil {
		log.Printf("Error marshaling data %s\n", err)
		return
	}
	log.Printf("data %v\n", data)

	var name string

	if err := event.ExtensionAs(types.Meta_Extension, &name); err != nil {
		log.Printf("Error getting extension %s\n", err)
		return
	}

	//make ce client
	c := ce.CEClient()

	for _, act := range strings.Split(data.Data, "\n") {
		if act == "EOF" {
			log.Println("reached EOF")
			return
		}

		record := (strings.Split(act, ","))
		metaDetails := types.Meta{
			Name: name,
			Date: (strings.Split(record[1], " "))[0],
		}
		log.Printf("meta %v\n", metaDetails)

		//record id
		idEvent := ce.GetCE(types.ID_Type, source)
		idEvent.SetData("application/json", metaDetails)
		c.SendMessage(idEvent)

		var actType string
		switch record[2] {
		case "Walking":
			actType = types.Walk_Type
		case "Running":
			actType = types.Run_Type
		default:
			continue
		}
		actData := types.Activity{
			ActivityDuration: getDurationSeconds(record[5]),
			ID:               metaDetails,
		}
		actEvent := ce.GetCE(actType, source)
		actEvent.SetData("application/json", actData)
		c.SendMessage(actEvent)
	}
}

func getDurationSeconds(dur string) int {
	durArray := strings.Split(dur, ":")
	var durationseconds int
	switch len(durArray) {
	case 1:
		durationseconds = getIntDuration(durArray[0])
	case 2:
		durationseconds = getIntDuration(durArray[0])*60 + getIntDuration(durArray[1])
	case 3:
		durationseconds = getIntDuration(durArray[0])*3600 + getIntDuration(durArray[1])*60 + getIntDuration(durArray[2])
	}
	return durationseconds
}

func getIntDuration(dur string) int {
	durValue, _ := strconv.Atoi(dur)
	return durValue
}

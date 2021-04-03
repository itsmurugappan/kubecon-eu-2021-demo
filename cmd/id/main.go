package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/itsmurugappan/kubecon/pkg/db"
	"github.com/itsmurugappan/kubecon/pkg/types"
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
	var data types.Meta
	err := event.DataAs(&data)
	if err != nil {
		log.Printf("Error marshaling data %s\n", err)
		return
	}
	db.RecordIDData(data.Name, data.Date)
}

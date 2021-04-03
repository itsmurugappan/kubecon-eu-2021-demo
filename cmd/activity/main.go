package main

import (
	"context"
	"log"
	"net/http"
	"os"

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
	var data types.Activity
	err := event.DataAs(&data)
	if err != nil {
		log.Printf("Error marshaling data %s\n", err)
		return
	}
	db.RecordActivityData(data.ID.Name, data.ID.Date, os.Getenv("Activity_Name"), data.ActivityDuration)
}

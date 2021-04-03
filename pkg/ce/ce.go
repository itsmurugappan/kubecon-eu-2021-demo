package ce

import (
	"context"
	"log"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"go.opencensus.io/plugin/ochttp"
	"knative.dev/pkg/tracing/propagation/tracecontextb3"
)

type ceClient struct {
	cloudEventsClient cloudevents.Client
	ctx               context.Context
}

func CEClient() *ceClient {
	p, err := cloudevents.NewHTTP(
		cloudevents.WithTarget(os.Getenv("K_SINK")),
		cloudevents.WithRoundTripper(&ochttp.Transport{
			Propagation: tracecontextb3.TraceContextEgress,
		}))
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p, cloudevents.WithUUIDs(), cloudevents.WithTimeNow())
	if err != nil {
		log.Fatalf("failed to create client: %s", err.Error())
	}

	return &ceClient{c, context.Background()}
}

func (c *ceClient) SendMessage(e cloudevents.Event) {
	if res := c.cloudEventsClient.Send(c.ctx, e); !cloudevents.IsACK(res) {
		log.Printf("failed to send cloudevent: %v\n", res)
	}
}

func GetCE(eventType, source string) cloudevents.Event {
	event := cloudevents.NewEvent("1.0")
	event.SetType(eventType)
	event.SetSource(source)
	event.SetID(uuid.New().String())
	return event
}

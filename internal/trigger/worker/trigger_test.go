package worker

import (
	// standard libraries
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	// third-party libraries
	ce "github.com/cloudevents/sdk-go/v2"

	// first-party libraries
	eb "github.com/linkall-labs/eventbus-go"
	"github.com/linkall-labs/eventbus-go/pkg/discovery"
	"github.com/linkall-labs/eventbus-go/pkg/discovery/record"
	"github.com/linkall-labs/eventbus-go/pkg/inmemory"

	// this project
	"github.com/linkall-labs/vanus/internal/primitive"
	"github.com/linkall-labs/vanus/internal/trigger/consumer"
	"github.com/linkall-labs/vanus/internal/trigger/info"
	"github.com/linkall-labs/vanus/internal/trigger/storage"
)

func Test_e2e(t *testing.T) {
	tg := NewTrigger(&primitive.Subscription{
		ID:               "test",
		Source:           "human",
		Types:            []string{"aaa"},
		Config:           map[string]string{},
		Filters:          []*primitive.SubscriptionFilter{{Exact: map[string]string{"type": "none"}}},
		Sink:             "http://localhost:18080",
		Protocol:         "vanus",
		ProtocolSettings: nil,
	}, consumer.NewEventLogOffset("test", storage.NewFakeStorage()))
	emit := 0
	pre := 0
	go func() {
		for {
			time.Sleep(time.Second)
			cur := emit
			t.Logf("%v TPS: %d", time.Now(), cur-pre)
			pre = cur
		}
	}()

	ebVRN := "vanus+local:eventbus:example"
	elVRN := "vanus+local:eventlog+inmemory:1?keepalive=true"
	br := &record.EventBus{
		VRN: ebVRN,
		Logs: []*record.EventLog{
			{
				VRN:  elVRN,
				Mode: record.PremWrite | record.PremRead,
			},
		},
	}
	inmemory.UseInMemoryLog("inmemory")
	ns := inmemory.UseNameService("vanus+local")
	// register metadata of eventbus
	vrn, err := discovery.ParseVRN(ebVRN)
	if err != nil {
		panic(err.Error())
	}
	ns.Register(vrn, br)

	go func() {
		w, err := eb.OpenBusWriter(ebVRN)
		if err != nil {
			t.Fatal(err)
		}

		// FIXME
		time.Sleep(10 * time.Second)

		for i := 0; i < 10000000000; i++ {
			// Create an Event.
			event := ce.NewEvent()
			event.SetID(fmt.Sprintf("%d", i))
			event.SetSource("example/uri")
			event.SetType("none")
			event.SetData(ce.ApplicationJSON, map[string]string{"hello": "world", "type": "none"})

			_, err = w.Append(context.Background(), &event)
			if err != nil {
				t.Log(err)
			}
		}

		w.Close()
	}()
	go func() {
		ls, err := eb.LookupReadableLogs(context.Background(), ebVRN)
		if err != nil {
			t.Fatal(err)
		}

		r, err := eb.OpenLogReader(ls[0].VRN)
		if err != nil {
			t.Fatal(err)
		}

		_, err = r.Seek(context.Background(), 0, io.SeekStart)
		if err != nil {
			t.Fatal(err)
		}

		for {
			events, err := r.Read(context.Background(), 5)
			if err != nil {
				t.Fatal(err)
			}

			if len(events) == 0 {
				time.Sleep(time.Second)
				continue
			}

			for _, e := range events {
				tg.EventArrived(context.Background(), &info.EventRecord{
					Event: e,
				})
				emit++
			}
		}

		r.Close()
	}()
	receive := 0
	receivePre := 0
	go http.ListenAndServe(":18080", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		receive++
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
		}
		var _ = string(body)
	}))
	go func() {
		for {
			time.Sleep(time.Second)
			cur := receive
			t.Logf("%v RECEIVE TPS: %d", time.Now(), cur-receivePre)
			receivePre = cur
		}
	}()
	tg.Start(context.Background())

	time.Sleep(time.Hour)
	tg.Stop()
}

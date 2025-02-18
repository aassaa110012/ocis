package service_test

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/ocis-pkg/store"
	ehsvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/eventhistory/v0"
	"github.com/owncloud/ocis/v2/services/eventhistory/pkg/config"
	"github.com/owncloud/ocis/v2/services/eventhistory/pkg/service"
	microevents "go-micro.dev/v4/events"
	microstore "go-micro.dev/v4/store"
)

var _ = Describe("EventHistoryService", func() {
	var (
		cfg = &config.Config{}

		eh  *service.EventHistoryService
		bus testBus
		sto microstore.Store
	)

	BeforeEach(func() {
		var err error
		sto = store.Create()
		bus = testBus(make(chan events.Event))
		eh, err = service.NewEventHistoryService(cfg, bus, sto, log.Logger{})
		Expect(err).ToNot(HaveOccurred())

	})

	AfterEach(func() {
		close(bus)
	})

	It("Records events, stores them and allows to retrieve them", func() {
		id := bus.Publish(events.UploadReady{})

		// service will store eventually
		time.Sleep(500 * time.Millisecond)

		resp := &ehsvc.GetEventsResponse{}
		err := eh.GetEvents(context.Background(), &ehsvc.GetEventsRequest{Ids: []string{id}}, resp)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp).ToNot(BeNil())

		Expect(len(resp.Events)).To(Equal(1))
		Expect(resp.Events[0].Id).To(Equal(id))

	})
})

type testBus chan events.Event

func (tb testBus) Consume(_ string, _ ...microevents.ConsumeOption) (<-chan microevents.Event, error) {
	ch := make(chan microevents.Event)
	go func() {
		for ev := range tb {
			b, _ := json.Marshal(ev.Event)
			ch <- microevents.Event{
				Payload: b,
				Metadata: map[string]string{
					events.MetadatakeyEventID:   ev.ID,
					events.MetadatakeyEventType: ev.Type,
				},
			}
		}
	}()
	return ch, nil
}

func (tb testBus) Publish(e interface{}) string {
	ev := events.Event{
		ID:    uuid.New().String(),
		Type:  reflect.TypeOf(e).String(),
		Event: e,
	}

	tb <- ev
	return ev.ID
}

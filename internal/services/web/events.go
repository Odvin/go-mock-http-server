package web

import (
	"fmt"
	"github.com/Odvin/go-mock-http-server/pkg/mediator"
	"net/http"
	"time"
)

func sendCompanyUpdateEvent(data any) error {
	fmt.Println(data)
	return nil
}

var ps = mediator.GetPubSub()

func (web *Web) events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")

	UpdateCompanyChanel, _ := ps.Subscribe("UpdateCompany")

	for data := range UpdateCompanyChanel {
		content := fmt.Sprintf("data: %s\n\n", data)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		time.Sleep(time.Second)
	}
}

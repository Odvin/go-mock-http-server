package web

import (
	"fmt"
	"net/http"

	"github.com/Odvin/go-mock-http-server/pkg/mediator"
)

var ps = mediator.GetPubSub()

func (hs *HttpServer) events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")

	messages, unsubscribe := ps.Subscribe("UpdateCompany")
	defer unsubscribe()

	for data := range messages {
		content := fmt.Sprintf("data: %s\n\n", data)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
	}
}

package web

import (
	"fmt"
	"net/http"
)

func (hs *HttpServer) events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	messages, unsubscribe := ps.Subscribe("UpdateCompany")
	defer unsubscribe()

	for {
		select {
		case data, ok := <-messages:
			if !ok {
				_, _ = w.Write([]byte("event: close\ndata: stream finished\n\n"))
				flusher.Flush()
				return
			}

			content := fmt.Sprintf("data: %s\n\n", data)
			_, _ = w.Write([]byte(content))
			flusher.Flush()

		case <-r.Context().Done():
			fmt.Println("Client disconnected from /events")
			return

		case <-hs.ctx.Done():
			fmt.Println("Server shutting down, closing SSE stream")
			_, _ = w.Write([]byte("event: close\ndata: server shutting down\n\n"))
			flusher.Flush()
			return
		}
	}
}

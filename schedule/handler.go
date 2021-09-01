package schedule

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/wevnasc/hermes/server"
	"github.com/wevnasc/hermes/storage"
)

type Handler struct {
	*server.Middlewares
	ctrl *controller
}

func (h *Handler) scheduleHandler(rw http.ResponseWriter, req *http.Request) {
	type request struct {
		To          []string `json:"to"`
		Template    string   `json:"template"`
		ScheduledTo string   `json:"scheduled_to"`
	}

	var body request

	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		server.NewError("error to parse body", http.StatusBadRequest).Json(rw)
		return
	}

	scheduleTo, err := time.Parse(time.RFC3339, body.ScheduledTo)

	if err != nil {
		server.NewError("invalid date format", http.StatusBadRequest).Json(rw)
		return
	}

	schedule := &schedule{
		to:          body.To,
		template:    body.Template,
		scheduledTo: scheduleTo,
	}

	if err := h.ctrl.schedule(schedule); err != nil {
		err.Json(rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/schedules", h.Resource(h.scheduleHandler, []string{http.MethodPost}))
}

func NewHandler(logger *log.Logger) *Handler {
	ctrl := newController(
		newMemoDB(),
		storage.NewLocal("templates"),
	)
	return &Handler{
		&server.Middlewares{Log: logger},
		ctrl,
	}
}

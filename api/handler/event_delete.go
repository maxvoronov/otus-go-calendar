package handler

import (
	"context"
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type eventDeleteRequest struct {
	ID uuid.UUID
}

// EventDeleteHandler Delete event from storage if exists
// POST /delete_event
// Fields:
// 	 id: UUID
func (h *Handler) EventDeleteHandler(req *http.Request) APIResponse {
	ctx := context.Background()
	data := &eventDeleteRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	event, err := h.Storage.GetByID(ctx, data.ID.String())
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	if event == nil {
		return h.Error(http.StatusNotFound, errors.New("event not found"))
	}

	if err := h.Storage.Remove(ctx, event); err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.sendJSON(http.StatusAccepted, nil)
}

func (data *eventDeleteRequest) parse(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	eventID := req.FormValue("id")
	if eventID == "" {
		return errors.New("event ID is required")
	}

	var err error
	if data.ID, err = uuid.FromString(eventID); err != nil {
		return err
	}

	return nil
}

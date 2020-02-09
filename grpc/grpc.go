package grpc

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	eventproto "github.com/maxvoronov/otus-go-calendar/grpc/proto"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create Method for creating new event
func (serv *server) Create(_ context.Context, req *eventproto.EventCreateRequest) (*eventproto.EventCreateResponse, error) {
	event, err := serv.Calendar.CreateEvent(
		req.GetTitle(),
		time.Unix(req.DateFrom.GetSeconds(), int64(req.DateFrom.GetNanos())),
		time.Unix(req.DateTo.GetSeconds(), int64(req.DateTo.GetNanos())),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &eventproto.EventCreateResponse{Event: convertToProtoEvent(event)}, nil
}

// Update Method for updating event if exists
func (serv *server) Update(_ context.Context, req *eventproto.EventUpdateRequest) (*eventproto.EventUpdateResponse, error) {
	event, err := serv.Calendar.GetEventByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	event.Title = req.GetTitle()
	event.DateFrom = time.Unix(req.DateFrom.GetSeconds(), int64(req.DateFrom.GetNanos()))
	event.DateTo = time.Unix(req.DateTo.GetSeconds(), int64(req.DateTo.GetNanos()))

	if err := serv.Calendar.UpdateEvent(event); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &eventproto.EventUpdateResponse{Event: convertToProtoEvent(event)}, nil
}

// Delete Method for removing event if exists
func (serv *server) Delete(_ context.Context, req *eventproto.EventDeleteRequest) (*eventproto.EventDeleteResponse, error) {
	event, err := serv.Calendar.GetEventByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err := serv.Calendar.RemoveEvent(event); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}

// GetListByPeriod Return all events by date period
func (serv *server) GetListByPeriod(_ context.Context, req *eventproto.EventGetListByPeriodRequest) (*eventproto.EventGetListByPeriodResponse, error) {
	events, err := serv.Calendar.GetEventsByPeriod(
		time.Unix(req.DateFrom.GetSeconds(), int64(req.DateFrom.GetNanos())),
		time.Unix(req.DateTo.GetSeconds(), int64(req.DateTo.GetNanos())),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	respEvents := make([]*eventproto.Event, 0, len(events))
	for _, event := range events {
		respEvents = append(respEvents, convertToProtoEvent(event))
	}

	return &eventproto.EventGetListByPeriodResponse{Events: respEvents}, nil
}

func convertToProtoEvent(event *domain.Event) *eventproto.Event {
	return &eventproto.Event{
		Id:       event.ID.String(),
		Title:    event.Title,
		DateFrom: &timestamp.Timestamp{Seconds: event.DateFrom.Unix(), Nanos: int32(event.DateFrom.UnixNano())},
		DateTo:   &timestamp.Timestamp{Seconds: event.DateTo.Unix(), Nanos: int32(event.DateTo.UnixNano())},
	}
}

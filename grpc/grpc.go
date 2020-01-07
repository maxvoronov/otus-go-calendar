package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	eventproto "github.com/maxvoronov/otus-go-calendar/grpc/proto"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// ServerConfig contains general options for gRPC server
type ServerConfig struct {
	Storage     domain.StorageInterface
	Logger      *logrus.Logger
	Host        string
	Port        string
	ConnTimeout time.Duration
}

// Create Method for creating new event
func (g *ServerConfig) Create(_ context.Context, req *eventproto.EventCreateRequest) (*eventproto.EventCreateResponse, error) {
	event := domain.NewEvent(
		req.GetTitle(),
		time.Unix(req.DateFrom.GetSeconds(), int64(req.DateFrom.GetNanos())),
		time.Unix(req.DateTo.GetSeconds(), int64(req.DateTo.GetNanos())),
	)

	if err := g.Storage.Save(event); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &eventproto.EventCreateResponse{Event: g.convertToProtoEvent(event)}, nil
}

// Update Method for updating event if exists
func (g *ServerConfig) Update(_ context.Context, req *eventproto.EventUpdateRequest) (*eventproto.EventUpdateResponse, error) {
	event, err := g.Storage.GetByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	event.Title = req.GetTitle()
	event.DateFrom = time.Unix(req.DateFrom.GetSeconds(), int64(req.DateFrom.GetNanos()))
	event.DateTo = time.Unix(req.DateTo.GetSeconds(), int64(req.DateTo.GetNanos()))

	if err := g.Storage.Save(event); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &eventproto.EventUpdateResponse{Event: g.convertToProtoEvent(event)}, nil
}

// Delete Method for removing event if exists
func (g *ServerConfig) Delete(_ context.Context, req *eventproto.EventDeleteRequest) (*eventproto.EventDeleteResponse, error) {
	event, err := g.Storage.GetByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err := g.Storage.Remove(event); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}

// GetListByPeriod Return all events by date period
func (g *ServerConfig) GetListByPeriod(_ context.Context, req *eventproto.EventGetListByPeriodRequest) (*eventproto.EventGetListByPeriodResponse, error) {
	events, err := g.Storage.GetByPeriod(
		time.Unix(req.DateFrom.GetSeconds(), int64(req.DateFrom.GetNanos())),
		time.Unix(req.DateTo.GetSeconds(), int64(req.DateTo.GetNanos())),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	respEvents := make([]*eventproto.Event, 0, len(events))
	for _, event := range events {
		respEvents = append(respEvents, g.convertToProtoEvent(event))
	}

	return &eventproto.EventGetListByPeriodResponse{Events: respEvents}, nil
}

func (g *ServerConfig) convertToProtoEvent(event *domain.Event) *eventproto.Event {
	return &eventproto.Event{
		Id:       event.ID.String(),
		Title:    event.Title,
		DateFrom: &timestamp.Timestamp{Seconds: event.DateFrom.Unix(), Nanos: int32(event.DateFrom.UnixNano())},
		DateTo:   &timestamp.Timestamp{Seconds: event.DateTo.Unix(), Nanos: int32(event.DateTo.UnixNano())},
	}
}

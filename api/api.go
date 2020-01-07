package api

import (
	"github.com/gorilla/mux"
	"github.com/maxvoronov/otus-go-calendar/api/handler"
	"github.com/maxvoronov/otus-go-calendar/api/middleware"
	"github.com/maxvoronov/otus-go-calendar/internal/config"
	"net"
	"net/http"
	"time"
)

// ServerConfig contains general options for HTTP server
type ServerConfig struct {
	Host           string
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func initRouter(h *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.Handle(h.VersionHandler))
	router.HandleFunc("/create_event", h.Handle(h.EventCreateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/update_event", h.Handle(h.EventUpdateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/delete_event", h.Handle(h.EventDeleteHandler)).Methods(http.MethodPost)
	router.HandleFunc("/events_for_day", h.Handle(h.EventsForDayHandler)).Methods(http.MethodGet)
	router.HandleFunc("/events_for_week", h.Handle(h.EventsForWeekHandler)).Methods(http.MethodGet)
	router.HandleFunc("/events_for_month", h.Handle(h.EventsForMonthHandler)).Methods(http.MethodGet)

	return router
}

// StartServer Start HTTP server for API
func StartServer(cfg *config.Config, config *ServerConfig) {
	h := &handler.Handler{
		Storage: cfg.Storage,
		Logger:  cfg.Logger,
	}
	router := initRouter(h)

	// Apply middleware
	router.Use(middleware.RecoverMiddleware(h.Logger))
	router.Use(middleware.LoggerMiddleware(h.Logger))

	server := &http.Server{
		Addr:           net.JoinHostPort(config.Host, config.Port),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		cfg.Logger.Fatalf("HTTP server error: %s", err)
	}
}
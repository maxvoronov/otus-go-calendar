package sql

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// Storage struct
type Storage struct {
	ConnPool *pgxpool.Pool
	Logger   *logrus.Logger
}

// NewStorage Create new database storage (postgres)
func NewStorage(dbconf *DatabaseConfig, logger *logrus.Logger) (*Storage, error) {
	storage := &Storage{}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbconf.User, dbconf.Password, dbconf.Host, dbconf.Port, dbconf.Database)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse connection configs")
	}

	cfg.MaxConns = 8
	cfg.ConnConfig.TLSConfig = nil
	cfg.ConnConfig.Logger = logrusadapter.NewLogger(logger)
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: 5 * time.Minute,
		Timeout:   1 * time.Second,
	}).DialContext

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to postgres")
	}
	storage.ConnPool = pool
	storage.Logger = logger

	return storage, nil
}

// GetAll Return list of all events
func (storage *Storage) GetAll(ctx context.Context) ([]*domain.Event, error) {
	sql := "SELECT * FROM events"
	result, err := storage.getEventsBySQL(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive events")
	}

	return result, nil
}

// GetByID Return event by ID
func (storage *Storage) GetByID(ctx context.Context, id string) (*domain.Event, error) {
	sql := "SELECT * FROM events WHERE id = $1 LIMIT 1"
	result, err := storage.getEventsBySQL(ctx, sql, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive event")
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result[0], nil
}

// GetByPeriod Return list of events by period
func (storage *Storage) GetByPeriod(ctx context.Context, from, to time.Time) ([]*domain.Event, error) {
	sql := "SELECT * FROM events WHERE date_from >= $1 AND date_to <= $2"
	result, err := storage.getEventsBySQL(ctx, sql, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive events")
	}

	return result, nil
}

// GetForNotification Return list of events for notifications
func (storage *Storage) GetForNotification(ctx context.Context, from, to time.Time) ([]*domain.Event, error) {
	sql := "SELECT * FROM events WHERE status = $1 AND date_from BETWEEN $2 AND $3"
	result, err := storage.getEventsBySQL(ctx, sql, domain.EventStatusNew, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive events")
	}

	return result, nil
}

// Save Create or update event in storage
func (storage *Storage) Save(ctx context.Context, event *domain.Event) error {
	sql := `
		INSERT INTO events (id,title,status,date_from,date_to) VALUES($1,$2,$3,$4,$5)
		ON CONFLICT (id) DO
		UPDATE SET title = $2, status = $3, date_from = $4, date_to = $5
	`
	_, err := storage.ConnPool.Exec(ctx, sql, event.ID.String(), event.Title, event.Status, event.DateFrom, event.DateTo)
	if err != nil {
		return errors.Wrap(err, "failed to update event")
	}

	return nil
}

// Remove event from storage
func (storage *Storage) Remove(ctx context.Context, event *domain.Event) error {
	sql := "DELETE FROM events WHERE id = $1"
	if _, err := storage.ConnPool.Exec(ctx, sql, event.ID.String()); err != nil {
		return errors.Wrap(err, "failed to remove event")
	}

	return nil
}

func (storage *Storage) getEventsBySQL(ctx context.Context, sql string, args ...interface{}) ([]*domain.Event, error) {
	rows, err := storage.ConnPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare query")
	}
	defer rows.Close()

	result := make([]*domain.Event, 0)
	for rows.Next() {
		var id, title, status string
		var dateFrom, dateTo time.Time
		if err := rows.Scan(&id, &title, &status, &dateFrom, &dateTo); err != nil {
			return nil, errors.Wrap(err, "failed to scan result into vars")
		}

		result = append(result, &domain.Event{
			ID:       uuid.FromStringOrNil(id),
			Title:    title,
			Status:   status,
			DateFrom: dateFrom,
			DateTo:   dateTo,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to load result")
	}

	return result, nil
}

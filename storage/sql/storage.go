package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

// Storage struct
type Storage struct {
	Ctx      context.Context
	ConnPool *pgxpool.Pool
	Logger   *logrus.Logger
}

// NewStorage Create new database storage (postgres)
func NewStorage(dbconf *DatabaseConfig, logger *logrus.Logger) (*Storage, error) {
	storage := &Storage{Ctx: context.Background()}
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

	pool, err := pgxpool.ConnectConfig(storage.Ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to postgres")
	}
	storage.ConnPool = pool
	storage.Logger = logger

	return storage, nil
}

// GetAll Return list of all events
func (storage *Storage) GetAll() ([]*domain.Event, error) {
	sql := "SELECT * FROM events"
	result, err := storage.getEventsBySQL(sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive events")
	}

	return result, nil
}

// GetByID Return event by ID
func (storage *Storage) GetByID(id string) (*domain.Event, error) {
	sql := "SELECT * FROM events WHERE id = $1 LIMIT 1"
	result, err := storage.getEventsBySQL(sql, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive event")
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result[0], nil
}

// GetByPeriod Return list of events by period
func (storage *Storage) GetByPeriod(from, to time.Time) ([]*domain.Event, error) {
	sql := "SELECT * FROM events WHERE date_from >= $1 AND date_to <= $2"
	result, err := storage.getEventsBySQL(sql, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "failed to receive events")
	}

	return result, nil
}

// Save Create or update event in storage
func (storage *Storage) Save(event *domain.Event) error {
	existEvent, err := storage.GetByID(event.ID.String())
	if err != nil {
		return err
	}

	// Create new event
	if existEvent == nil {
		sql := "INSERT INTO events (id,title,date_from,date_to) VALUES($1,$2,$3,$4)"
		_, err := storage.ConnPool.Exec(storage.Ctx, sql, event.ID.String(), event.Title, event.DateFrom, event.DateTo)
		if err != nil {
			return errors.Wrap(err, "failed to update event")
		}

		return nil
	}

	// or update event
	sql := "UPDATE events SET title = $1, date_from = $2, date_to = $3 WHERE id = $4"
	_, err = storage.ConnPool.Exec(storage.Ctx, sql, event.Title, event.DateFrom, event.DateTo, event.ID.String())
	if err != nil {
		return errors.Wrap(err, "failed to update event")
	}

	return nil
}

// Remove event from storage
func (storage *Storage) Remove(event *domain.Event) error {
	sql := "DELETE FROM events WHERE id = $1"
	if _, err := storage.ConnPool.Exec(storage.Ctx, sql, event.ID.String()); err != nil {
		return errors.Wrap(err, "failed to remove event")
	}

	return nil
}

func (storage *Storage) getEventsBySQL(sql string, args ...interface{}) ([]*domain.Event, error) {
	rows, err := storage.ConnPool.Query(storage.Ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare query")
	}
	defer rows.Close()

	result := make([]*domain.Event, 0)
	for rows.Next() {
		var id, title string
		var dateFrom, dateTo time.Time
		if err := rows.Scan(&id, &title, &dateFrom, &dateTo); err != nil {
			return nil, errors.Wrap(err, "failed to scan result into vars")
		}

		result = append(result, &domain.Event{
			ID:       uuid.FromStringOrNil(id),
			Title:    title,
			DateFrom: dateFrom,
			DateTo:   dateTo,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to load result")
	}

	return result, nil
}

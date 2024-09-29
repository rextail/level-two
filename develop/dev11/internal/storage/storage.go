package storage

import (
	"context"
	"dev11/internal/dberrs"
	"dev11/internal/entities"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Postgre interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	IsUniqueConstraintError(err error) (bool, error)
}

type Storage struct {
	DB Postgre
}

// InitStorage возвращает указатель на структуру Storage. Инициализирует таблицу с ивентами, если ее еще не существует.
func InitStorage(ctx context.Context, db Postgre) (*Storage, error) {
	const op = `internal.storage.New`

	query := `
	CREATE TABLE IF NOT EXISTS events(
		title TEXT UNIQUE,
		description TEXT,
		date DATE
)`
	_, err := db.Exec(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db}, nil
}

// CreateEvent создает в базе данных событие на основе запроса. Каждое событие обязано иметь уникальное название.
func (s *Storage) CreateEvent(ctx context.Context, req entities.CreateEventRequest) error {
	const op = `internal.storage.CreateEvent`

	query := `INSERT INTO events VALUES ($1, $2, $3)`

	_, err := s.DB.Exec(ctx, query, req.Title, req.Description, req.Date)
	if err != nil {
		notUniqueErr, err := s.DB.IsUniqueConstraintError(err)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		//Если ошибка связана с тем, что мы попытались создать ивент, который уже лежит в базе
		if notUniqueErr {
			return fmt.Errorf("%s: %w", op, dberrs.ErrorNotUniqueTitle)
		}

		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// UpdateEvent обновляет событие в базе данных на основе запроса.
// Если одно из полей пустое, его значение в базе не изменится.
func (s *Storage) UpdateEvent(ctx context.Context, req entities.UpdateEventRequest) error {
	const op = `internal.storage.UpdateEvent`

	query := `
		UPDATE events
		SET 
			description = CASE WHEN $2 <> '' THEN $2 ELSE description END,
			date = CASE WHEN $3 <> '' THEN $3 ELSE date END
		WHERE title = $1
	`

	_, err := s.DB.Exec(ctx, query, req.Title, req.Description, req.Date)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// DeleteEvent удаляет событие из базы по запросу.
func (s *Storage) DeleteEvent(ctx context.Context, req entities.DeleteRequest) error {
	const op = `internal.storage.DeleteEvent`

	query := `DELETE FROM events WHERE $1`

	_, err := s.DB.Exec(ctx, query, req.Title)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) EventsForDay(ctx context.Context) ([]entities.Event, error) {
	const op = `internal.storage.EventsForDay`

	var query string

	query = `SELECT * FROM events WHERE date = CURRENT_DATE`

	var events []entities.Event

	res, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer res.Close()

	for res.Next() {
		var event entities.Event
		if err = res.Scan(&event.Title, &event.Description, &event.Date); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		events = append(events, event)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return events, nil
}

func (s *Storage) EventsForWeek(ctx context.Context) ([]entities.Event, error) {
	const op = `internal.storage.EventsForWeek`

	var query string

	query = `
			SELECT * 
			FROM events
			WHERE EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM CURRENT_DATE)
			AND EXTRACT(MONTH FROM date) = EXTRACT(MONTH FROM CURRENT_DATE)
		`

	var events []entities.Event

	res, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer res.Close()

	for res.Next() {
		var event entities.Event
		if err = res.Scan(&event.Title, &event.Description, &event.Date); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		events = append(events, event)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return events, nil
}

func (s *Storage) EventsForYear(ctx context.Context) ([]entities.Event, error) {
	const op = `internal.storage.EventsForYear`

	var query string

	query = `SELECT * FROM events WHERE EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM CURRENT_DATE)`

	var events []entities.Event

	res, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer res.Close()

	for res.Next() {
		var event entities.Event
		if err = res.Scan(&event.Title, &event.Description, &event.Date); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		events = append(events, event)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return events, nil
}

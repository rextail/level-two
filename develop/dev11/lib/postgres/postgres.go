package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

const (
	defaultMaxConns     = 10
	defaultConnAttempts = 10
	defaultRetryTimeout = 5 * time.Second
	ErrorNotUniqueCode  = `23505`
)

type Postgres struct {
	maxConns     int
	connAttempts int
	retryTimeout time.Duration

	*pgxpool.Pool
}

func New(connStr string, opts ...Option) (*Postgres, error) {
	const op = `lib.postgres.postgres.New`

	pg := &Postgres{
		maxConns:     defaultMaxConns,
		connAttempts: defaultConnAttempts,
		retryTimeout: defaultRetryTimeout,
	}

	for _, opt := range opts {
		opt(pg)
	}

	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	poolCfg.MaxConns = int32(pg.maxConns)

	maxTries := pg.connAttempts

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolCfg)
		if err == nil {
			break
		}

		log.Printf("Postgres is connecting, tries left %d/%d", pg.connAttempts, maxTries)

		time.Sleep(pg.retryTimeout)
		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return pg, nil
}

func (p *Postgres) IsUniqueConstraintError(err error) (bool, error) {
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return false, errors.New("can't assert error to db error")
	}
	if pgErr.Code == ErrorNotUniqueCode {
		return true, nil
	}
	return false, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

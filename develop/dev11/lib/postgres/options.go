package postgres

import "time"

type Option func(p *Postgres)

func MaxConns(max int) Option {
	return func(p *Postgres) {
		p.maxConns = max
	}
}

func ConnAttempts(attempts int) Option {
	return func(p *Postgres) {
		p.connAttempts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(p *Postgres) {
		p.retryTimeout = timeout
	}
}

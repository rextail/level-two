package forday

import (
	"context"
	"dev11/internal/entities"
	"dev11/internal/transport/http/send"
	"dev11/lib/http/response"
	"dev11/lib/slogkz"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type DayEventer interface {
	EventsForDay(ctx context.Context) ([]entities.Event, error)
}

// New возвращает функцию-обработчик запросов на фильтрацию событий для текущего дня.
func New(log *slog.Logger, eventer DayEventer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = `transport.http.handlers.get.forday.New`

		if r.Method != http.MethodGet {
			fmt.Printf("%s: http method check failed", op)

			send.ErrorJSON(w, response.Error("not allowed method ", strconv.Itoa(http.StatusMethodNotAllowed)))

			return
		}

		log = log.With(
			"op", op,
			"remote_addr", r.RemoteAddr,
		)

		events, err := eventer.EventsForDay(r.Context())
		if err != nil {
			log.Error("Error", slogkz.Err(err))

			send.ErrorJSON(w, response.Error(err.Error(), response.UsecaseErrCode))

			return
		}

		send.ResultJSON(w, response.Result(events))
	}
}

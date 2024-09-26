package update

import (
	"context"
	"dev11/internal/entities"
	"dev11/internal/transport/http/send"
	"dev11/lib/http/response"
	"dev11/lib/slogkz"
	"encoding/json"
	v10 "github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type EventUpdater interface {
	UpdateEvent(ctx context.Context, req entities.UpdateEventRequest) error
}

// New возвращает функцию-обработчик запросов на обновление события. Принимает объект, непосредственно занимающийся
// обновлением события.
func New(log *slog.Logger, validator v10.Validate, updater EventUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = `transport.http.handlers.update.New`

		log = log.With(
			slog.String("op", op),
			slog.String("remote_addr", r.RemoteAddr),
		)

		var req entities.UpdateEventRequest

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&req); err != nil {
			log.Error("failed to decode request", slogkz.Err(err))

			send.ErrorJSON(w, response.Error(err.Error(), response.OtherErrCode))

			return
		}
		if err := validator.Struct(req); err != nil {
			validateErr := err.(v10.ValidationErrors)

			log.Error("failed to validate request", slogkz.Err(err))

			send.ErrorJSON(w, response.ValidationError(validateErr))

			return
		}

		if err := updater.UpdateEvent(r.Context(), req); err != nil {
			log.Error("failed to update event", slogkz.Err(err))

			send.ErrorJSON(w, response.Error(err.Error(), response.UsecaseErrCode))

			return
		}

		send.OkJSON(w)
	}
}

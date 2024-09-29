package create

import (
	"context"
	"dev11/internal/dberrs"
	"dev11/internal/entities"
	"dev11/internal/transport/http/send"
	"dev11/lib/http/response"
	"dev11/lib/slogkz"
	"encoding/json"
	"errors"
	"fmt"
	v10 "github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strconv"
)

type EventCreator interface {
	CreateEvent(ctx context.Context, event entities.CreateEventRequest) error
}

// New возвращает функцию-обработчик запросов на создание события. Принимает объект, непосредственно занимающийся
// созданием события.
func New(log *slog.Logger, validator *v10.Validate, creater EventCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = `transport.http.handlers.post.create.New`

		if r.Method != http.MethodPost {
			fmt.Printf("%s: http method check failed", op)

			send.ErrorJSON(w, response.Error("not allowed method ", strconv.Itoa(http.StatusMethodNotAllowed)))

			return
		}

		log = log.With(
			slog.String("op", op),
			slog.String("remote_addr", r.RemoteAddr),
		)

		var req entities.CreateEventRequest

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

		if err := creater.CreateEvent(r.Context(), req); err != nil {
			log.Error("failed to create event", slogkz.Err(err))

			if errors.Is(err, dberrs.ErrorNotUniqueTitle) {
				send.ErrorJSON(w, response.Error("event with this title already exists", response.UsecaseErrCode))
				return
			}

			send.ErrorJSON(w, response.Error(err.Error(), response.UsecaseErrCode))

			return
		}

		send.OkJSON(w)
	}
}

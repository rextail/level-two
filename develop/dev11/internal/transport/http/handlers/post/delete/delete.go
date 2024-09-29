package delete

import (
	"context"
	"dev11/internal/entities"
	"dev11/internal/transport/http/send"
	"dev11/lib/http/response"
	"dev11/lib/slogkz"
	"encoding/json"
	"fmt"
	v10 "github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strconv"
)

type EventDeleter interface {
	DeleteEvent(ctx context.Context, req entities.DeleteRequest) error
}

// New возвращает функцию-обработчик запросов на обновление события. Принимает объект, непосредственно занимающийся
// обновлением события.
func New(log *slog.Logger, validator *v10.Validate, deleter EventDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = `transport.http.handlers.post.delete.New`

		if r.Method != http.MethodPost {
			fmt.Printf("%s: http method check failed", op)

			send.SendJSON(w, response.Error("not allowed method ", strconv.Itoa(http.StatusMethodNotAllowed)))

			return
		}

		log = log.With(
			slog.String("op", op),
			slog.String("remote_addr", r.RemoteAddr),
		)

		var req entities.DeleteRequest

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&req); err != nil {
			log.Error("failed to decode request", slogkz.Err(err))

			send.SendJSON(w, response.Error(err.Error(), response.OtherErrCode))

			return
		}
		if err := validator.Struct(req); err != nil {
			validateErr := err.(v10.ValidationErrors)

			log.Error("failed to validate request", slogkz.Err(err))

			send.SendJSON(w, response.ValidationError(validateErr))

			return
		}

		if err := deleter.DeleteEvent(r.Context(), req); err != nil {
			log.Error("failed to delete event", slogkz.Err(err))

			send.SendJSON(w, response.Error(err.Error(), response.UsecaseErrCode))

			return
		}

		send.SendJSON(w, response.OK())
	}
}

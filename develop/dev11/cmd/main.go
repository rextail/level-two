package main

import (
	"context"
	"dev11/config"
	"dev11/internal/storage"
	"dev11/internal/transport/http/handlers/get/forday"
	"dev11/internal/transport/http/handlers/get/forweek"
	"dev11/internal/transport/http/handlers/get/foryear"
	"dev11/internal/transport/http/handlers/post/create"
	del "dev11/internal/transport/http/handlers/post/delete"
	"dev11/internal/transport/http/handlers/post/update"
	"dev11/internal/transport/http/middleware/logger"
	"dev11/lib/postgres"
	"dev11/lib/validation"
	v10 "github.com/go-playground/validator/v10"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad("config/config.yml")

	pg, err := postgres.New(cfg.ConnStr, postgres.MaxConns(cfg.MaxPoolSize))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()

	strge, err := storage.InitStorage(ctx, pg)
	if err != nil {
		log.Fatal(err)
	}

	validator := v10.New()
	validator.RegisterValidation("date", validation.ValidateDate)

	mux := http.NewServeMux()

	log := slog.Default()

	// Регистрируем обработчики с логированием
	mux.Handle("/create", create.New(log, validator, strge))
	mux.Handle("/update", update.New(log, validator, strge))
	mux.Handle("/delete", del.New(log, validator, strge))

	mux.Handle("/events_for_day", forday.New(log, strge))
	mux.Handle("/events_for_week", forweek.New(log, strge))
	mux.Handle("/events_for_year", foryear.New(log, strge))

	loggedMux := middleware.Logging(mux) // Применяем middleware к mux

	// Запускаем HTTP-сервер
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

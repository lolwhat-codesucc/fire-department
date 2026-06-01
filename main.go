package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"

	"fire-department/modules/handler"
	"fire-department/modules/postgres"
	"fire-department/modules/cli"
)

func main() {
	// Подключение к БД
	dsn := "postgres://doa@localhost:5432/fire_department?sslmode=disable"
	runMigrations(dsn)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}
	config.MaxConns = 10
	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if len(os.Args) > 1 && os.Args[1] == "cli" {
		// Запускаем интерактивный CLI
		cli.Run(pool)
		return
	}

	// Инициализация репозиториев
	rankRepo := postgres.NewRankRepo(pool)
	specRepo := postgres.NewSpecializationRepo(pool)
	districtRepo := postgres.NewDistrictRepo(pool)
	carModelRepo := postgres.NewCarModelRepo(pool)
	callStatusRepo := postgres.NewCallStatusRepo(pool)
	carRepo := postgres.NewCarRepo(pool)
	teamRepo := postgres.NewTeamRepo(pool)
	ffRepo := postgres.NewFirefighterRepo(pool)
	callRepo := postgres.NewCallRepo(pool)

	// Инициализация обработчиков
	rankHandler := handler.NewRankHandler(rankRepo)
	specHandler := handler.NewSpecializationHandler(specRepo)
	districtHandler := handler.NewDistrictHandler(districtRepo)
	carModelHandler := handler.NewCarModelHandler(carModelRepo)
	callStatusHandler := handler.NewCallStatusHandler(callStatusRepo)
	carHandler := handler.NewCarHandler(carRepo)
	teamHandler := handler.NewTeamHandler(teamRepo)
	ffHandler := handler.NewFirefighterHandler(ffRepo)
	callHandler := handler.NewCallHandler(callRepo)

	// Роутер
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	// Маршруты API v1
	r.Route("/api", func(r chi.Router) {
		// Ranks
		r.Route("/ranks", func(r chi.Router) {
			r.Post("/", rankHandler.Create)
			r.Get("/", rankHandler.GetAll)
			r.Get("/{id}", rankHandler.GetByID)
			r.Put("/{id}", rankHandler.Update)
			r.Delete("/{id}", rankHandler.Delete)
		})

		// Specializations
		r.Route("/specializations", func(r chi.Router) {
			r.Post("/", specHandler.Create)
			r.Get("/", specHandler.GetAll)
			r.Get("/{id}", specHandler.GetByID)
			r.Put("/{id}", specHandler.Update)
			r.Delete("/{id}", specHandler.Delete)
		})

		// Districts + specializations subroutes
		r.Route("/districts", func(r chi.Router) {
			r.Post("/", districtHandler.Create)
			r.Get("/", districtHandler.GetAll)
			r.Get("/{id}", districtHandler.GetByID)
			r.Put("/{id}", districtHandler.Update)
			r.Delete("/{id}", districtHandler.Delete)
			r.Post("/{id}/specializations/{specId}", districtHandler.AddSpecialization)
			r.Delete("/{id}/specializations/{specId}", districtHandler.RemoveSpecialization)
			r.Get("/{id}/specializations", districtHandler.GetSpecializations)
		})

		// Car models
		r.Route("/car-models", func(r chi.Router) {
			r.Post("/", carModelHandler.Create)
			r.Get("/", carModelHandler.GetAll)
			r.Get("/{id}", carModelHandler.GetByID)
			r.Put("/{id}", carModelHandler.Update)
			r.Delete("/{id}", carModelHandler.Delete)
		})

		// Call statuses
		r.Route("/call-statuses", func(r chi.Router) {
			r.Post("/", callStatusHandler.Create)
			r.Get("/", callStatusHandler.GetAll)
			r.Get("/{id}", callStatusHandler.GetByID)
			r.Put("/{id}", callStatusHandler.Update)
			r.Delete("/{id}", callStatusHandler.Delete)
		})

		// Cars
		r.Route("/cars", func(r chi.Router) {
			r.Post("/", carHandler.Create)
			r.Get("/", carHandler.GetAll) // ?model=...
			r.Get("/{id}", carHandler.GetByID)
			r.Put("/{id}", carHandler.Update)
			r.Delete("/{id}", carHandler.Delete)
		})

		// Teams
		r.Route("/teams", func(r chi.Router) {
			r.Post("/", teamHandler.Create)
			r.Get("/", teamHandler.GetAll) // ?district= &specialization=
			r.Get("/{number}", teamHandler.GetByNumber)
			r.Put("/{number}", teamHandler.Update)
			r.Delete("/{number}", teamHandler.Delete)
		})

		// Firefighters
		r.Route("/firefighters", func(r chi.Router) {
			r.Post("/", ffHandler.Create)
			r.Get("/", ffHandler.GetAll) // ?team=
			r.Get("/{id}", ffHandler.GetByID)
			r.Put("/{id}", ffHandler.Update)
			r.Delete("/{id}", ffHandler.Delete)
		})

		// Calls
		r.Route("/calls", func(r chi.Router) {
			r.Post("/", callHandler.Create)
			r.Get("/", callHandler.GetAll) // ?team= &car= &district= &status=
			r.Get("/{id}", callHandler.GetByID)
			r.Put("/{id}", callHandler.Update)
			r.Delete("/{id}", callHandler.Delete)
		})
	})

	// Запуск сервера
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Выключение сервера...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(shutdownCtx)
	}()

	log.Println("Сервер запущен на :8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Сервер остановлен")
}

func runMigrations(dsn string) {
	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		log.Fatal("Ошибка инициализации миграций:", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Ошибка применения миграций:", err)
	}
	log.Println("Миграции успешно применены")
}

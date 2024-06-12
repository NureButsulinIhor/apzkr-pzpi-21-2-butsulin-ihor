package main

import (
	"apz-backend/handlers/admin"
	"apz-backend/handlers/device"
	"apz-backend/handlers/manager"
	"apz-backend/handlers/oauth"
	"apz-backend/handlers/register"
	"apz-backend/handlers/worker"
	"apz-backend/internal/config"
	"apz-backend/internal/log"
	"apz-backend/storage/gormstorage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	cfg := config.NewConfig()

	// Logger
	logger, writeCloser := log.NewLogger(cfg.BuildMode, cfg.IsConsoleLogger, cfg.LogFilePath)
	defer writeCloser.Close()
	logger.Info("Logger initialised")

	// Storage
	storage, err := gormstorage.NewStorage(cfg.ConnectionString)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	_ = storage

	// Auth
	tokenAuth := jwtauth.New("HS256", []byte(cfg.AuthSecret), nil)
	_ = tokenAuth

	// Router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	// Without auth
	r.Route("/login", func(r chi.Router) {
		r.Post("/", oauth.Login(logger, tokenAuth, cfg.GoogleClientID, storage))
	})

	// Admin routes group
	r.Route("/register", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/device", func(r chi.Router) {
			r.Post("/", register.Device(logger, tokenAuth, storage))
			r.Delete("/", register.DeleteDevice(logger, tokenAuth, storage))
		})

		r.Route("/manager", func(r chi.Router) {
			r.Post("/", register.Manager(logger, tokenAuth, storage))
			r.Get("/all", register.GetManagers(logger, tokenAuth, storage))
		})

		r.Route("/worker", func(r chi.Router) {
			r.Post("/", register.Worker(logger, tokenAuth, storage))
			r.Delete("/", register.DeleteWorker(logger, tokenAuth, storage))
		})
	})
	r.Route("/admin", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/car", func(r chi.Router) {
			r.Post("/", admin.AddCar(logger, storage))
			r.Delete("/", admin.DeleteCar(logger, storage))
			r.Get("/{carID}", admin.GetCar(logger, storage))
			r.Get("/all", admin.GetCars(logger, storage))
		})

		r.Route("/item", func(r chi.Router) {
			r.Post("/", admin.AddItem(logger, storage))
			r.Put("/", admin.UpdateItem(logger, storage))
		})

		r.Route("/warehouse", func(r chi.Router) {
			r.Post("/", admin.AddWarehouse(logger, storage))
			r.Delete("/", admin.DeleteWarehouse(logger, storage))
			r.Get("/{warehouseID}", admin.GetWarehouse(logger, storage))
			r.Get("/all", admin.GetWarehouses(logger, storage))
		})

		r.Route("/slot", func(r chi.Router) {
			r.Post("/", admin.AddSlot(logger, storage))
			r.Delete("/", admin.DeleteSlot(logger, storage))
			r.Get("/{slotID}", admin.GetSlot(logger, storage))
			r.Put("/", admin.UpdateSlot(logger, storage))
		})
	})

	// Manager routes group
	r.Route("/manager", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/car", func(r chi.Router) {
			r.Get("/all", manager.GetCars(logger, storage))
		})

		r.Route("/task", func(r chi.Router) {
			r.Post("/", manager.AddTask(logger, storage))
			r.Get("/{taskID}", manager.GetTask(logger, storage))
			r.Get("/all", manager.GetTasks(logger, storage))
			r.Delete("/", manager.DeleteTask(logger, storage))
		})

		r.Route("/timetable", func(r chi.Router) {
			r.Post("/", manager.AddToTimetable(logger, storage))
		})

		r.Route("/transfer", func(r chi.Router) {
			r.Post("/", manager.AddTransfer(logger, storage))
			r.Get("/all/{carID}", manager.GetTransfers(logger, storage))
		})

		r.Route("/warehouse", func(r chi.Router) {
			r.Get("/", manager.GetWarehouse(logger, storage))
		})
	})

	// Device routes group
	r.Route("/device", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/polling", device.Polling(logger, storage))
	})

	// Worker routes group
	r.Route("/worker", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/task", func(r chi.Router) {
			r.Get("/", worker.GetTasks(logger, storage))
			r.Post("/", worker.SetDone(logger, storage))
		})
	})

	// Server
	server := &http.Server{
		Addr:              cfg.URL,
		Handler:           r,
		ReadHeaderTimeout: time.Duration(cfg.Timeout) * time.Second,
	}

	// Start server
	logger.Info("Starting server...", slog.String("url", cfg.URL))
	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}

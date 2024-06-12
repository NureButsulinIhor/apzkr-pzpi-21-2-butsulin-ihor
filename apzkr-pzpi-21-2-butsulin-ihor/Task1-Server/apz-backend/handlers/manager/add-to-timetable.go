package manager

import (
	"apz-backend/services/employee"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"time"
)

type AddToTimetableRequestData struct {
	WorkerID       uint `json:"workerID"`
	StartWorkShift int  `json:"startWorkShift"`
	EndWorkShift   int  `json:"endWorkShift"`
}

func AddToTimetable(logger *slog.Logger, storage employee.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.AddToTimetable"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody AddToTimetableRequestData
		err := render.DecodeJSON(r.Body, &requestBody)
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  "invalid request body",
				Body:   nil,
			})
			l.Debug("err in decoding json")
			return
		}

		err = employee.AddToTimetable(requestBody.WorkerID,
			time.Duration(requestBody.StartWorkShift)*time.Hour,
			time.Duration(requestBody.EndWorkShift)*time.Hour,
			employee.Configuration{
				Logger:  l,
				Storage: storage,
				Context: r.Context(),
			})
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  err.Error(),
				Body:   nil,
			})
			return
		}

		render.JSON(w, r, types.Response[any]{
			Status: true,
			Body:   nil,
		})
	}
}

package worker

import (
	"apz-backend/services/task"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetTasksResponse struct {
	Tasks []models.Task `json:"tasks"`
}

func GetTasks(logger *slog.Logger, storage task.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.worker.GetTasks"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		tasks, err := task.GetForWorker(task.Configuration{
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

		render.JSON(w, r, types.Response[GetTasksResponse]{
			Status: true,
			Body:   GetTasksResponse{Tasks: tasks},
		})
	}
}

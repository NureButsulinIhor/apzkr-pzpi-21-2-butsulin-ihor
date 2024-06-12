package manager

import (
	"apz-backend/services/task"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func GetTask(logger *slog.Logger, storage task.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.GetTask"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		taskIDStr := r.PathValue("taskID")
		taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  "invalid request id",
				Body:   nil,
			})
			l.Debug("invalid request id")
			return
		}

		taskModel, err := task.Get(uint(taskID),
			task.Configuration{
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

		render.JSON(w, r, types.Response[models.Task]{
			Status: true,
			Body:   *taskModel,
		})
	}
}

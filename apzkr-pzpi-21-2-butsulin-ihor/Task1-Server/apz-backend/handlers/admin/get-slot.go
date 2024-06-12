package admin

import (
	"apz-backend/services/slot"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func GetSlot(logger *slog.Logger, storage slot.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.admin.GetSlot"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		slotIDStr := r.PathValue("slotID")
		slotID, err := strconv.ParseUint(slotIDStr, 10, 64)
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

		slotModel, err := slot.Get(uint(slotID),
			slot.Configuration{
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

		render.JSON(w, r, types.Response[models.Slot]{
			Status: true,
			Body:   *slotModel,
		})
	}
}

package admin

import (
	"apz-backend/services/slot"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type DeleteSlotRequestData struct {
	SlotID uint `json:"slotID"`
}

func DeleteSlot(logger *slog.Logger, storage slot.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.admin.DeleteSlot"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody DeleteSlotRequestData
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

		err = slot.Delete(requestBody.SlotID,
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

		render.JSON(w, r, types.Response[any]{
			Status: true,
			Body:   nil,
		})
	}
}

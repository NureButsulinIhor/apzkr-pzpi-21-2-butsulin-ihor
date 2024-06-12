package manager

import (
	"apz-backend/services/transfer"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"time"
)

type AddTransferRequestData struct {
	OutTime time.Time `json:"outTime"`
	CarID   uint      `json:"carID"`
}

func AddTransfer(logger *slog.Logger, storage transfer.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.AddTransfer"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody AddTransferRequestData
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

		err = transfer.Add(requestBody.OutTime, requestBody.CarID,
			transfer.Configuration{
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

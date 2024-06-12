package admin

import (
	"apz-backend/services/item"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type UpdateItemRequestData struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	ItemID      uint    `json:"itemID"`
}

func UpdateItem(logger *slog.Logger, storage item.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.admin.UpdateItem"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody UpdateItemRequestData
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

		err = item.Update(requestBody.ItemID, requestBody.Name, requestBody.Description, requestBody.Weight,
			item.Configuration{
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

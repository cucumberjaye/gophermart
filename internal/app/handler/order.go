package handler

import (
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/cucumberjaye/gophermart/internal/app/middleware"
	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/cucumberjaye/gophermart/pkg/luhn"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (h *Handler) setOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	orderID, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Stack().Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok, err := luhn.Valid(string(orderID))
	if err != nil {
		log.Error().Err(err).Stack().Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !ok {
		log.Error().Err(ErrInvalidOrder).Stack().Send()
		http.Error(w, ErrInvalidOrder.Error(), http.StatusUnprocessableEntity)
		return
	}

	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(errors.New("id must be string")).Send()
		return
	}

	order.ID = string(orderID)
	order.UserID = userID
	err = h.service.SetOrder(order)
	if err != nil {
		if errors.Is(err, ErrOrderExists) {
			log.Error().Err(err).Stack().Send()
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if errors.Is(err, ErrUserOrderExists) {
			w.WriteHeader(http.StatusOK)
			return
		}
		log.Error().Err(err).Stack().Send()
		http.Error(w, "error on server", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(errors.New("id must be string")).Stack().Send()
		return
	}

	output, err := h.service.GetOrders(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(err).Stack().Send()
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, output)
}

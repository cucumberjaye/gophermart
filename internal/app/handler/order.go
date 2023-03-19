package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/cucumberjaye/gophermart/internal/app/middleware"
	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/cucumberjaye/gophermart/pkg/luhn"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (h *Handler) setOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	var orderID int

	err := render.Decode(r, &orderID)
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(errors.New("id must be string")).Send()
		return
	}

	if !luhn.Valid(orderID) {
		log.Error().Err(ErrInvalidOrder).Send()
		http.Error(w, ErrInvalidOrder.Error(), http.StatusUnprocessableEntity)
		return
	}

	order.ID = strconv.Itoa(orderID)
	order.UserID = userID
	err = h.service.SetOrder(order)
	if err != nil {
		if errors.Is(err, ErrOrderExists) {
			log.Error().Err(err).Send()
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if errors.Is(err, ErrUserOrderExists) {
			w.WriteHeader(http.StatusOK)
			return
		}
		log.Error().Err(err).Send()
		http.Error(w, "error on server", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(errors.New("id must be string")).Send()
		return
	}

	output, err := h.service.GetOrders(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(err).Send()
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, output)
}

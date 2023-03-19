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
	var orderId int

	err := render.Decode(r, &orderId)
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(errors.New("id must be string")).Send()
		return
	}

	if !luhn.Valid(orderId) {
		log.Error().Err(ErrInvalidOrder).Send()
		http.Error(w, ErrInvalidOrder.Error(), http.StatusUnprocessableEntity)
		return
	}

	order.Id = strconv.Itoa(orderId)
	order.UserId = userId
	err = h.service.SetOrder(order)
	if err != nil {
		if errors.Is(err, ErrBadOrder) {
			log.Error().Err(err).Send()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors.Is(err, ErrOrderExists) {
			log.Error().Err(err).Send()
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if errors.Is(err, ErrUserOrderExists) {
			w.WriteHeader(http.StatusAccepted)
			return
		}
		log.Error().Err(err).Send()
		http.Error(w, "error on server", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Err(errors.New("id must be string")).Send()
		return
	}

	output, err := h.service.GetOrders(userId)
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

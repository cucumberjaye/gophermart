package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cucumberjaye/gophermart/internal/app/middleware"
	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/cucumberjaye/gophermart/pkg/luhn"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getBalance(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Stack().Err(errors.New("id must be string")).Send()
		return
	}

	output, err := h.service.GetBalance(userID)
	if err != nil {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Stack().Err(err).Send()
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, output)
}

func (h *Handler) withdraw(w http.ResponseWriter, r *http.Request) {
	var input models.Withdraw

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok, err := luhn.Valid(input.Order)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		log.Error().Stack().Err(ErrInvalidOrder).Send()
		http.Error(w, ErrInvalidOrder.Error(), http.StatusUnprocessableEntity)
		return
	}

	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Stack().Err(errors.New("id must be string")).Send()
		return
	}

	err = h.service.Withdraw(userID, input)
	if err != nil {
		if errors.Is(err, ErrInsufficientFunds) {
			log.Error().Err(err).Stack().Send()
			http.Error(w, ErrInsufficientFunds.Error(), http.StatusPaymentRequired)
			return
		}

		log.Error().Stack().Err(err).Send()
		http.Error(w, ErrInvalidOrder.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handler) getWithdraws(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Stack().Err(errors.New("id must be string")).Send()
		return
	}

	output, err := h.service.GetWithdrawals(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, "error on server", http.StatusInternalServerError)
		log.Error().Stack().Err(err).Send()
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, output)
}

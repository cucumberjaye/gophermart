package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

var (
	ErrorLoginExists          error = errors.New("login already exists")
	ErrorWrongLoginOrPassword error = errors.New("wrong login or password")
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterUser

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(&input)
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(input)
	if err != nil {
		log.Error().Err(err).Send()
		if errors.Is(err, ErrorLoginExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := h.service.GenerateToken(models.LoginUser(input))
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "authorization",
		Value:   token,
		Expires: time.Now().Add(time.Hour),
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var input models.LoginUser

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(&input)
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	token, err := h.service.GenerateToken(input)
	if err != nil {
		log.Error().Err(err).Send()
		if errors.Is(err, ErrorWrongLoginOrPassword) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "authorization",
		Value:   token,
		Expires: time.Now().Add(time.Hour),
		Path:    "/",
	})
	w.WriteHeader(http.StatusOK)
}

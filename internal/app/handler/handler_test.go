package handler

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNew(t *testing.T) {
	type args struct {
		service MartService
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_InitRoutes(t *testing.T) {
	tests := []struct {
		name string
		h    *Handler
		want *chi.Mux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.InitRoutes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.InitRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

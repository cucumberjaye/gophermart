package handler

/*import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/cucumberjaye/gophermart/internal/app/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_getBalance(t *testing.T) {
	tests := []struct {
		name      string
		balance   models.Balance
		code      int
		response  string
		returning string
		err       error
	}{
		{
			name: "ok",
			balance: models.Balance{
				Current:   500,
				Withdrawn: 500,
			},
			code: 200,
			err:  nil,
		},
	}
	ctrl := gomock.NewController(t)
	svc := mocks.NewMockMartService(ctrl)
	h := &Handler{service: svc}

	r := h.InitRoutes()
	ts := httptest.NewServer(r)

	for _, tt := range tests {
		balance, err := json.Marshal(tt.balance)
		require.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, ts.URL+"/user/balance", nil)
		request.RequestURI = ""
		request.AddCookie(&http.Cookie{
			Name:  "authorization",
			Value: "d0cdf57eeda01eecb13c99625ee5680dad9804e4d2285f48d66c4827dd3d9aa77863ed11d5748f851a08dec507d5ec7ca1f1a859",
			Path:  "/",
		})

		if tt.name != "valid_err" {
			svc.EXPECT().GetBalance("123").Return(tt.balance, nil)
		}
		resp, err := http.DefaultClient.Do(request)
		require.NoError(t, err)

		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, balance, string(resBody))

		require.Equal(t, resp.StatusCode, tt.code)
	}
}

/*func TestHandler_withdraw(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	ctrl := gomock.NewController(t)
	svc := mocks.NewMockMartService(ctrl)
	h := &Handler{service: svc}

	r := h.InitRoutes()
	ts := httptest.NewServer(r)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.withdraw(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_getWithdraws(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	ctrl := gomock.NewController(t)
	svc := mocks.NewMockMartService(ctrl)
	h := &Handler{service: svc}

	r := h.InitRoutes()
	ts := httptest.NewServer(r)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.getWithdraws(tt.args.w, tt.args.r)
		})
	}
}*/

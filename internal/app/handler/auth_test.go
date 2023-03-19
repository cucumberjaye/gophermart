package handler

/*import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/cucumberjaye/gophermart/internal/app/sevice/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_register(t *testing.T) {

	tests := []struct {
		name     string
		body     io.Reader
		user     models.RegisterUser
		code     int
		response string
		err      error
	}{
		{
			name: "ok",
			code: http.StatusOK,
			user: models.RegisterUser{
				Login:    "lol",
				Password: "123456",
			},
			err: nil,
		},
		{
			name: "valid_err",
			code: http.StatusBadRequest,
			user: models.RegisterUser{
				Login:    "",
				Password: "123456",
			},
			response: "invalid request body: Key: 'RegisterUser.Login' Error:Field validation for 'Login' failed on the 'required' tag\n",
			err:      errors.New(""),
		},
		{
			name: "service_err",
			code: http.StatusInternalServerError,
			user: models.RegisterUser{
				Login:    "lol",
				Password: "123456",
			},
			response: "test\n",
			err:      errors.New("test"),
		},
		{
			name: "not_unique_login_err",
			code: http.StatusConflict,
			user: models.RegisterUser{
				Login:    "lol",
				Password: "123456",
			},
			response: ErrorLoginExists.Error() + "\n",
			err:      ErrorLoginExists,
		},
	}

	ctrl := gomock.NewController(t)
	as := mocks.NewMockAuthService(ctrl)
	h := &Handler{authService: as}

	r := h.InitRoutes()
	ts := httptest.NewServer(r)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := json.Marshal(tt.user)
			require.NoError(t, err)

			tt.body = bytes.NewBuffer(user)
			request := httptest.NewRequest(http.MethodPost, ts.URL+"/user/register", tt.body)
			request.RequestURI = ""

			if tt.name != "valid_err" {
				as.EXPECT().CreateUser(tt.user).Return(tt.err)
			}
			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			require.Equal(t, resp.StatusCode, tt.code)
			defer resp.Body.Close()

			resBody, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, tt.response, string(resBody))
		})
	}
}

func TestHandler_login(t *testing.T) {
	tests := []struct {
		name      string
		body      io.Reader
		user      models.LoginUser
		code      int
		response  string
		returning string
		err       error
	}{
		{
			name: "ok",
			code: http.StatusOK,
			user: models.LoginUser{
				Login:    "lol",
				Password: "123456",
			},
			returning: "test",
			err:       nil,
		},
		{
			name: "valid_err",
			code: http.StatusBadRequest,
			user: models.LoginUser{
				Login:    "",
				Password: "123456",
			},
			returning: "",
			response:  "invalid request body: Key: 'LoginUser.Login' Error:Field validation for 'Login' failed on the 'required' tag\n",
			err:       errors.New(""),
		},
		{
			name: "service_err",
			code: http.StatusInternalServerError,
			user: models.LoginUser{
				Login:    "lol",
				Password: "123456",
			},
			returning: "",
			response:  "test\n",
			err:       errors.New("test"),
		},
		{
			name: "not_unique_login_err",
			code: http.StatusUnauthorized,
			user: models.LoginUser{
				Login:    "lol",
				Password: "123456",
			},
			returning: "",
			response:  ErrorWrongLoginOrPassword.Error() + "\n",
			err:       ErrorWrongLoginOrPassword,
		},
	}
	ctrl := gomock.NewController(t)
	as := mocks.NewMockAuthService(ctrl)
	h := &Handler{authService: as}

	r := h.InitRoutes()
	ts := httptest.NewServer(r)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := json.Marshal(tt.user)
			require.NoError(t, err)

			tt.body = bytes.NewBuffer(user)
			request := httptest.NewRequest(http.MethodPost, ts.URL+"/user/login", tt.body)
			request.RequestURI = ""

			if tt.name != "valid_err" {
				as.EXPECT().GenerateToken(tt.user).Return(tt.returning, tt.err)
			}
			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			if tt.name == "ok" {
				c := resp.Cookies()
				require.Equal(t, c[0].Value, tt.returning)
			}

			require.Equal(t, resp.StatusCode, tt.code)
			defer resp.Body.Close()

			resBody, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, tt.response, string(resBody))
		})
	}
}*/

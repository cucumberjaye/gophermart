package middleware

import (
	"compress/gzip"
	"context"
	"github.com/cucumberjaye/gophermart/configs"
	"github.com/cucumberjaye/gophermart/pkg/token"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

func GzipCompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

func GzipDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reader io.Reader

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			reader = gz
			defer gz.Close()
		} else {
			reader = r.Body
		}

		r.Body = io.NopCloser(reader)

		next.ServeHTTP(w, r)
	})
}

type UserID string

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("authorization")
		if err == nil {
			id, err := token.CheckToken(c.Value, []byte(configs.SigningKey))
			if err == nil {
				ctx := context.WithValue(r.Context(), UserID("user_id"), id)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		http.Error(w, "bad authorization token", http.StatusUnauthorized)
		return
	})
}

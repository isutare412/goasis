package http

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	"github.com/isutare412/goasis/internal/pkgctx"
)

func recoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				slog.ErrorContext(r.Context(),
					"http handler panicked",
					"recover", v,
					"stackTrace", string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func insertContextBag(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := pkgctx.WithBag(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func issueRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if bag, ok := pkgctx.GetBag(r.Context()); ok {
			requestID := uuid.NewString()
			bag.RequestID = requestID
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func wrapResponseWriter(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

func accessLog(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timeBeforeServing := time.Now()

		next.ServeHTTP(w, r)

		elapsedTime := time.Since(timeBeforeServing)
		ww := w.(middleware.WrapResponseWriter)

		var statusCode = http.StatusOK
		if sc := ww.Status(); sc != 0 {
			statusCode = sc
		}

		accessLog := slog.With(
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("addr", r.RemoteAddr),
			slog.String("proto", r.Proto),
			slog.Int64("contentLength", r.ContentLength),
			slog.String("userAgent", r.UserAgent()),
			slog.Int("status", statusCode),
			slog.Int("bodyBytes", ww.BytesWritten()),
			slog.Duration("elapsed", elapsedTime),
		)

		if ct := r.Header.Get("Content-Type"); ct != "" {
			accessLog = accessLog.With(slog.String("contentType", ct))
		}

		accessLog.InfoContext(r.Context(), "handle http request")
	}

	return http.HandlerFunc(fn)
}

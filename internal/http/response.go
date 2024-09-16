package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/isutare412/goasis/internal/pkgerr"
	"github.com/isutare412/goasis/pkg/oapi"
)

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		ctx     = r.Context()
		errCode = errorStatusCode(err)
		errMsg  = clientErrorMessage(err, errCode)
		errResp = oapi.ErrorResponse{
			Message: &errMsg,
		}
	)

	errRespBytes, err := json.Marshal(&errResp)
	if err != nil {
		slog.ErrorContext(ctx, "failed to marshal error response", "error", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(errCode)
	if _, err := w.Write(errRespBytes); err != nil {
		slog.ErrorContext(ctx, "failed to write error response body", "error", err)
		return
	}
}

func clientErrorMessage(err error, statusCode int) string {
	if cerr, ok := pkgerr.AsCodeError(err); ok && cerr.ClientMsg != "" {
		return cerr.ClientMsg
	}

	return http.StatusText(statusCode)
}

func errorStatusCode(err error) int {
	if code, ok := oapiErrorStatusCode(err); ok {
		return code
	}

	return http.StatusInternalServerError
}

func oapiErrorStatusCode(err error) (code int, ok bool) {
	var cookieErr *oapi.UnescapedCookieParamError
	if errors.As(err, &cookieErr) {
		return http.StatusBadRequest, true
	}

	var unmarshalErr *oapi.UnmarshalingParamError
	if errors.As(err, &unmarshalErr) {
		return http.StatusBadRequest, true
	}

	var requiredParamErr *oapi.RequiredParamError
	if errors.As(err, &requiredParamErr) {
		return http.StatusBadRequest, true
	}

	var requiredHeaderErr *oapi.RequiredHeaderError
	if errors.As(err, &requiredHeaderErr) {
		return http.StatusBadRequest, true
	}

	var invalidParamFormatErr *oapi.InvalidParamFormatError
	if errors.As(err, &invalidParamFormatErr) {
		return http.StatusBadRequest, true
	}

	var tooManyValuesForParamErr *oapi.TooManyValuesForParamError
	if errors.As(err, &tooManyValuesForParamErr) {
		return http.StatusBadRequest, true
	}

	return 0, false
}

package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/isutare412/goasis/pkg/oapi"
)

func Test_errorStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "error is nil",
			args: args{err: nil},
			want: http.StatusInternalServerError,
		},
		{
			name: "error type is unknown",
			args: args{err: fmt.Errorf("unknown error")},
			want: http.StatusInternalServerError,
		},
		{
			name: "error type is OpenAPI error",
			args: args{err: &oapi.RequiredHeaderError{
				ParamName: "x-test-header",
				Err:       fmt.Errorf("header not found"),
			}},
			want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorStatusCode(tt.args.err); got != tt.want {
				t.Errorf("errorStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseError(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		err error
	}
	tests := []struct {
		name             string
		args             args
		wantMessageRegex string
		wantStatusCode   int
	}{
		{
			name: "unknown error",
			args: args{
				w:   httptest.NewRecorder(),
				err: fmt.Errorf("unknown error"),
			},
			wantMessageRegex: `"message":.+`,
			wantStatusCode:   http.StatusInternalServerError,
		},
		{
			name: "error type is OpenAPI error",
			args: args{
				w: httptest.NewRecorder(),
				err: &oapi.RequiredHeaderError{
					ParamName: "x-test-header",
					Err:       fmt.Errorf("header not found"),
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseError(tt.args.w, tt.args.r, tt.args.err)

			recorder := tt.args.w.(*httptest.ResponseRecorder)

			resp := recorder.Result()
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("statusCode = %v, want %v", resp.StatusCode, tt.wantStatusCode)
			}

			contentType := resp.Header.Get("content-type")
			if !strings.Contains(contentType, "application/json") {
				t.Errorf("contentType = %v, want %v", contentType, "application/json")
			}

			wantMessagePattern := regexp.MustCompile(tt.wantMessageRegex)
			if !wantMessagePattern.MatchString(string(respBody)) {
				t.Errorf("responseBody = %v, want %v", string(respBody), wantMessagePattern)
			}
		})
	}
}

package pkgerr

import (
	"errors"
	"testing"
)

var testError = errors.New("test error")

func TestCodeError_Unwrap(t *testing.T) {
	type fields struct {
		Code      Code
		Err       error
		ClientMsg string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Err should be returned when unwrapped",
			fields: fields{
				Code:      CodeNotFound,
				Err:       testError,
				ClientMsg: "some message for client",
			},
			wantErr: testError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CodeError{
				Code:      tt.fields.Code,
				Err:       tt.fields.Err,
				ClientMsg: tt.fields.ClientMsg,
			}

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("errors.Is(%v, %v) = false, want true", err, tt.wantErr)
			}

			var unwrapped CodeError
			if !errors.As(err, &unwrapped) {
				t.Fatalf("errors.As(%v, %v) = false, want true", err, unwrapped)
			}

			if unwrapped.Code != tt.fields.Code {
				t.Errorf("%v != %v, want equal", unwrapped.Code, tt.fields.Code)
			}
			if unwrapped.Err != tt.fields.Err {
				t.Errorf("%v != %v, want equal", unwrapped.Err, tt.fields.Err)
			}
			if unwrapped.ClientMsg != tt.fields.ClientMsg {
				t.Errorf("%v != %v, want equal", unwrapped.ClientMsg, tt.fields.ClientMsg)
			}
		})
	}
}

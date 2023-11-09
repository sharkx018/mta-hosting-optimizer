package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteCustomResp(t *testing.T) {
	type args struct {
		w            http.ResponseWriter
		headerStatus int
		response     interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test nil",
			args: args{
				w:            httptest.NewRecorder(),
				headerStatus: http.StatusOK,
				response:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteCustomResp(tt.args.w, tt.args.headerStatus, tt.args.response)
		})
	}
}

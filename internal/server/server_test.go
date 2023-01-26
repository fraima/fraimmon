package server

import (
	"net/http"
	"testing"

	"fraima.io/fraimmon/internal/storage"
)

func TestServer_Get(t *testing.T) {
	type fields struct {
		storage storage.Storage
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				storage: tt.fields.storage,
			}
			s.Get(tt.args.w, tt.args.r)
		})
	}
}

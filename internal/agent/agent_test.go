package agent

import (
	"testing"
)

func TestPusher(t *testing.T) {
	type args struct {
		urlList []string
	}
	tests := []struct {
		name    string
		fields  args
		wantErr bool
	}{
		{
			name: "test success",
			fields: args{
				urlList: []string{"http://localhost:8080/update/gauge/totalalloc/215064"},
			},
			wantErr: false,
		},
		{
			name: "test failed <connection refused>",
			fields: args{
				urlList: []string{"http://localhost:8081/update/gauge/totalalloc/215064"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Pusher(tt.fields.urlList); (err != nil) != tt.wantErr {
				t.Errorf("Pusher() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

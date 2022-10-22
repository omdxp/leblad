package leblad

import (
	"testing"
)

func Test_openJsonFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "data/WilayaList.json",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := openJsonFile(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("openJsonFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
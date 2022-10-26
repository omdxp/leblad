package leblad

import (
	"testing"
)

func TestLeblad_GetWilayaList(t *testing.T) {
	tests := []struct {
		name    string
		l       *Leblad
		wantErr bool
	}{
		{
			name:    "GetWilayaList",
			l:       New(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetWilayaList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetWilayaList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

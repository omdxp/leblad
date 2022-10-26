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

func TestLeblad_GetWilayaByZipCode(t *testing.T) {
	type args struct {
		zipCode int
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		wantErr bool
	}{
		{
			name:    "GetWilayaByZipCode",
			l:       New(),
			args:    args{zipCode: 16000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetWilayaByZipCode(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetWilayaByZipCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

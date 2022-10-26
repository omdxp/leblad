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
			name:    "GetWilayaByZipCode for existing zip code",
			l:       New(),
			args:    args{zipCode: 16000},
			wantErr: false,
		},
		{
			name:    "GetWilayaByZipCode for non-existing zip code",
			l:       New(),
			args:    args{zipCode: 99999},
			wantErr: true,
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

func TestLeblad_GetWilayaByCode(t *testing.T) {
	type args struct {
		matricule int
		fields    []string
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		want    Wilaya
		wantErr bool
	}{
		{
			name:    "GetWilayaByCode for existing wilaya",
			l:       New(),
			args:    args{matricule: 1},
			wantErr: false,
		},
		{
			name:    "GetWilayaByCode for non-existing wilaya",
			l:       New(),
			args:    args{matricule: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetWilayaByCode(tt.args.matricule, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetWilayaByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

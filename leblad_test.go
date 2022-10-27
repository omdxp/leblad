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

func TestLeblad_GetAdjacentWilayas(t *testing.T) {
	type args struct {
		matricule int
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "GetAdjacentWilayas for existing wilaya",
			l:       New(),
			args:    args{matricule: 1},
			wantErr: false,
		},
		{
			name:    "GetAdjacentWilayas for non-existing wilaya",
			l:       New(),
			args:    args{matricule: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetAdjacentWilayas(tt.args.matricule)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetAdjacentWilayas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeblad_GetZipCodesForWilaya(t *testing.T) {
	type args struct {
		matricule int
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "GetZipCodesForWilaya for existing wilaya",
			l:       New(),
			args:    args{matricule: 1},
			wantErr: false,
		},
		{
			name:    "GetZipCodesForWilaya for non-existing wilaya",
			l:       New(),
			args:    args{matricule: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetZipCodesForWilaya(tt.args.matricule)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetZipCodesForWilaya() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeblad_GetDairatsForWilaya(t *testing.T) {
	type args struct {
		matricule int
		fields    []string
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		want    []Daira
		wantErr bool
	}{
		{
			name:    "GetDairatsForWilaya for existing wilaya",
			l:       New(),
			args:    args{matricule: 1},
			wantErr: false,
		},
		{
			name:    "GetDairatsForWilaya for non-existing wilaya",
			l:       New(),
			args:    args{matricule: 999},
			wantErr: true,
		},
		{
			name:    "GetDairatsForWilaya for existing wilaya with fields",
			l:       New(),
			args:    args{matricule: 1, fields: []string{"name"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetDairatsForWilaya(tt.args.matricule, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetDairatsForWilaya() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeblad_GetWilayaByPhoneCode(t *testing.T) {
	type args struct {
		phoneCode int
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
			name:    "GetWilayaByPhoneCode for existing phone code",
			l:       New(),
			args:    args{phoneCode: 21},
			wantErr: false,
		},
		{
			name:    "GetWilayaByPhoneCode for non-existing phone code",
			l:       New(),
			args:    args{phoneCode: 999},
			wantErr: true,
		},
		{
			name:    "GetWilayaByPhoneCode for existing phone code with fields",
			l:       New(),
			args:    args{phoneCode: 21, fields: []string{"name"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetWilayaByPhoneCode(tt.args.phoneCode, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetWilayaByPhoneCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeblad_GetWilayaByDairaName(t *testing.T) {
	type args struct {
		dairaName string
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
			name:    "GetWilayaByDairaName for existing daira name",
			l:       New(),
			args:    args{dairaName: "ADRAR"},
			wantErr: false,
		},
		{
			name:    "GetWilayaByDairaName for non-existing daira name",
			l:       New(),
			args:    args{dairaName: "Non existing daira"},
			wantErr: true,
		},
		{
			name:    "GetWilayaByDairaName for existing daira name with fields",
			l:       New(),
			args:    args{dairaName: "ADRAR", fields: []string{"name"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetWilayaByDairaName(tt.args.dairaName, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetWilayaByDairaName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeblad_GetBaladyiatsForDaira(t *testing.T) {
	type args struct {
		dairaName string
		fields    []string
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		want    []Baladyia
		wantErr bool
	}{
		{
			name:    "GetBaladyiatsForDaira for existing daira name",
			l:       New(),
			args:    args{dairaName: "ADRAR"},
			wantErr: false,
		},
		{
			name:    "GetBaladyiatsForDaira for non-existing daira name",
			l:       New(),
			args:    args{dairaName: "Non existing daira"},
			wantErr: true,
		},
		{
			name:    "GetBaladyiatsForDaira for existing daira name with fields",
			l:       New(),
			args:    args{dairaName: "ADRAR", fields: []string{"name"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetBaladyiatsForDaira(tt.args.dairaName, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetBaladyiatsForDaira() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeblad_GetBaladyiatsForDairaCode(t *testing.T) {
	type args struct {
		dairaCode int
		fields    []string
	}
	tests := []struct {
		name    string
		l       *Leblad
		args    args
		want    []Baladyia
		wantErr bool
	}{
		{
			name:    "GetBaladyiatsForDairaCode for existing daira code",
			l:       New(),
			args:    args{dairaCode: 101},
			wantErr: false,
		},
		{
			name:    "GetBaladyiatsForDairaCode for non-existing daira code",
			l:       New(),
			args:    args{dairaCode: 999},
			wantErr: true,
		},
		{
			name:    "GetBaladyiatsForDairaCode for existing daira code with fields",
			l:       New(),
			args:    args{dairaCode: 101, fields: []string{"name"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l
			_, err := l.GetBaladyiatsForDairaCode(tt.args.dairaCode, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leblad.GetBaladyiatsForDairaCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

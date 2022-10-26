package leblad

import (
	"reflect"
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

func Test_unmarshalWilayaListJson(t *testing.T) {
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
			b, _ := openJsonFile(tt.name)
			_, err := unmarshalWilayaListJson(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalWilayaListJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_filterWilayaList(t *testing.T) {
	type args struct {
		wilayas *[]Wilaya
		fields  []string
	}
	tests := []struct {
		name string
		args args
		want *[]Wilaya
	}{
		{
			name: "filterWilayaList with 1 field",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
					},
				},
				fields: []string{
					"name",
				},
			},
			want: &[]Wilaya{
				{
					Name: "Adrar",
				},
			},
		},
		{
			name: "filterWilayaList with 2 fields",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
					},
				},
				fields: []string{
					"name",
					"matricule",
				},
			},
			want: &[]Wilaya{
				{
					Name:      "Adrar",
					Matricule: 1,
				},
			},
		},
		{
			name: "filterWilayaList with non existing field",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
					},
				},
				fields: []string{
					"nonExistingField",
				},
			},
			want: &[]Wilaya{
				{}, // empty struct
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterWilayaList(tt.args.wilayas, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterWilayaList() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_filterWilaya(t *testing.T) {
	type args struct {
		wilaya Wilaya
		fields []string
	}
	tests := []struct {
		name string
		args args
		want Wilaya
	}{
		{
			name: "filterWilaya with 1 field",
			args: args{
				wilaya: Wilaya{
					Name:      "Adrar",
					Matricule: 1,
				},
				fields: []string{
					"name",
				},
			},
			want: Wilaya{
				Name: "Adrar",
			},
		},
		{
			name: "filterWilaya with 2 fields",
			args: args{
				wilaya: Wilaya{
					Name:      "Adrar",
					Matricule: 1,
				},
				fields: []string{
					"name",
					"matricule",
				},
			},
			want: Wilaya{
				Name:      "Adrar",
				Matricule: 1,
			},
		},
		{
			name: "filterWilaya with non existing field",
			args: args{
				wilaya: Wilaya{
					Name:      "Adrar",
					Matricule: 1,
				},
				fields: []string{
					"nonExistingField",
				},
			},
			want: Wilaya{}, // empty struct
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterWilaya(tt.args.wilaya, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterWilaya() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWilayaIndexByZipCode(t *testing.T) {
	type args struct {
		zipCode int
	}
	tests := []struct {
		name    string
		args    args
		wilayas []Wilaya
		want    int
	}{
		{
			name: "getWilayaIndexByZipCode with valid zip code",
			args: args{
				zipCode: 1000,
			},
			wilayas: []Wilaya{
				{
					Name:        "Adrar",
					PostalCodes: []int{1000, 1001, 1002},
				},
				{
					Name:        "Chlef",
					PostalCodes: []int{2000, 2001, 2002},
				},
			},
			want: 0,
		},
		{
			name: "getWilayaIndexByZipCode with invalid zip code",
			args: args{
				zipCode: 999999,
			},
			wilayas: []Wilaya{
				{
					Name:        "Adrar",
					PostalCodes: []int{1000, 1001, 1002},
				},
				{
					Name:        "Chlef",
					PostalCodes: []int{2000, 2001, 2002},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWilayaIndexByZipCode(&tt.wilayas, tt.args.zipCode); got != tt.want {
				t.Errorf("getWilayaIndexByZipCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidZipCode(t *testing.T) {
	type args struct {
		zipCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isValidZipCode with valid zip code",
			args: args{
				zipCode: 1000,
			},
			want: true,
		},
		{
			name: "isValidZipCode with invalid zip code",
			args: args{
				zipCode: 999999,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidZipCode(tt.args.zipCode); got != tt.want {
				t.Errorf("isValidZipCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWilayaIndexByCode(t *testing.T) {
	type args struct {
		wilayas    *[]Wilaya
		wilayaCode int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getWilayaIndexByCode with valid wilaya code",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
					},
					{
						Name:      "Chlef",
						Matricule: 2,
					},
				},
				wilayaCode: 1,
			},
			want: 0,
		},
		{
			name: "getWilayaIndexByCode with invalid wilaya code",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
					},
					{
						Name:      "Chlef",
						Matricule: 2,
					},
				},
				wilayaCode: 999999,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWilayaIndexByCode(tt.args.wilayas, tt.args.wilayaCode); got != tt.want {
				t.Errorf("getWilayaIndexByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidWilayaCode(t *testing.T) {
	type args struct {
		wilayaCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isValidWilayaCode with valid wilaya code",
			args: args{
				wilayaCode: 1,
			},
			want: true,
		},
		{
			name: "isValidWilayaCode with invalid wilaya code",
			args: args{
				wilayaCode: 999999,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidWilayaCode(tt.args.wilayaCode); got != tt.want {
				t.Errorf("isValidWilayaCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

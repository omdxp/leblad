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

func Test_getAdjacentWilayas(t *testing.T) {
	type args struct {
		wilayas     *[]Wilaya
		wilayaIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "getAdjacentWilayas with valid wilaya index",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
						AdjacentWilayas: []int{
							2,
							3,
						},
					},
					{
						Name:      "Chlef",
						Matricule: 2,
						AdjacentWilayas: []int{
							1,
							3,
						},
					},
				},
				wilayaIndex: 0,
			},
			want: []int{
				2,
				3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAdjacentWilayas(tt.args.wilayas, tt.args.wilayaIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAdjacentWilayas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getZipCodes(t *testing.T) {
	type args struct {
		wilayas     *[]Wilaya
		wilayaIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "getZipCodes with valid wilaya index",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name:      "Adrar",
						Matricule: 1,
						PostalCodes: []int{
							1000,
							1001,
						},
					},
				},
				wilayaIndex: 0,
			},
			want: []int{
				1000,
				1001,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getZipCodes(tt.args.wilayas, tt.args.wilayaIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getZipCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDairats(t *testing.T) {
	type args struct {
		wilayas     *[]Wilaya
		wilayaIndex int
	}
	tests := []struct {
		name string
		args args
		want []Daira
	}{
		{
			name: "getDairats with valid wilaya index",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "Adrar",
						Dairats: []Daira{
							{
								Name: "ADRAR",
							},
							{
								Name: "OULED AHMED TIMMI",
							},
						},
					},
				},
				wilayaIndex: 0,
			},
			want: []Daira{
				{
					Name: "ADRAR",
				},
				{
					Name: "OULED AHMED TIMMI",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDairats(tt.args.wilayas, tt.args.wilayaIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDairats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterDaira(t *testing.T) {
	type args struct {
		d      Daira
		fields []string
	}
	tests := []struct {
		name string
		args args
		want Daira
	}{
		{
			name: "filterDaira with valid fields",
			args: args{
				d: Daira{
					Name: "ADRAR",
					Code: 1,
				},
				fields: []string{
					"name",
				},
			},
			want: Daira{
				Name: "ADRAR",
			},
		},
		{
			name: "filterDaira with invalid fields",
			args: args{
				d: Daira{
					Name: "ADRAR",
				},
				fields: []string{
					"invalid",
				},
			},
			want: Daira{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterDaira(tt.args.d, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterDaira() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterDairats(t *testing.T) {
	type args struct {
		dairats []Daira
		fields  []string
	}
	tests := []struct {
		name string
		args args
		want []Daira
	}{
		{
			name: "filterDairats with valid fields",
			args: args{
				dairats: []Daira{
					{
						Name: "ADRAR",
						Code: 1,
					},
					{
						Name: "OULED AHMED TIMMI",
						Code: 2,
					},
				},
				fields: []string{
					"name",
				},
			},
			want: []Daira{
				{
					Name: "ADRAR",
				},
				{
					Name: "OULED AHMED TIMMI",
				},
			},
		},
		{
			name: "filterDairats with invalid fields",
			args: args{
				dairats: []Daira{
					{
						Name: "ADRAR",
					},
					{
						Name: "OULED AHMED TIMMI",
					},
				},
				fields: []string{
					"invalid",
				},
			},
			want: []Daira{
				{}, {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterDairats(tt.args.dairats, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterDairats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPhoneCode(t *testing.T) {
	type args struct {
		phoneCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isValidPhoneCode with valid phone code",
			args: args{
				phoneCode: 21,
			},
			want: true,
		},
		{
			name: "isValidPhoneCode with invalid phone code",
			args: args{
				phoneCode: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPhoneCode(tt.args.phoneCode); got != tt.want {
				t.Errorf("isValidPhoneCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWilayaIndexByPhoneCode(t *testing.T) {
	type args struct {
		wilayas   *[]Wilaya
		phoneCode int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getWilayaIndexByPhoneCode with valid phone code",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "Adrar",
						PhoneCodes: []int{
							21,
						},
					},
					{
						Name: "Chlef",
						PhoneCodes: []int{
							23,
						},
					},
				},
				phoneCode: 21,
			},
			want: 0,
		},
		{
			name: "getWilayaIndexByPhoneCode with invalid phone code",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "Adrar",
						PhoneCodes: []int{
							21,
						},
					},
					{
						Name: "Chlef",
						PhoneCodes: []int{
							23,
						},
					},
				},
				phoneCode: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWilayaIndexByPhoneCode(tt.args.wilayas, tt.args.phoneCode); got != tt.want {
				t.Errorf("getWilayaIndexByPhoneCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWilayaIndexByDairaName(t *testing.T) {
	type args struct {
		wilayas   *[]Wilaya
		dairaName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getWilayaIndexByDairaName with valid daira name",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "Adrar",
						Dairats: []Daira{
							{
								Name: "ADRAR",
							},
						},
					},
					{
						Name: "Chlef",
						Dairats: []Daira{
							{
								Name: "CHLEF",
							},
						},
					},
				},
				dairaName: "ADRAR",
			},
			want: 0,
		},
		{
			name: "getWilayaIndexByDairaName with invalid daira name",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "Adrar",
						Dairats: []Daira{
							{
								Name: "ADRAR",
							},
						},
					},
					{
						Name: "Chlef",
						Dairats: []Daira{
							{
								Name: "CHLEF",
							},
						},
					},
				},
				dairaName: "invalid",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWilayaIndexByDairaName(tt.args.wilayas, tt.args.dairaName); got != tt.want {
				t.Errorf("getWilayaIndexByDairaName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDairaIndexByName(t *testing.T) {
	type args struct {
		dairats   []Daira
		dairaName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getDairaIndexByName with valid daira name",
			args: args{
				dairats: []Daira{
					{
						Name: "ADRAR",
					},
				},
				dairaName: "ADRAR",
			},
			want: 0,
		},
		{
			name: "getDairaIndexByName with invalid daira name",
			args: args{
				dairats: []Daira{
					{
						Name: "ADRAR",
					},
				},
				dairaName: "invalid",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDairaIndexByName(tt.args.dairats, tt.args.dairaName); got != tt.want {
				t.Errorf("getDairaIndexByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBaladyiats(t *testing.T) {
	type args struct {
		dairats    []Daira
		dairaIndex int
	}
	tests := []struct {
		name string
		args args
		want []Baladyia
	}{
		{
			name: "getBaladyiats with valid daira index",
			args: args{
				dairats: []Daira{
					{
						Name: "ADRAR",
						Baladyiats: []Baladyia{
							{
								Name: "ADRAR",
							},
						},
					},
				},
				dairaIndex: 0,
			},
			want: []Baladyia{
				{
					Name: "ADRAR",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBaladyiats(tt.args.dairats, tt.args.dairaIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBaladyiats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterBaladyia(t *testing.T) {
	type args struct {
		b      Baladyia
		fields []string
	}
	tests := []struct {
		name string
		args args
		want Baladyia
	}{
		{
			name: "filterBaladyia with valid fields",
			args: args{
				b: Baladyia{
					Name: "ADRAR",
				},
				fields: []string{"name"},
			},
			want: Baladyia{
				Name: "ADRAR",
			},
		},
		{
			name: "filterBaladyia with invalid fields",
			args: args{
				b: Baladyia{
					Name: "ADRAR",
				},
				fields: []string{"invalid"},
			},
			want: Baladyia{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterBaladyia(tt.args.b, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterBaladyia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterBaladyiats(t *testing.T) {
	type args struct {
		baladyiats []Baladyia
		fields     []string
	}
	tests := []struct {
		name string
		args args
		want []Baladyia
	}{
		{
			name: "filterBaladyiats with valid fields",
			args: args{
				baladyiats: []Baladyia{
					{
						Name: "ADRAR",
					},
				},
				fields: []string{"name"},
			},
			want: []Baladyia{
				{
					Name: "ADRAR",
				},
			},
		},
		{
			name: "filterBaladyiats with invalid fields",
			args: args{
				baladyiats: []Baladyia{
					{
						Name: "ADRAR",
					},
				},
				fields: []string{"invalid"},
			},
			want: []Baladyia{
				{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterBaladyiats(tt.args.baladyiats, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterBaladyiats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWilayaIndexByDairaCode(t *testing.T) {
	type args struct {
		wilayas   *[]Wilaya
		dairaCode int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getWilayaIndexByDairaCode with valid daira code",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "ADRAR",
						Dairats: []Daira{
							{
								Code: 1,
							},
						},
					},
				},
				dairaCode: 1,
			},
			want: 0,
		},
		{
			name: "getWilayaIndexByDairaCode with invalid daira code",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "ADRAR",
						Dairats: []Daira{
							{
								Code: 1,
							},
						},
					},
				},
				dairaCode: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWilayaIndexByDairaCode(tt.args.wilayas, tt.args.dairaCode); got != tt.want {
				t.Errorf("getWilayaIndexByDairaCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDairaIndexByCode(t *testing.T) {
	type args struct {
		dairats   []Daira
		dairaCode int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getDairaIndexByCode with valid daira code",
			args: args{
				dairats: []Daira{
					{
						Code: 1,
					},
				},
				dairaCode: 1,
			},
			want: 0,
		},
		{
			name: "getDairaIndexByCode with invalid daira code",
			args: args{
				dairats: []Daira{
					{
						Code: 1,
					},
				},
				dairaCode: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDairaIndexByCode(tt.args.dairats, tt.args.dairaCode); got != tt.want {
				t.Errorf("getDairaIndexByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWilayaIndexByName(t *testing.T) {
	type args struct {
		wilayas    *[]Wilaya
		wilayaName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getWilayaIndexByName with valid wilaya name",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "ADRAR",
					},
				},
				wilayaName: "ADRAR",
			},
			want: 0,
		},
		{
			name: "getWilayaIndexByName with invalid wilaya name",
			args: args{
				wilayas: &[]Wilaya{
					{
						Name: "ADRAR",
					},
				},
				wilayaName: "invalid",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWilayaIndexByName(tt.args.wilayas, tt.args.wilayaName); got != tt.want {
				t.Errorf("getWilayaIndexByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPhoneCodes(t *testing.T) {
	type args struct {
		wilaya Wilaya
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "getPhoneCodes with valid wilaya",
			args: args{
				wilaya: Wilaya{
					PhoneCodes: []int{21, 22},
				},
			},
			want: []int{21, 22},
		},
		{
			name: "getPhoneCodes with empty phone codes",
			args: args{
				wilaya: Wilaya{
					PhoneCodes: []int{},
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPhoneCodes(tt.args.wilaya); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPhoneCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leblad

import (
	"encoding/json"
	"os"
)

// openJsonFile opens a local json file with the given name
func openJsonFile(n string) ([]byte, error) {
	bytes, err := os.ReadFile(n)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// unmarshalWilayaListJson unmarshals the given json bytes into a slice of wilayas
func unmarshalWilayaListJson(b []byte) (*[]Wilaya, error) {
	var wilayas *[]Wilaya
	if err := json.Unmarshal(b, &wilayas); err != nil {
		return nil, err
	}
	return wilayas, nil
}

// filterWilayaList filters the fields of the given slice of wilayas
func filterWilayaList(wilayas *[]Wilaya, fields ...string) *[]Wilaya {
	var filtered []Wilaya
	for _, w := range *wilayas {
		var f Wilaya
		for _, field := range fields {
			switch field {
			case "matricule":
				f.Matricule = w.Matricule
			case "name_ar":
				f.NameAr = w.NameAr
			case "name_ber":
				f.NameBer = w.NameBer
			case "name_en":
				f.NameEn = w.NameEn
			case "name":
				f.Name = w.Name
			case "phoneCodes":
				f.PhoneCodes = w.PhoneCodes
			case "postalCodes":
				f.PostalCodes = w.PostalCodes
			case "dairats":
				f.Dairats = w.Dairats
			case "adjacentWilayas":
				f.AdjacentWilayas = w.AdjacentWilayas
			}
		}
		filtered = append(filtered, f)
	}
	return &filtered
}

package leblad

import (
	"encoding/json"
	"os"
)

const (
	ZIP_COUNT    = 48073
	WILAYA_COUNT = 48
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
		filtered = append(filtered, filterWilaya(w, fields...))
	}
	return &filtered
}

// filterWilaya filters the fields of the given wilaya
func filterWilaya(w Wilaya, fields ...string) Wilaya {
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
	return f
}

// getWilayaIndexByZipCode returns the index of the wilaya that contains the given zip code
func getWilayaIndexByZipCode(wilayas *[]Wilaya, zipCode int) int {
	for i, w := range *wilayas {
		for _, pc := range w.PostalCodes {
			if pc == zipCode {
				return i
			}
		}
	}
	return -1
}

// getWilayaIndexByCode returns the index of the wilaya that has the given code
func getWilayaIndexByCode(wilayas *[]Wilaya, wilayaCode int) int {
	for i, w := range *wilayas {
		if w.Matricule == wilayaCode {
			return i
		}
	}
	return -1
}

// getAdjacentWilayas returns the adjacent wilayas of the given wilaya
func getAdjacentWilayas(wilayas *[]Wilaya, wilayaIndex int) []int {
	return (*wilayas)[wilayaIndex].AdjacentWilayas
}

// getZipCodes returns the zip codes of the given wilaya
func getZipCodes(wilayas *[]Wilaya, wilayaIndex int) []int {
	return (*wilayas)[wilayaIndex].PostalCodes
}

// getDairats returns the dairats of the given wilaya
func getDairats(wilayas *[]Wilaya, wilayaIndex int) []Daira {
	return (*wilayas)[wilayaIndex].Dairats
}

// isValidZipCode returns true if the given zip code is valid
func isValidZipCode(zipCode int) bool {
	return zipCode >= 1000 && zipCode <= ZIP_COUNT
}

// isValidWilayaCode returns true if the given wilaya code is valid
func isValidWilayaCode(wilayaCode int) bool {
	return wilayaCode >= 1 && wilayaCode <= WILAYA_COUNT
}

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

// filterDairats filters the fields of the given slice of dairats
func filterDairats(dairats []Daira, fields ...string) []Daira {
	var filtered []Daira
	for _, d := range dairats {
		filtered = append(filtered, filterDaira(d, fields...))
	}
	return filtered
}

// filterDaira filters the fields of the given daira
func filterDaira(d Daira, fields ...string) Daira {
	var f Daira
	for _, field := range fields {
		switch field {
		case "code":
			f.Code = d.Code
		case "name":
			f.Name = d.Name
		case "name_ar":
			f.NameAr = d.NameAr
		case "name_en":
			f.NameEn = d.NameEn
		case "baladyiats":
			f.Baladyiats = d.Baladyiats
		}
	}
	return f
}

// filterBaladyiats filters the fields of the given slice of baladyiats
func filterBaladyiats(baladyiats []Baladyia, fields ...string) []Baladyia {
	var filtered []Baladyia
	for _, b := range baladyiats {
		filtered = append(filtered, filterBaladyia(b, fields...))
	}
	return filtered
}

// filterBaladyia filters the fields of the given baladyia
func filterBaladyia(b Baladyia, fields ...string) Baladyia {
	var f Baladyia
	for _, field := range fields {
		switch field {
		case "code":
			f.Code = b.Code
		case "name":
			f.Name = b.Name
		case "name_ar":
			f.NameAr = b.NameAr
		case "name_en":
			f.NameEn = b.NameEn
		}
	}
	return f
}

// getWilayaIndexByPhoneCode returns the index of the wilaya that contains the given phone code
func getWilayaIndexByPhoneCode(wilayas *[]Wilaya, phoneCode int) int {
	for i, w := range *wilayas {
		for _, c := range w.PhoneCodes {
			if c == phoneCode {
				return i
			}
		}
	}
	return -1
}

// getWilayaIndexByDairaName returns the index of the wilaya that contains the given daira name
func getWilayaIndexByDairaName(wilayas *[]Wilaya, dairaName string) int {
	for i, w := range *wilayas {
		for _, d := range w.Dairats {
			if d.Name == dairaName {
				return i
			}
		}
	}
	return -1
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

// getBaladyiats returns the baladyiats of the given daira
func getBaladyiats(dairats []Daira, dairaIndex int) []Baladyia {
	return dairats[dairaIndex].Baladyiats
}

// getDairaIndexByName returns the index of the daira that has the given name
func getDairaIndexByName(dairats []Daira, dairaName string) int {
	for i, d := range dairats {
		if d.Name == dairaName {
			return i
		}
	}
	return -1
}

// getWilayaIndexByDairaCode returns the index of the wilaya that contains the given daira code
func getWilayaIndexByDairaCode(wilayas *[]Wilaya, dairaCode int) int {
	for i, w := range *wilayas {
		for _, d := range w.Dairats {
			if d.Code == dairaCode {
				return i
			}
		}
	}
	return -1
}

// getDairaIndexByCode returns the index of the daira that has the given code
func getDairaIndexByCode(dairats []Daira, dairaCode int) int {
	for i, d := range dairats {
		if d.Code == dairaCode {
			return i
		}
	}
	return -1
}

// getWilayaIndexByName returns the index of the wilaya that has the given name
func getWilayaIndexByName(wilayas *[]Wilaya, wilayaName string) int {
	for i, w := range *wilayas {
		if w.Name == wilayaName {
			return i
		}
	}
	return -1
}

// getPhoneCodes returns the phone codes of the given wilaya
func getPhoneCodes(wilaya Wilaya) []int {
	return wilaya.PhoneCodes
}

// getFirstPhoneCode returns the first phone code of the given wilaya
func getFirstPhoneCode(wilaya Wilaya) int {
	return wilaya.PhoneCodes[0]
}

// isValidZipCode returns true if the given zip code is valid
func isValidZipCode(zipCode int) bool {
	return zipCode >= 1000 && zipCode <= ZIP_COUNT
}

// isValidWilayaCode returns true if the given wilaya code is valid
func isValidWilayaCode(wilayaCode int) bool {
	return wilayaCode >= 1 && wilayaCode <= WILAYA_COUNT
}

// isValidPhoneCode returns true if the given phone code is valid
func isValidPhoneCode(phoneCode int) bool {
	return phoneCode >= 21 && phoneCode <= 59
}

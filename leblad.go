package leblad

import (
	"path/filepath"
	"runtime"
)

var (
	f       string
	dirPath string
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	f = file
	dirPath = filepath.Dir(f)
}

// Leblad is the main struct for the Leblad library
type Leblad struct{}

// New returns a new Leblad struct
func New() *Leblad {
	return &Leblad{}
}

// GetWilayaList returns a slice of all wilayas.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaList(fields ...string) ([]Wilaya, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaListError{}
	}
	// filter the results
	if len(fields) > 0 {
		wilayas = filterWilayaList(wilayas, fields...)
	}
	return *wilayas, nil
}

// GetWilayaByZipCode returns a wilaya by its zip code.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaByZipCode(zipCode int, fields ...string) (*Wilaya, error) {
	// check if the zip code is valid
	if !isValidZipCode(zipCode) {
		return nil, &WilayaByZipCodeError{zipCode}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaListError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByZipCode(wilayas, zipCode)
	if index == -1 {
		return nil, &WilayaByZipCodeError{zipCode}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return &w, nil
	}
	return &(*wilayas)[index], nil
}

// GetWilayaByCode returns a wilaya by its matricule.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaByCode(matricule int, fields ...string) (*Wilaya, error) {
	// check if the matricule is valid
	if !isValidWilayaCode(matricule) {
		return nil, &WilayaByCodeError{matricule}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaListError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByCode(wilayas, matricule)
	if index == -1 {
		return nil, &WilayaByCodeError{matricule}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return &w, nil
	}
	return &(*wilayas)[index], nil
}

// GetAdjacentWilayas returns a slice of adjacent wilayas by a given wilaya code
func (l *Leblad) GetAdjacentWilayas(matricule int) ([]int, error) {
	// check if the matricule is valid
	if !isValidWilayaCode(matricule) {
		return nil, &WilayaByCodeError{matricule}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &AdjacentWilayasError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &AdjacentWilayasError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByCode(wilayas, matricule)
	if index == -1 {
		return nil, &WilayaByCodeError{matricule}
	}
	// get the adjacent wilayas
	adjacentWilayas := getAdjacentWilayas(wilayas, index)
	return adjacentWilayas, nil
}

// GetZipCodesForWilaya returns a slice of zip codes for a given wilaya code
func (l *Leblad) GetZipCodesForWilaya(matricule int) ([]int, error) {
	// check if the matricule is valid
	if !isValidWilayaCode(matricule) {
		return nil, &WilayaByCodeError{matricule}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &ZipCodesForWilayaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &ZipCodesForWilayaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByCode(wilayas, matricule)
	if index == -1 {
		return nil, &WilayaByCodeError{matricule}
	}
	// get the zip codes
	zipCodes := getZipCodes(wilayas, index)
	return zipCodes, nil
}

// GetDairatsForWilaya returns a slice of dairats for a given wilaya code.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetDairatsForWilaya(matricule int, fields ...string) ([]Daira, error) {
	// check if the matricule is valid
	if !isValidWilayaCode(matricule) {
		return nil, &WilayaByCodeError{matricule}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &DairatsForWilayaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &DairatsForWilayaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByCode(wilayas, matricule)
	if index == -1 {
		return nil, &WilayaByCodeError{matricule}
	}
	// get the dairats
	dairats := getDairats(wilayas, index)
	// filter the results
	if len(fields) > 0 {
		d := filterDairats(dairats, fields...)
		return d, nil
	}
	return dairats, nil
}

// GetWilayaByPhoneCode returns a wilaya by its phone code.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaByPhoneCode(phoneCode int, fields ...string) (*Wilaya, error) {
	// check if the phone code is valid
	if !isValidPhoneCode(phoneCode) {
		return nil, &WilayaByPhoneCodeError{phoneCode}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaListError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByPhoneCode(wilayas, phoneCode)
	if index == -1 {
		return nil, &WilayaByPhoneCodeError{phoneCode}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return &w, nil
	}
	return &(*wilayas)[index], nil
}

// GetWilayaByDairaName returns a wilaya by its daira name.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaByDairaName(dairaName string, fields ...string) (*Wilaya, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaListError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByDairaName(wilayas, dairaName)
	if index == -1 {
		return nil, &WilayaByDairaNameError{dairaName}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return &w, nil
	}
	return &(*wilayas)[index], nil
}

// GetBaladyiatsForDaira returns a slice of baladyiats for a given daira name.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetBaladyiatsForDaira(dairaName string, fields ...string) ([]Baladyia, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &BaladyiatsForDairaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &BaladyiatsForDairaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByDairaName(wilayas, dairaName)
	if index == -1 {
		return nil, &WilayaByDairaNameError{dairaName}
	}
	// get the index of the daira
	dairaIndex := getDairaIndexByName((*wilayas)[index].Dairats, dairaName)
	if dairaIndex == -1 {
		return nil, &DairaByDairaNameError{dairaName}
	}
	// get the baladyiats
	baladyiats := getBaladyiats((*wilayas)[index].Dairats, dairaIndex)
	// filter the results
	if len(fields) > 0 {
		b := filterBaladyiats(baladyiats, fields...)
		return b, nil
	}
	return baladyiats, nil
}

// GetBaladyiatsForDairaCode returns a slice of baladyiats for a given daira code.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetBaladyiatsForDairaCode(dairaCode int, fields ...string) ([]Baladyia, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &BaladyiatsForDairaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &BaladyiatsForDairaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByDairaCode(wilayas, dairaCode)
	if index == -1 {
		return nil, &WilayaByDairaCodeError{dairaCode}
	}
	// get the index of the daira
	dairaIndex := getDairaIndexByCode((*wilayas)[index].Dairats, dairaCode)
	if dairaIndex == -1 {
		return nil, &DairaByDairaCodeError{dairaCode}
	}
	// get the baladyiats
	baladyiats := getBaladyiats((*wilayas)[index].Dairats, dairaIndex)
	// filter the results
	if len(fields) > 0 {
		b := filterBaladyiats(baladyiats, fields...)
		return b, nil
	}
	return baladyiats, nil
}

// GetPhoneCodesForWilaya returns a slice of phone codes for a given wilaya name.
func (l *Leblad) GetPhoneCodesForWilaya(wilayaName string) ([]int, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &PhoneCodesForWilayaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &PhoneCodesForWilayaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByName(wilayas, wilayaName)
	if index == -1 {
		return nil, &WilayaByWilayaNameError{wilayaName}
	}
	// get the phone codes
	phoneCodes := getPhoneCodes((*wilayas)[index])
	return phoneCodes, nil
}

// GetFirstPhoneCodeForWilaya returns the first phone code for a given wilaya name.
func (l *Leblad) GetFirstPhoneCodeForWilaya(wilayaName string) (int, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return 0, &FirstPhoneCodeForWilayaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return 0, &FirstPhoneCodeForWilayaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByName(wilayas, wilayaName)
	if index == -1 {
		return 0, &WilayaByWilayaNameError{wilayaName}
	}
	// get the first phone code
	phoneCode := getFirstPhoneCode((*wilayas)[index])
	return phoneCode, nil
}

// GetBaladyiatsForWilaya returns a slice of baladyiats for a given wilaya name.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetBaladyiatsForWilaya(wilayaName string, fields ...string) ([]Baladyia, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &BaladyiatsForWilayaError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &BaladyiatsForWilayaError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByName(wilayas, wilayaName)
	if index == -1 {
		return nil, &WilayaByWilayaNameError{wilayaName}
	}
	// get the baladyiats
	baladyiats := getBaladyiatsForWilaya((*wilayas)[index])
	// filter the results
	if len(fields) > 0 {
		b := filterBaladyiats(baladyiats, fields...)
		return b, nil
	}
	return baladyiats, nil
}

// GetWilayaByBaladyiaName returns a wilaya for a given baladyia name.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaByBaladyiaName(baladyiaName string, fields ...string) (*Wilaya, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaByBaladyiaNameError{baladyiaName}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaByBaladyiaNameError{baladyiaName}
	}
	// get the index of the wilaya
	index := getWilayaIndexByBaladyiaName(wilayas, baladyiaName)
	if index == -1 {
		return nil, &WilayaByBaladyiaNameError{}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return &w, nil
	}
	return &(*wilayas)[index], nil
}

// GetDairaByBaladyiaName returns a daira for a given baladyia name.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetDairaByBaladyiaName(baladyiaName string, fields ...string) (*Daira, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &DairaByBaladyiaNameError{baladyiaName}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &DairaByBaladyiaNameError{baladyiaName}
	}
	// get the index of the wilaya
	index := getWilayaIndexByBaladyiaName(wilayas, baladyiaName)
	if index == -1 {
		return nil, &WilayaByBaladyiaNameError{baladyiaName}
	}
	// get the index of the daira
	dairaIndex := getDairaIndexByBaladyiaName((*wilayas)[index].Dairats, baladyiaName)
	if dairaIndex == -1 {
		return nil, &DairaByBaladyiaNameError{baladyiaName}
	}
	// filter the results
	if len(fields) > 0 {
		d := filterDaira((*wilayas)[index].Dairats[dairaIndex], fields...)
		return &d, nil
	}
	return &(*wilayas)[index].Dairats[dairaIndex], nil
}

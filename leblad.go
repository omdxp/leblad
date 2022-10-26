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
func (l *Leblad) GetWilayaByZipCode(zipCode int, fields ...string) (Wilaya, error) {
	// check if the zip code is valid
	if !isValidZipCode(zipCode) {
		return Wilaya{}, &WilayaByZipCodeError{zipCode}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return Wilaya{}, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return Wilaya{}, &WilayaListError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByZipCode(wilayas, zipCode)
	if index == -1 {
		return Wilaya{}, &WilayaByZipCodeError{zipCode}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return w, nil
	}
	return (*wilayas)[index], nil
}

// GetWilayaByCode returns a wilaya by its matricule.
// It has a variadic argument that can be used to filter the results
func (l *Leblad) GetWilayaByCode(matricule int, fields ...string) (Wilaya, error) {
	// check if the matricule is valid
	if !isValidWilayaCode(matricule) {
		return Wilaya{}, &WilayaByCodeError{matricule}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return Wilaya{}, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return Wilaya{}, &WilayaListError{}
	}
	// get the index of the wilaya
	index := getWilayaIndexByCode(wilayas, matricule)
	if index == -1 {
		return Wilaya{}, &WilayaByCodeError{matricule}
	}
	// filter the results
	if len(fields) > 0 {
		w := filterWilaya((*wilayas)[index], fields...)
		return w, nil
	}
	return (*wilayas)[index], nil
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

// GetDairatsForWilaya returns a slice of dairats for a given wilaya code
func (l *Leblad) GetDairatsForWilaya(matricule int) ([]Daira, error) {
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
	return dairats, nil
}

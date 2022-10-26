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

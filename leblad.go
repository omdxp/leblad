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

// GetWilayaList returns a slice of all wilayas
func (l *Leblad) GetWilayaList() ([]Wilaya, error) {
	bytes, err := openJsonFile(filepath.Join(dirPath, "data", "WilayaList.json"))
	if err != nil {
		return nil, &WilayaListError{}
	}
	wilayas, err := unmarshalWilayaListJson(bytes)
	if err != nil {
		return nil, &WilayaListError{}
	}
	return *wilayas, nil
}

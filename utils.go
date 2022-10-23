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

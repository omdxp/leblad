package leblad

// WilayaListError is the error returned when getting the wilaya list
type WilayaListError struct{}

// Error returns the error message
func (e *WilayaListError) Error() string {
	return "error getting wilaya list"
}

// WilayaByZipCodeError is the error returned when getting a wilaya by zip code
type WilayaByZipCodeError struct{}

// Error returns the error message
func (e *WilayaByZipCodeError) Error() string {
	return "error getting wilaya by zip code"
}

// WilayaByCodeError is the error returned when getting a wilaya by code
type WilayaByCodeError struct{}

// Error returns the error message
func (e *WilayaByCodeError) Error() string {
	return "error getting wilaya by code"
}

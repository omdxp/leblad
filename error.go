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

// AdjacentWilayasError is the error returned when getting adjacent wilayas
type AdjacentWilayasError struct{}

// Error returns the error message
func (e *AdjacentWilayasError) Error() string {
	return "error getting adjacent wilayas"
}

// ZipCodesForWilayaError is the error returned when getting zip codes for a wilaya
type ZipCodesForWilayaError struct{}

// Error returns the error message
func (e *ZipCodesForWilayaError) Error() string {
	return "error getting zip codes for wilaya"
}

// WilayaByPhoneCodeError is the error returned when getting a wilaya by phone code
type WilayaByPhoneCodeError struct{}

// Error returns the error message
func (e *WilayaByPhoneCodeError) Error() string {
	return "error getting wilaya by phone code"
}

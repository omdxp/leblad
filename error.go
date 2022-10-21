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

// WilayaByDairaError is the error returned when getting a wilaya by daira
type WilayaByDairaError struct{}

// Error returns the error message
func (e *WilayaByDairaError) Error() string {
	return "error getting wilaya by daira"
}

// BaladiyatsForDairaError is the error returned when getting baladiyats for a daira
type BaladiyatsForDairaError struct{}

// Error returns the error message
func (e *BaladiyatsForDairaError) Error() string {
	return "error getting baladiyats for daira"
}

// DairatsForWilayaError is the error returned when getting dairats for a wilaya
type DairatsForWilayaError struct{}

// Error returns the error message
func (e *DairatsForWilayaError) Error() string {
	return "error getting dairats for wilaya"
}

// PhoneCodesForWilayaError is the error returned when getting phone codes for a wilaya
type PhoneCodesForWilayaError struct{}

// Error returns the error message
func (e *PhoneCodesForWilayaError) Error() string {
	return "error getting phone codes for wilaya"
}

// FirstPhoneCodeForWilayaError is the error returned when getting the first phone code for a wilaya
type FirstPhoneCodeForWilayaError struct{}

// Error returns the error message
func (e *FirstPhoneCodeForWilayaError) Error() string {
	return "error getting first phone code for wilaya"
}

// BaladiyatsForWilayaError is the error returned when getting baladiyats for a wilaya
type BaladiyatsForWilayaError struct{}

// Error returns the error message
func (e *BaladiyatsForWilayaError) Error() string {
	return "error getting baladiyats for wilaya"
}

// WilayaByBaladiyaError is the error returned when getting a wilaya by baladiya
type WilayaByBaladiyaError struct{}

// Error returns the error message
func (e *WilayaByBaladiyaError) Error() string {
	return "error getting wilaya by baladiya"
}

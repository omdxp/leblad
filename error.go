package leblad

import "strconv"

// WilayaListError is the error returned when getting the wilaya list
type WilayaListError struct{}

// Error returns the error message
func (e *WilayaListError) Error() string {
	return "error getting wilaya list"
}

// WilayaByZipCodeError is the error returned when getting a wilaya by zip code
type WilayaByZipCodeError struct {
	ZipCode int
}

// Error returns the error message
func (e *WilayaByZipCodeError) Error() string {
	return "error getting wilaya by zip code " + strconv.Itoa(e.ZipCode)
}

// WilayaByCodeError is the error returned when getting a wilaya by code
type WilayaByCodeError struct {
	Code int
}

// Error returns the error message
func (e *WilayaByCodeError) Error() string {
	return "error getting wilaya by code " + strconv.Itoa(e.Code)
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
type WilayaByPhoneCodeError struct {
	PhoneCode int
}

// Error returns the error message
func (e *WilayaByPhoneCodeError) Error() string {
	return "error getting wilaya by phone code " + strconv.Itoa(e.PhoneCode)
}

// WilayaByDairaError is the error returned when getting a wilaya by daira
type WilayaByDairaNameError struct {
	DairaName string
}

// Error returns the error message
func (e *WilayaByDairaNameError) Error() string {
	return "error getting wilaya by daira " + e.DairaName
}

// BaladiyatsForDairaError is the error returned when getting baladiyats for a daira
type BaladyiatsForDairaError struct{}

// Error returns the error message
func (e *BaladyiatsForDairaError) Error() string {
	return "error getting baladiyats for daira"
}

// DairatsForWilayaError is the error returned when getting dairats for a wilaya
type DairatsForWilayaError struct{}

// Error returns the error message
func (e *DairatsForWilayaError) Error() string {
	return "error getting dairats for wilaya"
}

// WilayaByDairaCodeError is the error returned when getting a wilaya by daira code
type WilayaByDairaCodeError struct {
	DairaCode int
}

// Error returns the error message
func (e *WilayaByDairaCodeError) Error() string {
	return "error getting wilaya by daira code " + strconv.Itoa(e.DairaCode)
}

// DairaByDairaCodeError is the error returned when getting a daira by daira code
type DairaByDairaCodeError struct {
	DairaCode int
}

// Error returns the error message
func (e *DairaByDairaCodeError) Error() string {
	return "error getting daira by daira code " + strconv.Itoa(e.DairaCode)
}

// DairaByDairaNameError is the error returned when getting a daira by daira name
type DairaByDairaNameError struct {
	DairaName string
}

// Error returns the error message
func (e *DairaByDairaNameError) Error() string {
	return "error getting daira by daira name " + e.DairaName
}

// WilayaByWilayaNameError is the error returned when getting a wilaya by wilaya name
type WilayaByWilayaNameError struct {
	WilayaName string
}

// Error returns the error message
func (e *WilayaByWilayaNameError) Error() string {
	return "error getting wilaya by wilaya name " + e.WilayaName
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

// DairaByBaladiyaError is the error returned when getting a daira by baladiya
type DairaByBaladiyaError struct{}

// Error returns the error message
func (e *DairaByBaladiyaError) Error() string {
	return "error getting daira by baladiya"
}

// FullAdjacentWilayasError is the error returned when getting full adjacent wilayas
type FullAdjacentWilayasError struct{}

// Error returns the error message
func (e *FullAdjacentWilayasError) Error() string {
	return "error getting full adjacent wilayas"
}

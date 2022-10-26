package leblad

// Wilaya is a struct that represents a wilaya
type Wilaya struct {
	Matricule       int     `json:"mattricule"`
	NameAr          string  `json:"name_ar"`
	NameBer         string  `json:"name_ber"`
	NameEn          string  `json:"name_en"`
	Name            string  `json:"name"`
	PhoneCodes      []int   `json:"phoneCodes"`
	PostalCodes     []int   `json:"postalCodes"`
	Dairats         []Daira `json:"dairats"`
	AdjacentWilayas []int   `json:"adjacentWilayas"`
}

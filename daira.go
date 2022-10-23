package leblad

// Daira is a struct that represents a daira
type Daira struct {
	Code       int        `json:"code"`
	Name       string     `json:"name"`
	NameAr     string     `json:"name_ar"`
	NameEn     string     `json:"name_en"`
	Baladyiats []Baladyia `json:"baladyiats"`
}

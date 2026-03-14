package ranges

type rangeAddRequest struct {
	Name     string  `json:"name"`
	Address1 string  `json:"address1"`
	Address2 string  `json:"address2"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	Zipcode  string  `json:"zipcode"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}

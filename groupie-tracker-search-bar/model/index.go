package model


type index struct{
	ID int    `json:"id"`
	Relations  map[string][]string `json:"datesLocation"`
}

package model

type DatesLocations struct {
    ID             int                 `json:"id"`
    LocationsDates map[string][]string `json:"datesLocations"`
}
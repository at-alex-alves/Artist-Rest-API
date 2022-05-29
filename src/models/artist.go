package models

type Artist struct {
	Id             string     `json:"id,omitempty"`
	FirstName      string     `json:"firstName,omitempty"`
	LastName       string     `json:"lastName,omitempty"`
	MostFamousWork Artwork    `json:"mostFamousWork,omitempty"`
	Birthplace     Birthplace `json:"birthplace,omitempty"`
}

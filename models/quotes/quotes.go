package models

type Authors struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string   `json:"name,omitempty" bson:"name,omitempty"`
	Url      string   `json:"url,omitempty" bson:"url,omitempty"`
	Category []string `json:"category" bson:"category"`
}

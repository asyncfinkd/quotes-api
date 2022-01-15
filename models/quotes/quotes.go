package models

type Authors struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string   `json:"name,omitempty" bson:"name,omitempty"`
	Url      string   `json:"url,omitempty" bson:"url,omitempty"`
	Category []string `json:"category" bson:"category"`
}

type Quotes struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Text     string   `json:"text,omitempty" bson:"text,omitempty"`
	Author   string   `json:"author,omitempty" bson:"author,omitempty"`
	Category []string `json:"category" bson:"category,omitempty"`
}

type Users struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

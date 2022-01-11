package constant

type Quotes struct {
	ID       uint     `json:"id"`
	Text     string   `json:"text"`
	Author   string   `json:"author"`
	Category []string `json:"category"`
}

type AuthorGallery struct {
	ID       uint     `json:"id"`
	Url      string   `json:"url"`
	Category []string `json:"category"`
	Name     string   `json:"name"`
}

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

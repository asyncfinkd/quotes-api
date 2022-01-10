package constant

type Quotes struct {
	ID       uint     `json:"id" validate:"required,omitempty"`
	Text     string   `json:"text" validate:"required"`
	Author   string   `json:"author" validate:"required"`
	Category []string `json:"category" validate:"category"`
}

type AuthorGallery struct {
	ID       uint     `json:"id" validate:"required,omitempty"`
	Url      string   `json:"url" validate:"required"`
	Category []string `json:"category" validate:"required"`
	Name     string   `json:"name" validate:"required"`
}

type User struct {
	ID       uint   `json:"id" validate:"required,omitempty"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,gte=10"`
	Role     string `json:"role" validate:"required"`
}

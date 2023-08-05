package models

type RegisterInput struct {
	FullName string `json:"full_name" validate:"required,min=4,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=25"`
}

type RegisterResponse struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RegisterService interface {
	Create(user *User) error
	FetchByEmail(email string) (User, error)
}

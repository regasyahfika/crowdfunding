package user

type RegisterUserInput struct {
	Name           string `json:"name" binding:"required"`
	Occupation     string `json:"occupation" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Role           string `json:"role" binding:"required"`
	Password       string `json:"password" binding:"required"`
	AvatarFileName string
}

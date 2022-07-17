package request

type ExistEmail struct {
	Email string `json:"email" binding:"required,email,lte=50" maximum:"50"`
}

package request

type ExistEmail struct {
	Email string `form:"email" binding:"required,email,lte=50" maximum:"50"` // 邮箱
}

type SendEmail struct {
	Email string `json:"email" binding:"required,email,lte=50" maximum:"50"` // 邮箱
}

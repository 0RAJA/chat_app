package request

type Register struct {
	Email    string `json:"email" binding:"required,email,lte=50" maximum:"50"`                // 邮箱
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"` // 密码
	Code     string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"`       // 验证码
}

type Login struct {
	Email    string `json:"email" binding:"required,email,lte=50" maximum:"50"`                // 邮箱
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"` // 密码
}

type UpdateUserEmail struct {
	Email string `json:"email" binding:"required,email,lte=50" maximum:"50"`          // 邮箱
	Code  string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"` // 验证码
}

type UpdateUserPassword struct {
	Code        string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"`           // 验证码
	NewPassword string `json:"new_password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"` // 新密码
}

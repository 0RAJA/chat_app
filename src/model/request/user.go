package request

type CreateUser struct {
	Email    string `json:"email" binding:"required,email,lte=50" maximum:"50"`
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
	Code     string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email,lte=50" maximum:"50"`
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
}

type UpdateUserEmail struct {
	Email string `json:"email" binding:"required,email,lte=50" maximum:"50"`
	Code  string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"`
}

type UpdateUserPassword struct {
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
}

type ExistEmail struct {
	Email string `json:"email" binding:"required,email,lte=50" maximum:"50"`
}

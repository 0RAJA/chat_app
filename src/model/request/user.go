package request

type Register struct {
	Email    string `json:"email" binding:"required,email,lte=50" maximum:"50"`
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
	Code     string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email,lte=50" maximum:"50"`
	Password string `json:"password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
}

type UpdateUserEmail struct {
	Email string `json:"email" binding:"required,email,lte=50" maximum:"50"`
	Code  string `json:"code" binding:"required,gte=6,lte=6" minimum:"6" maximum:"6"`
}

type UpdateUserPassword struct {
	OldPassword string `json:"old_password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
	NewPassword string `json:"new_password" binding:"required,gte=6,lte=50" minimum:"6" maximum:"50"`
}

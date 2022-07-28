package v1

type group struct {
	User        user
	Email       email
	Account     account
	Application application
	Setting     setting
	Message     message
}

var Group = new(group)

package routing

type group struct {
	User  user
	Email email
	File   file
	Account     account
	Application application
	Setting     setting
	Message     message
	Notify notify
}

var Group = new(group)

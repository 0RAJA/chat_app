package v1

type group struct {
	User  user
	Email email
	File file
	Account account
	Application application
	Notify notify
	Setting     setting
	Message     message
}

var Group = new(group)

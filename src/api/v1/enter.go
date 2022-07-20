package v1

type group struct {
	User  user
	Email email
	File file
	Account     account
	Application application
}

var Group = new(group)

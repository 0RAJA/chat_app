package logic

type group struct {
	User  user
	Email email
	File file
	Account     account
	Application application
}

var Group = new(group)

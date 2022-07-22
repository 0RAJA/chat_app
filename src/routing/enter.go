package routing

type group struct {
	User    user
	Email   email
	Account account
	File   file
	Application application
}

var Group = new(group)

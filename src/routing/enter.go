package routing

type group struct {
	User    user
	Email   email
	Account account
	File   file
}

var Group = new(group)

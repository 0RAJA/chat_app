package v1

type group struct {
	User    user
	Email   email
	Account account
}

var Group = new(group)

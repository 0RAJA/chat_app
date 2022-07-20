package v1

type group struct {
	User  user
	Email email
	File file
	Account account
}

var Group = new(group)

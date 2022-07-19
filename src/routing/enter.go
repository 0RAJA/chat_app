package routing

type group struct {
	User  user
	Email email
	File   file
}

var Group = new(group)

package v1

type group struct {
	User  user
	Email email
	File file
}

var Group = new(group)

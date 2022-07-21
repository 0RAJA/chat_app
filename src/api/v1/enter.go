package v1

type group struct {
	User        user
	Email       email
	Account     account
	Application application
}

var Group = new(group)

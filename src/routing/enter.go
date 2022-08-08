package routing

type group struct {
	User        user
	Email       email
	File        file
	Account     account
	Application application
	Notify      notify
	Setting     setting
	MGroup      mGroup
	Message     message
	Ws          ws
}

var Group = new(group)

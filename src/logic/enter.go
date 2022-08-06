package logic

type group struct {
	User        user
	Email       email
	File        file
	Account     account
	Application application
	Notify      notify
	Setting     setting
	MGroup      mGroup
	Auto        auto
	Message     message
}

var Group = new(group)

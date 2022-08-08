package logic

import (
	"github.com/0RAJA/chat_app/src/logic/chat"
)

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
	Chat        chat.Group
}

var Group = new(group)

package model

import (
	"encoding/json"

	"github.com/0RAJA/Rutils/pkg/token"
)

type TokenType string

const (
	UserToken    TokenType = "user"
	AccountToken TokenType = "account"
)

type Content struct {
	Type TokenType `json:"type"`
	ID   int64     `json:"id"`
}

func NewTokenContent(t TokenType, ID int64) *Content {
	return &Content{Type: t, ID: ID}
}

func (c *Content) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Content) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	return nil
}

// Token 结合token.Payload和Token
type Token struct {
	Payload *token.Payload
	Content *Content
}

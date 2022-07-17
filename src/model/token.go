package model

import (
	"encoding/json"
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

func (c *Content) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Content) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	return nil
}

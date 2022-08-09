package server

type AccountLogin struct {
	EnToken string `json:"en_token"` // 加密后的Token
	Address string `json:"address"`  // 地址
}

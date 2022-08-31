package server

// chat 中 server 端有关处理的结构

type AccountLogin struct {
	EnToken string `json:"en_token"` // 加密后的Token
	Address string `json:"address"`  // 地址
}

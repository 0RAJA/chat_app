package server

type UpdateAccount struct {
	EnToken   string `json:"en_token"`  // 加密后的Token
	Name      string `json:"name"`      // 昵称
	Gender    string `json:"gender"`    // 性别
	Signature string `json:"signature"` // 签名
}

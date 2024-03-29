package domain

type Session struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
	IP       string `json:"ip"`
	UserID   UserID `json:"userId"`
}

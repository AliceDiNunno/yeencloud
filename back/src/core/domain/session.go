package domain

type Session struct {
	Token    string
	ExpireAt int64
	IP       string
	UserID   string
}

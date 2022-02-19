package entity

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token	string `json:"token"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	TokenString string `json:"token"`
}
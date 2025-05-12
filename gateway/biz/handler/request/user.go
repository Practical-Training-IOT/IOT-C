package request

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoUpdate struct {
	Department string `json:"department"`
	Email      string `json:"email"`
	Phone      int    `json:"phone"`
	Username   string `json:"username"`
}

package user

type LoginInputData struct {
	Username string `example:"username" json:"username"`
	Password string `example:"pass" json:"password"`
}

type LoginOutputData struct {
	Token string `example:"gadsgasfdasfds" json:"token"`
}

type LoginErrorData struct {
	Error string `example:"Invalid credentials" json:"error"`
}

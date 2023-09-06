package model

type SignInInput struct {
	Passport string
	Password string
}

type SignUpInput struct {
	Nickname   string
	Passport   string
	Password   string
	VerifyKey  string
	VerifyCode string
}

type CaptchaInput struct {
}
